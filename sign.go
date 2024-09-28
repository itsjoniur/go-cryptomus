package cryptomus

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
)

func (c *Cryptomus) signRequest(apiKey string, reqBody []byte) string {
	data := base64.StdEncoding.EncodeToString(reqBody)
	hash := md5.Sum([]byte(data + apiKey))
	return hex.EncodeToString(hash[:])
}

func (c *Cryptomus) VerifySign(apiKey string, reqBody []byte) error {
	var jsonBody map[string]any
	err := json.Unmarshal(reqBody, &jsonBody)
	if err != nil {
		return err
	}

	reqSign, ok := jsonBody["sign"].(string)
	if !ok {
		return errors.New("missing signature field in request body")
	}
	delete(jsonBody, "sign")

	expectedSign := c.signRequest(apiKey, reqBody)
	if reqSign != expectedSign {
		return errors.New("invalid signature")
	}
	return nil
}
