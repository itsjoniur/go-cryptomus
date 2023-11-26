package gocryptomus

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createTestInvoice(t *testing.T) *Payment {
	invoiceReq := &InvoiceRequest{
		Amount:   "10",
		Currency: "USD",
		OrderId:  "xxy",
		InvoiceRequestOptions: &InvoiceRequestOptions{
			Network:     "tron",
			UrlCallback: "https://example.com/cryptomus/callback",
		},
	}
	invoice, err := TestCryptomus.CreateInvoice(invoiceReq)
	require.NoError(t, err)
	require.NotEmpty(t, invoice)

	return invoice
}

func TestCreateInvoice(t *testing.T) {
	createTestInvoice(t)
}

func TestGenerateInvoiceQRCode(t *testing.T) {
	invoice := createTestInvoice(t)
	qrCode, err := TestCryptomus.GeneratePaymentQRCode(invoice.UUID)
	require.NoError(t, err)
	require.NotEmpty(t, qrCode)
}

func TestGetPaymentInfo(t *testing.T) {
	invoice := createTestInvoice(t)
	payment, err := TestCryptomus.GetPaymentInfo(&PaymentInfoRequest{PaymentUUID: invoice.UUID})
	require.NoError(t, err)
	require.NotEmpty(t, payment)
}

func TestGeyPaymentHistory(t *testing.T) {
	payments, err := TestCryptomus.GetPaymentHistory(time.Now(), time.Now())
	require.NoError(t, err)
	require.NotEmpty(t, payments)
}
