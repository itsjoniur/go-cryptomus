package gocryptomus

import (
	"encoding/json"
	"errors"
	"time"
)

const (
	createInvoiceEndpoit          = "/payment"
	generateInvoiceQRCodeEndpoint = "/payment/qr"
	paymentInfoEndpoint           = "/payment/info"
	paymentHistoryEndpoint        = "/payment/list"
)

type InvoiceRequest struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	OrderId  string `json:"order_id"`
	*InvoiceRequestOptions
}

type InvoiceRequestOptions struct {
	Network                string `json:"network,omitempty"`
	UrlReturn              string `json:"url_return,omitempty"`
	UrlSuccess             string `json:"url_success,omitempty"`
	UrlCallback            string `json:"url_callback,omitempty"`
	IsPaymentMultiple      bool   `json:"is_payment_multiple,omitempty"`
	Lifetime               uint16  `json:"lifetime,omitempty"`
	ToCurrency             string `json:"to_currency,omitempty"`
	Subtract               uint8  `json:"subtract,omitempty"`
	AccuarcyPaymentPercent uint8  `json:"accuarcy_payment_percent,omitempty"`
	AdditionalData         string `json:"additional_data,omitempty"`
	Currencies             []struct {
		Currency string `json:"currency"`
		Network  string `json:"network,omitempty"`
	} `json:"currencies,omitempty"`
	ExceptCurrencies []struct {
		Currency string `json:"currency"`
		Network  string `json:"network,omitempty"`
	} `json:"except_currencies,omitempty"`
	CourseSource     string `json:"course_source,omitempty"`
	FromReferralCode string `json:"from_referral_code,omitempty"`
	DiscountPercent  int8   `json:"discount_percent,omitempty"`
	IsRefresh        bool   `json:"is_refresh,omitempty"`
}

type Payment struct {
	UUID                    string    `json:"uuid"`
	OrderId                 string    `json:"order_id"`
	Amount                  string    `json:"amount"`
	PaymentAmount           string    `json:"payment_amount,omitempty"`
	PaymentAmountUSD        string    `json:"payment_amount_usd,omitempty"`
	PayerAmount             string    `json:"payer_amount,omitempty"`
	PayerAmountExchangeRate string    `json:"payer_amount_exchange_rate,omitempty"`
	DiscountPercent         string    `json:"discount_percent,omitempty"`
	Discount                string    `json:"discount,omitempty"`
	PayerCurrency           string    `json:"payer_currency,omitempty"`
	Currency                string    `json:"currency"`
	MerchantAmount          uint32    `json:"merchant_amount,omitempty"`
	Network                 string    `json:"network,omitempty"`
	Address                 string    `json:"address,omitempty"`
	From                    string    `json:"from,omitempty"`
	TxId                    string    `json:"txid,omitempty"`
	PaymentStatus           string    `json:"payment_status"`
	Status                  string    `json:"status,omitempty"`
	Url                     string    `json:"url"`
	ExpiredAt               float64   `json:"expired_at"`
	IsFinal                 bool      `json:"is_final"`
	AdditionalData          string    `json:"additional_data,omitempty"`
	Comments                string    `json:"comments,omitempty"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

type invoiceRawResponse struct {
	Result *Payment
	State  int8
}

type paymentQRCodeRawResponse struct {
	Result struct {
		Image string `json:"image"`
	} `json:"result"`
	State int8 `json:"state"`
}

type PaymentInfoRequest struct {
	PaymentUUID string `json:"uuid,omitempty"`
	OrderId     string `json:"order_id,omitempty"`
}

type PaymentHistoryResponse struct {
	Payments []*Payment
	Paginate *PaymentHistoryPaginate
}

type PaymentHistoryPaginate struct {
	Count          int16   `json:"count"`
	HasPages       bool   `json:"hasPages"`
	NextCursor     string `json:"nextCursor,omitempty"`
	PreviousCursor string `json:"previousCursor,omitempty"`
	PerPage        int16   `json:"perPage"`
}

type paymentHistoryRawResponse struct {
	State    int8                    `json:"state"`
	Result   []*Payment              `json:"result"`
	Paginate *PaymentHistoryPaginate `json:"paginate"`
}

func (c *Cryptomus) CreateInvoice(invoiceReq *InvoiceRequest) (*Payment, error) {
	res, err := c.fetch("POST", createInvoiceEndpoit, invoiceReq)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &invoiceRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (c *Cryptomus) GeneratePaymentQRCode(paymentUUID string) (string, error) {
	payload := map[string]any{"merchant_payment_uuid": paymentUUID}
	res, err := c.fetch("POST", generateInvoiceQRCodeEndpoint, payload)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	response := &paymentQRCodeRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return "", err
	}

	return response.Result.Image, nil

}

func (c *Cryptomus) GetPaymentInfo(paymentInfoReq *PaymentInfoRequest) (*Payment, error) {
	if paymentInfoReq.PaymentUUID == "" || paymentInfoReq.OrderId == "" {
		return nil, errors.New("you should pass one of required values [PaymentUUID, OrderId]")
	}

	res, err := c.fetch("POST", paymentInfoEndpoint, paymentInfoReq)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &invoiceRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (c *Cryptomus) GetPaymentHistory(dateFrom, dateTo time.Time) (*PaymentHistoryResponse, error) {
	payload := map[string]any{"date_from": dateFrom, "date_to": dateTo}
	res, err := c.fetch("POST", paymentHistoryEndpoint, payload)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &paymentHistoryRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	paymentHistory := &PaymentHistoryResponse{
		Payments: response.Result,
		Paginate: response.Paginate,
	}
	return paymentHistory, nil
}
