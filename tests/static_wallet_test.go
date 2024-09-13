package tests

import (
	"testing"

	"github.com/itsjoniur/go-cryptomus"

	"github.com/stretchr/testify/require"
)

func TestCreateStaticWallet(t *testing.T) {
	staticWalletReq := &cryptomus.StaticWalletRequest{
		Currency: "TRX",
		Network:  "tron",
		OrderId:  "xxx",
		StaticWalletRequestOptions: &cryptomus.StaticWalletRequestOptions{
			UrlCallback: "https://example.com/cryptomus/callback",
		},
	}

	staticWallet, err := TestCryptomus.CreateStaticWallet(staticWalletReq)
	require.NoError(t, err)
	require.NotEmpty(t, staticWallet)
}
