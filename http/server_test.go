package http

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "github.com/roundpartner/seq/buffer"
    "strings"
    "github.com/roundpartner/seq/claim"
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

func TestGetReturnsMessage(t *testing.T) {
    buf = buffer.Create(1)
    claims = claim.New()
    buffer.Add(buf, "Hello World")
    rr := recordGet(t)
    if "{\"id\":1,\"body\":\"Hello World\"}" != rr.Body.String() {
        t.Errorf("response: %s", rr.Body.String())
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
    rr := recordPost(t, "")
    if rr.Code != http.StatusNoContent {
        t.Fail()
    }
}

func TestPostAddsToBuffer(t *testing.T) {
    recordPost(t, "")
    _, ok := buffer.Pop(buf)
    if false == ok {
        t.Fail()
    }
}

func TestPostAddsMessageToBuffer(t *testing.T) {
    recordPost(t, "Hello World")
    message, _ := buffer.Pop(buf)
    if "Hello World" != message {
        t.Fail()
    }
}

func recordPost(t *testing.T, body string) *httptest.ResponseRecorder {
    buf = buffer.Create(1)
    claims = claim.New()
    rr := httptest.NewRecorder()
    r := strings.NewReader(body)
    req, err := http.NewRequest("POST", "/", r)
    if err != nil {
        t.Fatal(err)
    }
    Post(rr, req)
    return rr
}

func TestDelete(t *testing.T) {
    rr := httptest.NewRecorder()
    req, err := http.NewRequest("DELETE", "/1", nil)
    if err != nil {
        t.Fatal(err)
    }

    r := router()
    r.ServeHTTP(rr, req)

    if rr.Code != http.StatusNotFound {
        t.Errorf("code: %d r: %s", rr.Code, rr.Body.String())
        t.Fail()
    }
}

func TestDeleteReturnsNoContent(t *testing.T) {
    buf = buffer.Create(1)
    claims = claim.New()
    claim.NewC(claims, buf)
    buffer.Add(buf, "Hello World")
    c, _ := claim.Next(claims, buf)
    rr := httptest.NewRecorder()
    req, err := http.NewRequest("DELETE", "/1", nil)
    if err != nil {
        t.Errorf("c: %d code: %d r: %s", c.Id, rr.Code, rr.Body.String())
        t.Fatal(err)
    }

    r := router()
    r.ServeHTTP(rr, req)

    if rr.Code != http.StatusNoContent {
        t.Errorf("c: %d code: %d r: %s", c.Id, rr.Code, rr.Body.String())
        t.Fail()
    }
}