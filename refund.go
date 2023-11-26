package gocryptomus

import (
	"encoding/json"
	"errors"
)

const (
	refundEndpoint               = "/payment/refund"
	blockedAddressRefundEndpoint = "/wallet/blocked-address-refund"
)

type RefundRequest struct {
	Address     string `json:"address"`
	IsSubtract  bool   `json:"is_subtract"`
	PaymentUUID string `json:"uuid,omitempty"`
	OrderId     string `json:"order_id,omitempty"`
}

type refundRawResponse struct {
	Result []string `json:"result,omitempty"`
	State  int8     `json:"state"`
}

type BlockedAddressRefundRequest struct {
	WalletUUID string `json:"uuid,omitempty"`
	OrderId    string `json:"order_id,omitempty"`
	Address    string `json:"address"`
}

type BlockedAddressRefundResponse struct {
	Commision string `json:"commision"`
	Amount    string `json:"amount"`
}

type blockedAddressRefundRawResponse struct {
	Result *BlockedAddressRefundResponse `json:"result"`
	State  int8                          `json:"state"`
}

func (c *Cryptomus) Refund(refundRequest *RefundRequest) (bool, error) {
	res, err := c.fetch("POST", refundEndpoint, refundRequest)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	response := &refundRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return false, err
	}

	return len(response.Result) == 0, nil
}

func (c *Cryptomus) BlockedAddressRefund(refundRequest *BlockedAddressRefundRequest) (*BlockedAddressRefundResponse, error) {
	if refundRequest.WalletUUID == "" || refundRequest.OrderId == "" {
		return nil, errors.New("you should pass one of required values [WalletUUID, OrderId]")
	}

	res, err := c.fetch("POST", blockedAddressRefundEndpoint, refundRequest)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &blockedAddressRefundRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}
