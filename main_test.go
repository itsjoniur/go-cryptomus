package gocryptomus

import (
	"net/http"
	"os"
	"testing"
)

var TestCryptomus *Cryptomus

func TestMain(m *testing.M) {
	httpClient := http.Client{}
	// TestCryptomus = &Cryptomus{
	// 	Client:        &httpClient,
	// 	Merchant:      "replace with your merchant id",
	// 	PaymentApiKey: "replace with your payment API key",
	// 	PayoutApiKey:  "replace with your payout API key",
	// }
	merchant := "replace with your merchant id"
	paymentAPIKey := "replace with your payment API key"
	payoutAPIKey := "replace with your payout API key"
	TestCryptomus = New(&httpClient, merchant, paymentAPIKey, payoutAPIKey)

	os.Exit(m.Run())
}
