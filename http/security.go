package http

import (
    "bytes"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "encoding/base64"
)

var hmackey = "an example key"

func getHexHMAC(js []byte) string {
    return hex.EncodeToString(getHMAC(js))
}

func getBase64HMAC(js []byte) string {
    return base64.StdEncoding.EncodeToString(getHMAC(js))
}

func getHMAC(js []byte) []byte {
    key := bytes.NewBufferString(hmackey).Bytes()
    mac := hmac.New(sha256.New, key)
    mac.Write(js)
    digest := mac.Sum(nil)
    return digest
}