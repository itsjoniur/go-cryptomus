package cryptomus

import (
	"encoding/json"
	"errors"
)

const (
	resendWebhookEndpoint      = "/payment/resend"
	testPaymentWebhookEndpoint = "/test-webhook/payment"
	testPayoutWebhookEndpoint  = "/test-webhook/payout"
)

type WebhookConvert struct {
	ToCurrency string `json:"to_currency"`
	Commission string `json:"commission"`
	Rate       string `json:"rate"`
	Amount     string `json:"amount"`
}

type Webhook struct {
	Type              string         `json:"type"`
	UUID              string         `json:"uuid"`
	OrderId           string         `json:"order_id"`
	Amount            string         `json:"amount"`
	PaymentAmount     string         `json:"payment_amount"`
	PaymentAmountUSD  string         `json:"payment_amount_usd"`
	MerchantAmount    string         `json:"merchant_amount"`
	Commission        string         `json:"commission"`
	IsFinal           bool           `json:"is_final"`
	Status            string         `json:"status"`
	From              string         `json:"from"`
	WalletAddressUUID string         `json:"wallet_address_uuid"`
	Network           string         `json:"network"`
	Currency          string         `json:"currency"`
	PayerCurrency     string         `json:"payer_currency"`
	AdditionalData    string         `json:"additional_data"`
	Convert           WebhookConvert `json:"convert"`
	TxId              string         `json:"txid"`
	Sign              string         `json:"sign"`
}

type ResendWebhookRequest struct {
	PaymentUUID string `json:"uuid,omitempty"`
	OrderId     string `json:"order_id,omitempty"`
}

type resendWebhookRawResponse struct {
	Result []string `json:"result"`
	State  int8     `json:"state"`
}

type TestWebhookRequest struct {
	UrlCallback string `json:"url_callback"`
	Currency    string `json:"currency"`
	Network     string `json:"network"`
	UUID        string `json:"uuid,omitempty"`
	OrderId     string `json:"order_id,omitempty"`
	Status      string `json:"status"`
}

type TestWebhookResponse struct {
	Result []string `json:"result"`
	State  int8     `json:"state"`
}

func (c *Cryptomus) ParseWebhook(reqBody []byte, verifySign bool) (*Webhook, error) {
	var apiKey string
	response := &Webhook{}

	err := json.Unmarshal(reqBody, response)
	if err != nil {
		return nil, err
	}

	switch response.Type {
	case "payment":
		apiKey = c.paymentApiKey
	case "payout":
		apiKey = c.payoutApiKey
	default:
		return nil, errors.New("unknown webhook type")
	}

	if verifySign {
		err = c.VerifySign(apiKey, reqBody)
		if err != nil {
			return nil, err
		}
	}

	return response, err
}

func (c *Cryptomus) ResendWebhook(resendRequest *ResendWebhookRequest) (bool, error) {
	if resendRequest.PaymentUUID == "" || resendRequest.OrderId == "" {
		return false, errors.New("you should pass one of required values [PaymentUUID, OrderId]")
	}

	res, err := c.fetch("POST", resendWebhookEndpoint, resendRequest)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	response := &resendWebhookRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return false, err
	}

	return len(response.Result) == 0, nil
}

func (c *Cryptomus) TestPaymentWebhook(testRequest *TestWebhookRequest) (*TestWebhookResponse, error) {
	res, err := c.fetch("POST", testPaymentWebhookEndpoint, testRequest)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &TestWebhookResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Cryptomus) TestPayoutWebhook(testRequest *TestWebhookRequest) (*TestWebhookResponse, error) {
	res, err := c.fetch("POST", testPayoutWebhookEndpoint, testRequest)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &TestWebhookResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}
