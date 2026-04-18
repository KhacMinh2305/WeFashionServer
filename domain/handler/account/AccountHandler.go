package account

import (
	"WeFashionServer/domain/entity"
	"WeFashionServer/domain/handler/authentication"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/infrastructure/model"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// --- helpers ---
func validateTokenOrAbort(ctx *gin.Context) bool {
	_, err := authentication.ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, entity.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Error:      "Unauthorized",
			Detail:     err.Error(),
		})
		return false
	}
	return true
}

// POST /api/account/register
func RegisterAccount(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
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
	if strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.Password) == "" {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Missing required fields",
			Detail:     "Email, username, and password are required.",
		})
		return
	}
	// Check email tồn tại trong User
	var countEmail int64
	database.DB.Model(&model.User{}).Where("email = ?", req.Email).Count(&countEmail)
	// Check username tồn tại trong Account
	var countUsername int64
	database.DB.Model(&model.Account{}).Where("username = ?", req.Username).Count(&countUsername)
	if countEmail > 0 || countUsername > 0 {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Email or username already exists",
			Detail:     "Email or username already exists.",
		})
		return
	}
	hashed, err := hashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Hash password failed",
			Detail:     err.Error(),
		})
		return
	}
	user := model.User{
		Name:        req.Username,
		AvatarUrl:   "https://avatars.githubusercontent.com/u/100721386?v=4",
		Email:       req.Email,
		PhoneNumber: "",
		Bio:         "",
	}
	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Create user failed",
			Detail:     err.Error(),
		})
		return
	}
	account := model.Account{
		Username: req.Username,
		Password: hashed,
		UserId:   user.Id,
	}
	if err := database.DB.Create(&account).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Create account failed",
			Detail:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[any]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       nil,
	})
}

// POST /api/account/login
func LoginAccount(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
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
	ctx.JSON(http.StatusOK, entity.SuccessReponse[any]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       nil,
	})
}

// POST /api/account/forgot-password
func ForgotPassword(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
		return
	}
	var req struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid request body",
			Detail:     err.Error(),
		})
		return
	}
	if !isValidEmail(req.Email) {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid email format",
			Detail:     "Email format is invalid.",
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
	if !isValidCode(req.Code) {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid code",
			Detail:     "Code must be 4 digits.",
		})
		return
	}
	launchTime := time.Now().UnixMilli()
	expiredAt := launchTime + 5*60*1000
	if err := sendEmail(req.Email, req.Code); err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Send email failed",
			Detail:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[map[string]interface{}]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data: map[string]interface{}{
			"launch_time": launchTime,
			"expired_at":  expiredAt,
		},
	})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
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
	body := fmt.Sprintf("Mã xác thực của bạn là: %s, có hiệu lực trong 5 phút", code)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", from, password, host)

	addr := host + ":" + port
	return smtp.SendMail(addr, auth, from, []string{to}, msg)
}
