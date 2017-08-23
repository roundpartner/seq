package http

import (
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
)

func Serve() {
    router := mux.NewRouter()
    router.HandleFunc("/", Get).Methods("GET")
    router.HandleFunc("/", Post).Methods("POST")
    http.ListenAndServe(":6060", router)
}

func Get(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "{}")
}

func Post(w http.ResponseWriter, req *http.Request) {
    w.WriteHeader(http.StatusNoContent)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
}