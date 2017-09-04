package http

import (
    "testing"
    "bytes"
)

func TestGetHMAC(t *testing.T) {
    payload := bytes.NewBufferString("hello world").Bytes()
    digest := getHMAC(payload)
    if "6ec035d91dc104db569a01a4d8c16fb13f125dc298992edfb8e66d3a837fe0c5" != digest {
        t.Errorf("digest does not match: %s", digest)
        t.Fail()
    }

}