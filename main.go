package main

import (
	"github.com/vavar/go-elementary/apiproxy"
	"context"
	"os/signal"
	"syscall"
	"os"
	"github.com/vavar/go-elementary/captcha"
	"github.com/vavar/go-elementary/fizzbuzz"
	"github.com/vavar/go-elementary/auth"
	"github.com/vavar/go-elementary/trace"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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

func JSONContentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w,r)
	})
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

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	r := mux.NewRouter()
	r.Use(JSONContentMiddleware, trace.NewTraceMiddleware(logger))
	r.HandleFunc("/login", auth.LoginHandler)
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/api/users", apiproxy.UserHandler)
	r.Handle("/fizzbuzz/{number}", auth.SecureMiddleware(fizzBuzzGolliraHandler))
	r.Handle("/captcha", captcha.NewCaptchaHandler(db)).Methods(http.MethodGet)
	r.Handle("/verify", captcha.NewVerifyCaptchaHandler(db)).Methods(http.MethodPost)

	srv := http.Server {
		Addr: fmt.Sprintf(":%s", viper.GetString("PORT")),
		Handler: r,
	}

	fmt.Printf("HTTP Server Started at %s\n", viper.GetString("PORT"))
	log.Fatal(srv.ListenAndServe())

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
