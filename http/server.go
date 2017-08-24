package http

import (
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
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
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.WriteHeader(http.StatusOK)
        fmt.Fprint(w, "{}")
        return
    }
    js, err := json.Marshal(message)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func Post(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    body, error := ioutil.ReadAll(req.Body)
    if error != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    res := buffer.Add(buf, string(body))
    if false == res {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}