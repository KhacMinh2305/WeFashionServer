package repositoryimpl

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"WeFashionServer/domain/repository"

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

func getOrderCodePrefix() int64 {
	num, err := strconv.Atoi(os.Getenv("PAYOS_ORDER_CODE_PREFIX"))
	if err != nil {
		return 0
	}
	return int64(num)
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

func (p *PaymentRepositoryImpl) CreatePaymentRequest(data repository.PaymentData) (*payos.CreatePaymentLinkResponse, error) {
	if PayOsClient == nil || !initialized || !isReady {
		return nil, errors.New("Client is not ready now")
	}
	paymentLink, err := PayOsClient.PaymentRequests.Create(context.Background(), payos.CreatePaymentLinkRequest{
		OrderCode:    getOrderCodePrefix() + data.OrderCode,
		Amount:       data.Amount,
		Description:  data.Description,
		CancelUrl:    data.CancelUrl,
		ReturnUrl:    data.ReturnUrl,
		Items:        data.Items,
		BuyerName:    data.BuyerName,
		BuyerEmail:   data.BuyerEmail,
		BuyerPhone:   data.BuyerPhone,
		BuyerAddress: data.BuyerAddress,
	})
	if err != nil || paymentLink == nil {
		fmt.Println("------------Error------------")
		fmt.Println(err.Error())
		return nil, errors.New("Create payment link failed")
	}
	return paymentLink, nil
}

func (p *PaymentRepositoryImpl) VerifyPayment(input map[string]interface{}) (interface{}, error) {
	return PayOsClient.Webhooks.VerifyData(context.Background(), input)
}

func (p *PaymentRepositoryImpl) ResolveOrderIdFromOderCode(orderCode int64) (int, error) {
	orderId := int64(orderCode) - getOrderCodePrefix()
	if orderId <= 0 {
		return 0, errors.New("Invalid order id")
	}
	return int(orderId), nil
}
