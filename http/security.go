package http

import (
    "bytes"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
)

var hmackey = "hello world"

func getHMAC(js []byte) string {
    key := bytes.NewBufferString(hmackey).Bytes()
    mac := hmac.New(sha256.New, key)
    mac.Write(js)
    digest := mac.Sum(nil)
    return hex.EncodeToString(digest)
}