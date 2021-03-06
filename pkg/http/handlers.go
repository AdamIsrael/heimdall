package http

import (
	"encoding/json"
	"fmt"
	"github.com/adamisrael/heimdall/pkg/client"
	"github.com/adamisrael/heimdall/pkg/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"

	"html"
	"net/http"
	//    "strings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {
	var client client.APIClient

	_ = json.NewDecoder(req.Body).Decode(&client)

	if err := db.Where("username = ? and password = ?",
		client.Username, client.Password).First(&client).Error; err != nil {
		json.NewEncoder(w).Encode(Exception{Message: "Invalid credentials."})
		return
	} else {
		// Don't encode the password in the token
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
			"username": client.Username,
		})
		// TODO: move the secret to an external config file
		tokenString, error := token.SignedString([]byte("secret"))
		if error != nil {
			fmt.Println(error)
		}
		json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
	}

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
	var user client.APIClient
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// Here we are implementing the NotImplemented handler. Whenever an API endpoint is hit
// we will simply return the message "Not Implemented"
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

func UserIndex(w http.ResponseWriter, r *http.Request) {
	users := user.Users{
		user.User{Name: "Administrator"},
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
	fmt.Fprintf(w, "User id: %s", userid)
}

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "1.0.0")

}
