package http

import (
    "testing"
    "net/http/httptest"
    "net/http"
)

func TestWriteJson(t *testing.T) {
    rr := httptest.NewRecorder()
    WriteJson(rr, "Hello World")
    if http.StatusOK != rr.Code {
        t.Fail()
    }
}

func TestWriteJsonSetsContentType(t *testing.T) {
    rr := httptest.NewRecorder()
    WriteJson(rr, "Hello World")
    if "application/json; charset=utf-8" != rr.Header().Get("Content-Type") {
        t.Fail()
    }
}

func TestWriteJsonSendsMessage(t *testing.T) {
    rr := httptest.NewRecorder()
    WriteJson(rr, "Hello World")
    if "\"Hello World\"" != rr.Body.String() {
        t.Errorf("got \"%s\"",rr.Body.String())
        t.Fail()
    }
}

func TestWriteJsonSendsMessageWithInterface(t *testing.T) {
    rr := httptest.NewRecorder()
    msg := rr
    WriteJson(rr, msg)
    if "{\"Code\":200,\"HeaderMap\":{},\"Body\":{},\"Flushed\":false}" != rr.Body.String() {
        t.Errorf("got \"%s\"",rr.Body.String())
        t.Fail()
    }
}

func TestWriteEmptyJson(t *testing.T) {
    rr := httptest.NewRecorder()
    WriteEmptyJson(rr)
    if http.StatusOK != rr.Code {
        t.Fail()
    }
}

func TestWriteEmptyJsonSetsHMAC(t *testing.T) {
    rr := httptest.NewRecorder()
    WriteEmptyJson(rr)
    if "gBPtzLexSoIuVxap0j4hNLWxk24pBw6ZlpKQLwr8RVk=" != rr.Header().Get("HMAC") {
        t.Errorf("hmac: %s", rr.Header().Get("HMAC"))
        t.Fail()
    }
}

func TestWriteEmptyJsonSetsContentType(t *testing.T) {
    rr := httptest.NewRecorder()
    WriteEmptyJson(rr)
    if "application/json; charset=utf-8" != rr.Header().Get("Content-Type") {
        t.Fail()
    }
}

func TestWriteEmptyJsonSendsEmptyMessage(t *testing.T) {
    rr := httptest.NewRecorder()
    WriteEmptyJson(rr)
    if "[]" != rr.Body.String() {
        t.Errorf("got \"%s\"",rr.Body.String())
        t.Fail()
    }
}