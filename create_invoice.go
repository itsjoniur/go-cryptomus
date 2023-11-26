package gocryptomus

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const CreateInvoiceEndpoit = "/payment"

type NewInvoice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	OrderId  string `json:"order_id"`
	*NewInvoiceOptions
}

type NewInvoiceOptions struct {
	Network                string `json:"network,omitempty"`
	UrlReturn              string `json:"url_return,omitempty"`
	UrlSuccess             string `json:"url_success,omitempty"`
	UrlCallback            string `json:"url_callback,omitempty"`
	IsPaymentMultiple      bool   `json:"is_payment_multiple,omitempty"`
	Lifetime               uint8  `json:"lifetime,omitempty"`
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

type NewInvoiceResponse struct {
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

type NewInvoiceRawResponse struct {
	Result *NewInvoiceResponse
	State  int8
}

func (c *Cryptomus) CreateInvoice(amount string, currency, orderId string, opts *NewInvoiceOptions) (*NewInvoiceResponse, error) {
	invoice := NewInvoice{
		Amount:            amount,
		Currency:          currency,
		OrderId:           orderId,
		NewInvoiceOptions: opts,
	}
	payload, err := json.Marshal(invoice)
	if err != nil {
		return nil, err
	}

	sign := c.SignRequest(c.PaymentApiKey, payload)
	req, err := http.NewRequest("POST", APIURL+CreateInvoiceEndpoit, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("merchant", c.Merchant)
	req.Header.Set("sign", sign)
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &NewInvoiceRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}
