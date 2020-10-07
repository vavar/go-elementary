package main 

import (
	"encoding/base64"
	"time"
	"fmt"
	"encoding/json"
	"net/http"
	"log"

	"github.com/gorilla/mux"
	jwt "github.com/dgrijalva/jwt-go"
)

type Credential struct {
	Email	string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

var hmacSampleSecret = []byte("ureehhh")

func loginHandler(w http.ResponseWriter, req *http.Request) {
	var cred Credential
	if err:= json.NewDecoder(req.Body).Decode(&cred); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	defer req.Body.Close()

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add( 3* time.Minute).Unix(),
		"email":base64.StdEncoding.EncodeToString([]byte(cred.Email)),
	}).SignedString(hmacSampleSecret)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "unknow error",
		})
		return
	}

	resp := AuthResponse{ Token: tokenString }
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("done")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler)

	log.Fatal(http.ListenAndServe(":8080",r))
}