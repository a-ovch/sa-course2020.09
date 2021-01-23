package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"github.com/kelseyhightower/envconfig"
)

const appName = "user"

type Config struct {
	Port int `default:"80"`
}

func main() {
	var c Config
	err := envconfig.Process(appName, &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/ready", readyHandler)
	r.Use(logMiddleware)

	log.Printf("Starting HTTP server on port %d", c.Port)

	// Add graceful shutdown
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(c.Port),
		Handler: r,
	}
	log.Fatal(srv.ListenAndServe())
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "{\"status\": \"OK\"}")
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "{\"host\": \"%v\"}", r.Host)
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
