package utils

import (
	"fmt"
	"net/smtp"
	"os"
	"time"
)

func SendOrderPaidEmail(to string, orderCode int64, paidAt time.Time) error {
	from := "doankhacminh2301@gmail.com"
	password := os.Getenv("EAP")

	host := "smtp.gmail.com"
	port := "587"

	subject := "WeFashion - Xac nhan thanh toan"
	body := fmt.Sprintf(
		"Ban da thanh toan thanh cong cho don hang %d vao luc %s. Don hang da duoc xac nhan va se som duoc giao toi ban trong thoi gian ngan nhat.",
		orderCode,
		paidAt.Format("2006-01-02 15:04:05"),
	)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", from, password, host)

	addr := host + ":" + port
	return smtp.SendMail(addr, auth, from, []string{to}, msg)
}
