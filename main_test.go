package gocryptomus

import (
	"net/http"
	"os"
	"testing"
)

var TestCryptomus *Cryptomus

func TestMain(m *testing.M) {
	httpClient := http.Client{}
	merchant := "replace with your merchant id"
	paymentAPIKey := "replace with your payment API key"
	payoutAPIKey := "replace with your payout API key"
	TestCryptomus = New(&httpClient, merchant, paymentAPIKey, payoutAPIKey)

	os.Exit(m.Run())
}
