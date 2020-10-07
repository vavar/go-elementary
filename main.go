package main

import (
	"github.com/vavar/go-elementary/day2/captcha"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/vavar/go-elementary/fizzbuzz"
	_ "github.com/mattn/go-sqlite3"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func fizzBuzzHandler(w http.ResponseWriter, req *http.Request) {
	log.Println(req.RequestURI)
	res := fizzbuzz.FizzBuzz(3)
	io.WriteString(w, res)
}

func fizzBuzzGolliraHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	traceparent := req.Header.Get("traceparent")

	n, err := strconv.Atoi(vars["number"])
	if err != nil {
		log.Printf("ERROR: %s fizzBuzz handler %s", traceparent, err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error %s", err)
		return
	}

	log.Printf("INFO: we got the %d\n", n)
	res := fizzbuzz.FizzBuzz(n)
	io.WriteString(w, res)
}


func main() {

	viper.SetDefault("PORT", "8000")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("WARN: viper read config failed - %s\n", err)
	}

	viper.AutomaticEnv()

	db,err := sql.Open("sqlite3", "./captcha.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/fizzbuzz/", fizzBuzzHandler)
	r.HandleFunc("/fizzbuzz/{number}", fizzBuzzGolliraHandler)
	r.Handle("/captcha", captcha.NewCaptchaHandler(db)).Methods(http.MethodGet)
	r.Handle("/captcha", captcha.NewVerifyCaptchaHandler(db)).Methods(http.MethodPost)

	fmt.Printf("HTTP Server Started at %s\n", viper.GetString("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("PORT")), r))
}
