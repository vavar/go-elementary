package captcha

import (
	"fmt"
	"encoding/json"
	"math/rand"
	"time"
	"net/http"
	"database/sql"

	"github.com/google/uuid"
)

//CaptchaHandler ...
type CaptchaHandler struct {
	db *sql.DB
}

//NewCaptchaHandler ...
func NewCaptchaHandler(db *sql.DB) CaptchaHandler {
	return CaptchaHandler{ db: db }
}

func (c CaptchaHandler) ServeHTTP( w http.ResponseWriter, req *http.Request) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	pattern := r.Intn(2) + 1
	leftOper := r.Intn(9) + 1
	oper := r.Intn(3) + 1
	rightOper := r.Intn(9) + 1

	ans := Answer(leftOper, oper, rightOper)
	ref:= uuid.New().String()
	
	_, err := c.db.Exec(`INSERT INTO captcha VALUES (?,?)`, ref, ans )
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
	}

	cc := NewCaptcha(pattern,leftOper,oper,rightOper)
	json.NewEncoder(w).Encode(map[string]string{
		"captcha": cc.String(),
		"ref": ref,
	})
}

type answerRequest struct {
	Ref string `json:"ref"`
	Answer int `json:"answer"`
}

//VerifyCaptchaHandler ...
type VerifyCaptchaHandler struct {
	db *sql.DB
}

//NewVerifyCaptchaHandler ...
func NewVerifyCaptchaHandler(db *sql.DB) VerifyCaptchaHandler {
	return VerifyCaptchaHandler{ db: db }
}

//AnswerHandler ...
func (v VerifyCaptchaHandler) ServeHTTP( w http.ResponseWriter, req *http.Request) {
	var payload answerRequest
	if err:= json.NewDecoder(req.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	defer req.Body.Close()
	
	row := v.db.QueryRow(`SELECT ans FROM captcha where ref=?`, payload.Ref)
	var correct int 
	if err := row.Scan(&correct); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Who are you",
		})
	}

	if correct != payload.Answer {
		fmt.Println(payload)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}