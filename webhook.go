package gocryptomus

import (
	"encoding/json"
	"errors"
)

const (
	resendWebhookEndpoint = "/payment/resend"
)

type ResendWebhookRequest struct {
	PaymentUUID string `json:"uuid,omitempty"`
	OrderId     string `json:"order_id,omitempty"`
}

type resendWebhookRawResponse struct {
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
