package http

import (
    "net/http"
    "encoding/json"
    "bytes"
)

func WriteJson(w http.ResponseWriter, output interface{}) {
    js, err := json.Marshal(output)
    if err != nil {
        js = bytes.NewBufferString("{\"error\":\"Marshal Error\"}").Bytes()
        writeJsonError(w, js, http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write(js)
}

func WriteEmptyJson(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    js := bytes.NewBufferString("[]").Bytes()
    w.Write(js)
}