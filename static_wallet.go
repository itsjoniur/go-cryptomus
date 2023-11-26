package gocryptomus

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const CreateStaticWalletEndpoint = "/wallet"

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

type StaticWalletRawResponse struct {
	Result *StaticWalletResponse `json:"result"`
	State  int8                  `json:"state"`
}

func (c *Cryptomus) CreateStaticWallet(staticWalletReq *StaticWalletRequest) (*StaticWalletResponse, error) {
	payload, err := json.Marshal(staticWalletReq)
	if err != nil {
		return nil, err
	}

	sign := c.SignRequest(c.PaymentApiKey, payload)
	req, err := http.NewRequest("POST", APIURL+CreateStaticWalletEndpoint, bytes.NewBuffer(payload))
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

	response := &StaticWalletRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}
