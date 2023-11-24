package gocryptomus

const APIURL = "https://api.cryptomus.com/v1/"

type Cryptomus struct {
	Merchant      string
	PaymentApiKey string
	PayoutApiKey  string
}

func New(merchant, paymentApiKey, payoutApiKey string) *Cryptomus {
	return &Cryptomus{
		Merchant:      merchant,
		PaymentApiKey: paymentApiKey,
		PayoutApiKey:  payoutApiKey,
	}
}
