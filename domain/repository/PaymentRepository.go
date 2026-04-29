package repository

import (
	"github.com/payOSHQ/payos-lib-golang/v2"
)

type PaymentData struct {
	OrderCode    int64
	Amount       int
	Description  string
	CancelUrl    string
	ReturnUrl    string
	Items        []payos.PaymentLinkItem
	BuyerName    *string
	BuyerEmail   *string
	BuyerPhone   *string
	BuyerAddress *string
}

type PaymentCofirmation struct {
	Code      string                 `json:"code"`
	Desc      string                 `json:"desc"`
	Success   bool                   `json:"success"`
	Data      PaymentCofirmationData `json:"data"`
	Signature string                 `json:"signature"`
}

type PaymentCofirmationData struct {
	OrderCode              int64  `json:"orderCode"`
	Amount                 int    `json:"amount"`
	Description            string `json:"description"`
	AccountNumber          string `json:"accountNumber"`
	Reference              string `json:"reference"`
	TransactionDateTime    string `json:"transactionDateTime"`
	Currency               string `json:"currency"`
	PaymentLinkID          string `json:"paymentLinkId"`
	Code                   string `json:"code"`
	Desc                   string `json:"desc"`
	CounterAccountBankID   string `json:"counterAccountBankId"`
	CounterAccountBankName string `json:"counterAccountBankName"`
	CounterAccountName     string `json:"counterAccountName"`
	CounterAccountNumber   string `json:"counterAccountNumber"`
	VirtualAccountName     string `json:"virtualAccountName"`
	VirtualAccountNumber   string `json:"virtualAccountNumber"`
}

type PaymentRepository interface {
	Initialize()
	CreatePaymentRequest(data PaymentData) (*payos.CreatePaymentLinkResponse, error)
	VerifyPayment(input map[string]interface{}) (interface{}, error)
	ResolveOrderIdFromOderCode(orderCode int64) (int, error)
}
