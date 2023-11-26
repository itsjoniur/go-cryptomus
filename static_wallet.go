package gocryptomus

import (
	"encoding/json"
	"errors"
)

const (
	createStaticWalletEndpoint         = "/wallet"
	generateStaticWalletQRCodeEndpoint = "/wallet/qr"
	blockWalletAddressEndpoint         = "/wallet/block-address"
)

type StaticWalletRequest struct {
	Currency string `json:"currency"`
	Network  string `json:"network"`
	OrderId  string `json:"order_id"`
	*StaticWalletRequestOptions
}

type StaticWalletRequestOptions struct {
	UrlCallback      string `json:"url_callback,omitempty"`
	FromReferralCode string `json:"from_referral_code,omitempty"`
}

type StaticWalletResponse struct {
	OrderId    string `json:"order_id"`
	WalletUUID string `json:"wallet_uuid"`
	UUID       string `json:"uuid"`
	Address    string `json:"address"`
	Network    string `json:"network"`
	Currency   string `json:"currency"`
	Url        string `json:"url"`
}

type staticWalletRawResponse struct {
	Result *StaticWalletResponse `json:"result"`
	State  int8                  `json:"state"`
}

type staticWalletQRCodeRawResponse struct {
	Result struct {
		Image string `json:"image"`
	} `json:"result"`
	State int8 `json:"state"`
}

type BlockAddressRequest struct {
	WalletUUID    string `json:"uuid,omitempty"`
	OrderId       string `json:"order_id,omitempty"`
	IsForceRefund bool   `json:"is_force_refund,omitempty"`
}

type BlockAddressResponse struct {
	WalletUUID string `json:"uuid"`
	Status     string `json:"status"`
}

type blockAddressRawResponse struct {
	Result *BlockAddressResponse
	State  int8
}

func (c *Cryptomus) CreateStaticWallet(staticWalletReq *StaticWalletRequest) (*StaticWalletResponse, error) {
	res, err := c.fetch("POST", createStaticWalletEndpoint, staticWalletReq)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &staticWalletRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (c *Cryptomus) GenerateStaticWalletQRCode(walletUUID string) (string, error) {
	payload := map[string]any{"wallet_address_uuid": walletUUID}
	res, err := c.fetch("POST", generateStaticWalletQRCodeEndpoint, payload)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	response := &staticWalletQRCodeRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return "", err
	}

	return response.Result.Image, nil
}

func (c *Cryptomus) BlockAddress(blockAddressReq *BlockAddressRequest) (*BlockAddressResponse, error) {
	if blockAddressReq.WalletUUID == "" || blockAddressReq.OrderId == "" {
		return nil, errors.New("you should pass one of required values [WalletUUID, OrderId]")
	}

	res, err := c.fetch("POST", blockWalletAddressEndpoint, blockAddressReq)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &blockAddressRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}
