// Kiểm tra dữ liệu đầu vào đăng ký tài khoản
package account

import (
	"WeFashionServer/domain/entity"
	"WeFashionServer/domain/handler/authentication"
	"WeFashionServer/domain/helper"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/infrastructure/model"
	"context"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func validateRegisterInput(req *struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}, ctx *gin.Context) bool {
	if strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.Password) == "" {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Missing required fields",
			Detail:     "Email, username, and password are required.",
		})
		return false
	}
	if ok, err := helper.ValidateEmail(context.Background(), req.Email, false); !ok {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid email format",
			Detail:     err.Error(),
		})
		return false
	}
	return true
}

// Kiểm tra email hoặc username đã tồn tại
func checkEmailUsernameExists(email, username string, ctx *gin.Context) bool {
	var countEmail int64
	database.DB.Model(&model.User{}).Where("email = ?", email).Count(&countEmail)
	var countUsername int64
	database.DB.Model(&model.Account{}).Where("username = ?", username).Count(&countUsername)
	if countEmail > 0 || countUsername > 0 {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Email or username already exists",
			Detail:     "Email or username already exists.",
		})
		return false
	}
	return true
}

// Tạo user mới
func createUser(username, email string) (model.User, error) {
	user := model.User{
		Name:        username,
		AvatarUrl:   "https://avatars.githubusercontent.com/u/100721386?v=4",
		Email:       email,
		PhoneNumber: "",
		Bio:         "",
	}
	err := database.DB.Create(&user).Error
	return user, err
}

// Tạo account mới
func createAccount(username, hashedPassword string, userId int) error {
	account := model.Account{
		Username: username,
		Password: hashedPassword,
		UserId:   userId,
	}
	return database.DB.Create(&account).Error
}

func hash(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(bytes), err
}

func isValidCode(code string) bool {
	if len(code) != 4 {
		return false
	}
	for _, c := range code {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func sendEmail(to, code string) error {
	from := "doankhacminh2301@gmail.com"
	password := os.Getenv("EAP")

	host := "smtp.gmail.com"
	port := "587"

	subject := "WeFashion - Lấy lại tài khoản"
	body := fmt.Sprintf("Mã xác thực của bạn là: %s", code)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", from, password, host)

	addr := host + ":" + port
	return smtp.SendMail(addr, auth, from, []string{to}, msg)
}

// POST /api/account/register
func RegisterAccount(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	var req struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid request body",
			Detail:     err.Error(),
		})
		return
	}
	if !validateRegisterInput(&req, ctx) {
		return
	}
	if !checkEmailUsernameExists(req.Email, req.Username, ctx) {
		return
	}
	hashed, err := hash(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Hash password failed",
			Detail:     err.Error(),
		})
		return
	}
	user, err := createUser(req.Username, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Create user failed",
			Detail:     err.Error(),
		})
		return
	}
	if err := createAccount(req.Username, hashed, user.Id); err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Create account failed",
			Detail:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[model.User]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       user,
	})
}

// POST /api/account/login
func LoginAccount(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid request body",
			Detail:     err.Error(),
		})
		return
	}
	var account model.Account
	if err := database.DB.Where("username = ?", req.Username).First(&account).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, entity.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Error:      "Username not found",
			Detail:     "Username does not exist.",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, entity.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Error:      "Invalid password",
			Detail:     "Password is incorrect.",
		})
		return
	}
	// Truy vấn user tương ứng
	var user model.User
	if err := database.DB.Where("id = ?", account.UserId).First(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "User not found",
			Detail:     "Cannot find user for this account.",
		})
		return
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[model.User]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       user,
	})
}

// POST /api/account/forgot-password
func ForgotPassword(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	var req struct {
		Email string `json:"email"`
		// Code  string `json:"code"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid request body",
			Detail:     err.Error(),
		})
		return
	}
	if ok, err := helper.ValidateEmail(context.Background(), req.Email, false); !ok {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid email format",
			Detail:     err.Error(),
		})
		return
	}
	var user model.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "Email not found",
			Detail:     "No user with this email.",
		})
		return
	}
	code := fmt.Sprintf("%04d", time.Now().UnixMilli()%10000)
	if err := sendEmail(req.Email, code); err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Send email failed",
			Detail:     err.Error(),
		})
		return
	}
	credential, codeGenErr := hash(code)
	if codeGenErr != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Build credential error",
			Detail:     codeGenErr.Error(),
		})
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[map[string]interface{}]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data: map[string]interface{}{
			// "launch_time": launchTime,
			// "expired_at":  expiredAt,
			"credential": credential,
		},
	})
}

// POST /api/account/forgot-password/validate
func ValidateCode(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	var req struct {
		Email      string `json:"email"`
		Code       string `json:"code"`
		Credential string `json:"credential"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid request body",
			Detail:     err.Error(),
		})
		return
	}
	if ok, err := helper.ValidateEmail(context.Background(), req.Email, false); !ok {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid email format",
			Detail:     err.Error(),
		})
		return
	}
	if !isValidCode(req.Code) {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid code",
			Detail:     "Code must be 4 digits.",
		})
		return
	}
	if req.Credential == "" {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid credential",
			Detail:     "",
		})
		return
	}
	var user model.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "Email not found",
			Detail:     "No user with this email.",
		})
		return
	}

	var isValid bool
	var detail string
	if err := bcrypt.CompareHashAndPassword([]byte(req.Credential), []byte(req.Code)); err != nil {
		isValid = false
		detail = "Incorrect code"
	} else {
		isValid = true
		detail = "Correct code"
	}

	ctx.JSON(http.StatusOK, entity.SuccessReponse[map[string]interface{}]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data: map[string]interface{}{
			"is_valid": isValid,
			"Detail":   detail,
		},
	})
}

func ChangePassword(ctx *gin.Context) {
	// Validate token first
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}

	// Parse request body
	var req struct {
		Username    string `json:"username"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid request body",
			Detail:     err.Error(),
		})
		return
	}

	// Check for empty fields
	if strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.OldPassword) == "" || strings.TrimSpace(req.NewPassword) == "" {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Missing required fields",
			Detail:     "Username, old_password, and new_password are required.",
		})
		return
	}

	// Check if account exists
	var account model.Account
	if err := database.DB.Where("username = ?", req.Username).First(&account).Error; err != nil {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "Username not found",
			Detail:     "Account with this username does not exist.",
		})
		return
	}

	// Check old password
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.OldPassword)); err != nil {
		ctx.JSON(http.StatusUnauthorized, entity.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Error:      "Invalid old password",
			Detail:     "Old password is incorrect.",
		})
		return
	}

	// Hash new password
	hashed, err := hash(req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Hash password failed",
			Detail:     err.Error(),
		})
		return
	}

	// Update password in database
	if err := database.DB.Model(&account).Update("password", hashed).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Update password failed",
			Detail:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, entity.SuccessReponse[map[string]interface{}]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data: map[string]interface{}{
			"message": "Password changed successfully",
		},
	})
}
