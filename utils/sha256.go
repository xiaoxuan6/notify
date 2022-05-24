package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func GenSignedByHmacSHA256(secret string, timestamp int64) string {

	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret

	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return ""
	}

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature
}
