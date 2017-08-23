package http

import (
    "testing"
    "net/http"
    "net/http/httptest"
)

func TestGet(t *testing.T) {
    rr := recordGet(t)
    if rr.Code != http.StatusOK {
        t.Fail()
    }
}

func TestGetContentTypeIsJson(t *testing.T) {
    rr := recordGet(t)
    if "application/json; charset=utf-8" != rr.Header().Get("Content-Type") {
        t.Fail()
    }
}

func TestGetReturnsEmptyJson(t *testing.T) {
    rr := recordGet(t)
    if "{}" != rr.Body.String() {
        t.Fail()
    }
}

func recordGet(t *testing.T) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }
    Get(rr, req)
    return rr
}

func TestPost(t *testing.T) {
    rr := recordPost(t)
    if rr.Code != http.StatusNoContent {
        t.Fail()
    }
}

func recordPost(t *testing.T) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    req, err := http.NewRequest("POST", "/", nil)
    if err != nil {
        t.Fatal(err)
    }
    Post(rr, req)
    return rr
}