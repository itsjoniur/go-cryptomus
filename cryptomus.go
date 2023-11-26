package gocryptomus

import "net/http"

const APIURL = "https://api.cryptomus.com/v1"

type Cryptomus struct {
	Merchant      string
	PaymentApiKey string
	PayoutApiKey  string
	Client        *http.Client
}

func New(client *http.Client, merchant, paymentApiKey, payoutApiKey string) *Cryptomus {
	return &Cryptomus{
		Client:        client,
		Merchant:      merchant,
		PaymentApiKey: paymentApiKey,
		PayoutApiKey:  payoutApiKey,
	}
}
