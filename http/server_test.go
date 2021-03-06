package http

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "strings"
    "github.com/roundpartner/seq/buffer"
    "runtime"
)

func TestGet(t *testing.T) {
    rs := New(buffer.NewSimpleBuffer())
    rr := recordGet(t, rs)
    if rr.Code != http.StatusOK {
        t.Fail()
    }
}

func TestGetContentTypeIsJson(t *testing.T) {
    rs := New(buffer.NewSimpleBuffer())

    rr := recordGet(t, rs)
    if "application/json; charset=utf-8" != rr.Header().Get("Content-Type") {
        t.Fail()
    }
}

func TestGetReturnsEmptyJson(t *testing.T) {
    rs := New(buffer.NewSimpleBuffer())
    rr := recordGet(t, rs)
    if "[]" != rr.Body.String() {
        t.Fail()
    }
}

func TestGetSetsHMACHeader(t *testing.T) {
    rs := New(buffer.NewSimpleBuffer())
    rs.sb.Add("\"JSON Encoded Content\"")
    rr := recordGet(t, rs)
    if "" == rr.Header().Get("HMAC") {
        t.Fail()
    }
}

func TestGetReturnsMessage(t *testing.T) {
    rs := New(buffer.NewSimpleBuffer())
    rs.sb.Add("\"JSON Encoded Content\"")
    rr := recordGet(t, rs)
    if "[{\"id\":1,\"body\":\"JSON Encoded Content\"}]" != rr.Body.String() {
        t.Errorf("response: %s", rr.Body.String())
        t.Fail()
    }
}

func recordGet(t *testing.T, rs *RestServer) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }
    runtime.Gosched()
    rs.Get(rr, req)
    return rr
}

func TestPost(t *testing.T) {
    rs := New(buffer.NewSimpleBuffer())
    rr := recordPost(t, rs, "")
    if rr.Code != http.StatusNoContent {
        t.Fail()
    }
}

func TestPostAddsToBuffer(t *testing.T) {
    rs := New(buffer.NewSimpleBuffer())
    rs.sb = buffer.NewSimpleBuffer()
    recordPost(t, rs, "")
    runtime.Gosched()
    _, ok := buffer.Pop(rs.sb)
    //_, ok := rs.sb.Pop()
    if false == ok {
        t.Fail()
    }
}

func TestPostAddsMessageToBuffer(t *testing.T) {
    rs := New(buffer.NewSimpleBuffer())
    recordPost(t, rs, "Hello World")
    runtime.Gosched()
    message, _ := rs.sb.Pop()
    if "Hello World" != message {
        t.Fail()
    }
}

func recordPost(t *testing.T, rs *RestServer, body string) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    r := strings.NewReader(body)
    req, err := http.NewRequest("POST", "/", r)
    if err != nil {
        t.Fatal(err)
    }

    rs.Post(rr, req)
    return rr
}

func TestDelete(t *testing.T) {
    rr := httptest.NewRecorder()
    req, err := http.NewRequest("DELETE", "/1", nil)
    if err != nil {
        t.Fatal(err)
    }

    rs := New(buffer.NewSimpleBuffer())
    rs.router().ServeHTTP(rr, req)

    if rr.Code != http.StatusNotFound {
        t.Errorf("code: %d r: %s", rr.Code, rr.Body.String())
        t.Fail()
    }
}

func TestDeleteReturnsNoContent(t *testing.T) {
    rs := New(buffer.NewSimpleBuffer())

    rs.sb.Add("\"Delete Me\"")
    runtime.Gosched()
    c, _ := rs.clm.Next()
    rr := httptest.NewRecorder()
    req, err := http.NewRequest("DELETE", "/1", nil)
    if err != nil {
        t.Errorf("c: %d code: %d r: %s", c.Id, rr.Code, rr.Body.String())
        t.Fatal(err)
    }

    rs.router().ServeHTTP(rr, req)

    if rr.Code != http.StatusNoContent {
        t.Errorf("c: %d code: %d r: %s", c.Id, rr.Code, rr.Body.String())
        t.Fail()
    }
}