package gocryptomus

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateStaticWallet(t *testing.T) {
	staticWalletReq := &StaticWalletRequest{
		Currency: "TRX",
		Network:  "tron",
		OrderId:  "xxx",
		StaticWalletRequestOptions: &StaticWalletRequestOptions{
			UrlCallback: "https://example.com/cryptomus/callback",
		},
	}

	staticWallet, err := TestCryptomus.CreateStaticWallet(staticWalletReq)
	require.NoError(t, err)
	require.NotEmpty(t, staticWallet)
}
