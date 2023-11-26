package gocryptomus

import "net/http"

const APIURL = "https://api.cryptomus.com/v1"

type Cryptomus struct {
	merchant      string
	paymentApiKey string
	payoutApiKey  string
	client        *http.Client
}

func New(client *http.Client, merchant, paymentApiKey, payoutApiKey string) *Cryptomus {
	return &Cryptomus{
		client:        client,
		merchant:      merchant,
		paymentApiKey: paymentApiKey,
		payoutApiKey:  payoutApiKey,
	}
}
