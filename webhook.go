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
