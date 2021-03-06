package http

import (
    "testing"
    "net/http/httptest"
    "net/http"
)

func TestInternalError(t *testing.T) {
    rr := httptest.NewRecorder()
    InternalError(rr, "Testing Error")
    if http.StatusInternalServerError != rr.Code {
        t.Fail()
    }
}

func TestInternalErrorSetsJsonHeader(t *testing.T) {
    rr := httptest.NewRecorder()
    InternalError(rr, "Testing Error")
    if "application/json; charset=utf-8" != rr.Header().Get("Content-Type") {
        t.Fail()
    }
}

func TestInternalErrorSetsMessage(t *testing.T) {
    rr := httptest.NewRecorder()
    InternalError(rr, "Testing Error")
    if "{\"error\":\"Testing Error\"}" != rr.Body.String() {
        t.Errorf("got \"%s\"",rr.Body.String())
        t.Fail()
    }
}