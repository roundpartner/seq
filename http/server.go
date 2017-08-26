package http

import (
    "github.com/gorilla/mux"
    "net/http"
    "github.com/roundpartner/seq/buffer"
    "io/ioutil"
    "encoding/json"
)

var buf chan buffer.Message = nil

func Serve() {
    buf = buffer.Create(1)

    router := mux.NewRouter()
    router.HandleFunc("/", Get).Methods("GET")
    router.HandleFunc("/", Post).Methods("POST")
    http.ListenAndServe(":6060", router)
}

func Get(w http.ResponseWriter, req *http.Request) {
    message, ok := buffer.Pop(buf)
    if false == ok {
        WriteEmptyJson(w)
        return
    }
    js, err := json.Marshal(message)
    if err != nil {
        InternalError(w, err.Error())
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func Post(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        InternalError(w, err.Error())
        return
    }
    res := buffer.Add(buf, string(body))
    if false == res {
        InternalError(w, "Buffer Full")
        return
    }
    w.WriteHeader(http.StatusNoContent)
}