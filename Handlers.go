package main

import (
    "encoding/json"
    "fmt"

    "github.com/dgrijalva/jwt-go"
    "github.com/gorilla/context"
    "github.com/gorilla/mux"
    "github.com/mitchellh/mapstructure"

    "html"
    "net/http"
//    "strings"
)

type ApiUser struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type JwtToken struct {
    Token string `json:"token"`
}

type Exception struct {
    Message string `json:"message"`
}

func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {
    var user ApiUser
    _ = json.NewDecoder(req.Body).Decode(&user)
    
    // Don't encode the password in the token
    token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
        "username": user.Username,
    })
    // TODO: move the secret to an external config file
    tokenString, error := token.SignedString([]byte("secret"))
    if error != nil {
        fmt.Println(error)
    }
    json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        authorizationHeader := req.Header.Get("authorization")
        if authorizationHeader != "" {
            token, error := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("There was an error")
                }
                return []byte("secret"), nil
            })
            if error != nil {
                json.NewEncoder(w).Encode(Exception{Message: error.Error()})
                return
            }
            if token.Valid {
                fmt.Println("Token is valid")

                context.Set(req, "decoded", token.Claims)
                next(w, req)
            } else {
                json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
            }
        } else {
            json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization header."})
        }
    })
}

func TestEndpoint(w http.ResponseWriter, req *http.Request) {
    decoded := context.Get(req, "decoded")
    var user ApiUser
    mapstructure.Decode(decoded.(jwt.MapClaims), &user)
    json.NewEncoder(w).Encode(user)
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// Here we are implementing the NotImplemented handler. Whenever an API endpoint is hit
// we will simply return the message "Not Implemented"
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Not Implemented"))
})

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

func Version(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "1.0.0")

}
