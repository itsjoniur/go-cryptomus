package tests

import (
	"net/http"
	"os"
	"testing"

	"github.com/itsjoniur/go-cryptomus"
)

var TestCryptomus *cryptomus.Cryptomus

func TestMain(m *testing.M) {
	httpClient := http.Client{}
	merchant := "replace with your merchant id"
	paymentAPIKey := "replace with your payment API key"
	payoutAPIKey := "replace with your payout API key"
	TestCryptomus = cryptomus.New(&httpClient, merchant, paymentAPIKey, payoutAPIKey)

	os.Exit(m.Run())
}
