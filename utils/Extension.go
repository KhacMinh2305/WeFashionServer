package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func sendEmail(to, title, emailBody string) error {

	payload := fmt.Sprintf(`{
        "personalizations": [{"to": [{"email": "%s"}]}],
        "from": {"email": "%s", "name": "WeFashion"},
        "subject": %s,
        "content": [{"type": "text/plain", "value": "%s"}]
    }`, to, os.Getenv("SENDGRID_FROM_EMAIL"), title, emailBody)

	req, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", strings.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SENDGRID_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("sendgrid error %d: %s", resp.StatusCode, string(b))
	}

	return nil
}

func SendVerificationEmail(to, code string) error {
	title := "WeFashion - Lấy lại tài khoản"
	body := fmt.Sprintf("Mã xác thực của bạn là: %s", code)
	return sendEmail(to, title, body)
}

func SendOrderPaidEmail(to string, orderCode int64, paidAt time.Time) error {
	title := "WeFashion - Xac nhan thanh toan"
	emailBody := fmt.Sprintf(
		"Ban da thanh toan thanh cong cho don hang %d vao luc %s. Don hang da duoc xac nhan va se som duoc giao den ban.",
		orderCode,
		paidAt.Format("2006-01-02 15:04:05"),
	)
	return sendEmail(to, title, emailBody)
}
