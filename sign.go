package gocryptomus

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func (c *Cryptomus) signRequest(apiKey string, reqBody []byte) string {
	data := base64.StdEncoding.EncodeToString(reqBody)
	hash := md5.Sum([]byte(data + apiKey))
	return hex.EncodeToString(hash[:])
}
