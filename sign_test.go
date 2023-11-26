package gocryptomus

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSignRequest(t *testing.T) {
	httpClient := http.Client{}
	c := New(&httpClient, "abcd", "bcda", "")

	b := map[string]any{
		"amount":   10,
		"currency": "usd",
		"order_id": "OrDeRiD",
	}
	d, err := json.Marshal(b)
	require.NoError(t, err)
	require.NotEmpty(t, d)

	sign := c.signRequest(c.paymentApiKey, d)
	require.Equal(t, "0f66e9e002a45d4285cc77a227dadfeb", sign)
}
