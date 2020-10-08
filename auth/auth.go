package auth 

import (
	"encoding/base64"
	"time"
	"fmt"
	"encoding/json"
	"net/http"
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

func LoginHandler(w http.ResponseWriter, req *http.Request) {
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
}

func SecureMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func( w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("authorization")
		if len(tokenString) <= 7 {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		tokenString = tokenString[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if t, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Println(t)
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret,nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w,r)
	})
}