package gocryptomus

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateInvoice(t *testing.T) {
	invoiceReq := &InvoiceRequest{
		Amount:   "10",
		Currency: "USD",
		OrderId:  "xxx",
		InvoiceRequestOptions: &InvoiceRequestOptions{
			Network:     "tron",
			UrlCallback: "https://example.com/cryptomus/callback",
		},
	}
	invoice, err := TestCryptomus.CreateInvoice(invoiceReq)
	require.NoError(t, err)
	require.NotEmpty(t, invoice)
}
