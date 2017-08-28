package http

import (
    "github.com/gorilla/mux"
    "net/http"
    "github.com/roundpartner/seq/buffer"
    "io/ioutil"
    "encoding/json"
    "github.com/roundpartner/seq/claim"
    "strconv"
)

var buf chan buffer.Message = nil
var claims *claim.Elastic = nil

func Serve() {
    sb := buffer.New(1)
    buf = sb.Messages
    claims = claim.New()

    http.ListenAndServe(":6060", router())
}

func router() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/", Get).Methods("GET")
    router.HandleFunc("/", Post).Methods("POST")
    router.HandleFunc("/{id}", Delete).Methods("DELETE")
    return router
}

func Get(w http.ResponseWriter, req *http.Request) {
    message, ok := claim.Next(claims, buf)
    if false == ok {
        WriteEmptyJson(w)
        return
    }
    js, err := json.Marshal(message)
    if err != nil {
        InternalError(w, err.Error())
        return
    }
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

func Delete(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id, err := strconv.ParseInt(params["id"], 10, 32)

    if err != nil {
        InternalError(w, err.Error())
        return
    }
    qry := claim.Query{Id: int(id), Out: make(chan claim.Item), Delete: true}
    claims.Query <- qry
    ec := claim.Item{}
    if ec == <- qry.Out {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}