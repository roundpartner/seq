package http

import (
    "testing"
    "bytes"
)

func TestGetBase64HMAC(t *testing.T) {
    payload := bytes.NewBufferString("hello world").Bytes()
    digest := getBase64HMAC(payload)
    if "o+DObPy4SD7mcLLw5YwfZhXAdC7B5eH0H7lfE9/oxak=" != digest {
        t.Errorf("digest does not match: %s", digest)
        t.Fail()
    }
}

func TestGetHexHMAC(t *testing.T) {
    payload := bytes.NewBufferString("hello world").Bytes()
    digest := getHexHMAC(payload)
    if "a3e0ce6cfcb8483ee670b2f0e58c1f6615c0742ec1e5e1f41fb95f13dfe8c5a9" != digest {
        t.Errorf("digest does not match: %s", digest)
        t.Fail()
    }
}