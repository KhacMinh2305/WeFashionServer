package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func sendEmail(to, title, emailBody string) error {
	type EmailContent struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}
	type EmailAddress struct {
		Email string `json:"email"`
	}
	type Personalization struct {
		To []EmailAddress `json:"to"`
	}
	type FromAddress struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	type Payload struct {
		Personalizations []Personalization `json:"personalizations"`
		From             FromAddress       `json:"from"`
		Subject          string            `json:"subject"`
		Content          []EmailContent    `json:"content"`
	}

	payload := Payload{
		Personalizations: []Personalization{
			{To: []EmailAddress{{Email: to}}},
		},
		From:    FromAddress{Email: os.Getenv("SENDGRID_FROM_EMAIL"), Name: "WeFashion"},
		Subject: title,
		Content: []EmailContent{
			{Type: "text/plain", Value: emailBody},
		},
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", bytes.NewReader(jsonBytes))
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
