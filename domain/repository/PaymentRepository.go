package repository

import "github.com/payOSHQ/payos-lib-golang/v2"

type PaymentRepository interface {
	Initialize()
	CreatePaymentRequest() (*payos.CreatePaymentLinkResponse, error)
}
