package repositoryimpl

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/payOSHQ/payos-lib-golang/v2"
)

type PaymentRepositoryImpl struct{}

type confirmWebhookRequest struct {
	WebhookUrl string `json:"webhookUrl"`
}

var PayOsClient *payos.PayOS

var initialized = false

var isReady = false

func getApiKey() string {
	return os.Getenv("PAYOS_API_KEY")
}

func getClientId() string {
	return os.Getenv("PAYOS_CLIENT_ID")
}

func getCheckSumKey() string {
	return os.Getenv("PAYOS_CHECKSUM_KEY")
}

func getWebhook() string {
	return os.Getenv("PAYOS_WEBHOOK")
}

func (p *PaymentRepositoryImpl) Initialize() {
	client, err := payos.NewPayOS(&payos.PayOSOptions{
		ClientId:    getClientId(),
		ApiKey:      getApiKey(),
		ChecksumKey: getCheckSumKey(),
	})
	if err != nil {
		log.Fatal(err)
	} else {
		PayOsClient = client
		initialized = true
		p.registerOrConfirmWebhook(context.Background())
	}
}

func (p *PaymentRepositoryImpl) registerOrConfirmWebhook(ctx context.Context) {
	go func() {
		reqCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()

		webhook := os.Getenv("PAYOS_WEBHOOK")
		payload, err := json.Marshal(confirmWebhookRequest{WebhookUrl: webhook})
		if err != nil {
			log.Println("confirm webhook: marshal error:", err)
			return
		}

		req, err := http.NewRequestWithContext(
			reqCtx,
			http.MethodPost,
			"https://api-merchant.payos.vn/confirm-webhook",
			bytes.NewBuffer(payload),
		)
		if err != nil {
			log.Println("confirm webhook: new request error:", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", getApiKey())
		req.Header.Set("x-client-id", getClientId())

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("confirm webhook: request error:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			log.Println("confirm webhook: non-2xx status:", resp.StatusCode)
			return
		}
		isReady = true
		log.Println("register webhook success", resp.StatusCode)
	}()
}

func (p *PaymentRepositoryImpl) CreatePaymentRequest() (*payos.CreatePaymentLinkResponse, error) {
	if PayOsClient == nil || !initialized || !isReady {
		return nil, errors.New("Client is not ready now")
	}
	paymentLink, err := PayOsClient.PaymentRequests.Create(context.Background(), payos.CreatePaymentLinkRequest{
		OrderCode:   123,
		Amount:      2000,
		Description: "Payment ",
		ReturnUrl:   "https://go.dev/",
		CancelUrl:   "https://www.google.com/?hl=vi",
	})
	if err != nil || paymentLink == nil {
		return nil, errors.New("Create payment link failed")
	}
	return paymentLink, nil
}
