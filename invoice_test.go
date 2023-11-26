package gocryptomus

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateInvoice(t *testing.T) {
	httpClient := http.Client{}
	merchant := "c7dc770e-656a-4cff-b1d2-8971efe4b17e"
	apiKey := "Oz127LTxuzHBR0c0vh7gCYiTu471Gr7GuLXtzszFNwWRV3qXUWM5BnUwLVrrh9x0KvmLYfbxbnKz1tPteLPmDpcHxKtOdmLxGLuJwgLXn24y27cisDZ7asTFdXYIiFvG"
	c := New(&httpClient, merchant, apiKey, "")

	invoiceReq := &InvoiceRequest{
		Amount:   "10",
		Currency: "USD",
		OrderId:  "xxx",
		InvoiceRequestOptions: &InvoiceRequestOptions{
			Network:     "tron",
			UrlCallback: "https://example.com/cryptomus/callback",
		},
	}
	invoice, err := c.CreateInvoice(invoiceReq)
	require.NoError(t, err)
	require.NotEmpty(t, invoice)
}
