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

var sb *buffer.SimpleBuffer = nil
var claims *claim.Elastic = nil
var clm *claim.C = nil

func Serve() {
    sb = buffer.New(1)
    claims = claim.New()
    clm = claim.NewC(claims, sb)

    http.ListenAndServe(":6060", router())
}

func router() *mux.Router {
    router := mux.NewRouter()
    rs := RestServer{}
    router.HandleFunc("/", rs.Get).Methods("GET")
    router.HandleFunc("/", rs.Post).Methods("POST")
    router.HandleFunc("/{id}", rs.Delete).Methods("DELETE")
    return router
}

type RestServer struct {

}

func (rs *RestServer) Get(w http.ResponseWriter, req *http.Request) {
    message, ok := clm.Next()
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

func (rs *RestServer) Post(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        InternalError(w, err.Error())
        return
    }
    res := sb.Add(string(body))
    if false == res {
        InternalError(w, "Buffer Full")
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func (rs *RestServer) Delete(w http.ResponseWriter, req *http.Request) {
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