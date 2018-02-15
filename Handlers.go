package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"

    "html"
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
    users := Users{
        User{Name: "Administrator"},
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(users); err != nil {
        panic(err)
    }

}

func UserShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userid := vars["id"]
    fmt.Fprintf(w, "User id:", userid)
}
