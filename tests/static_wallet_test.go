package tests

import (
	"github.com/itsjoniur/go-cryptomus"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateStaticWallet(t *testing.T) {
	staticWalletReq := &gocryptomus.StaticWalletRequest{
		Currency: "TRX",
		Network:  "tron",
		OrderId:  "xxx",
		StaticWalletRequestOptions: &gocryptomus.StaticWalletRequestOptions{
			UrlCallback: "https://example.com/cryptomus/callback",
		},
	}

	staticWallet, err := TestCryptomus.CreateStaticWallet(staticWalletReq)
	require.NoError(t, err)
	require.NotEmpty(t, staticWallet)
}
