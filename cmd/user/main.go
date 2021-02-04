package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"

	"github.com/a-ovch/sa-course2020.09/pkg/app/user"
	"github.com/a-ovch/sa-course2020.09/pkg/app/user/transport"
	"github.com/a-ovch/sa-course2020.09/pkg/common/infrastructure/database"
	"github.com/a-ovch/sa-course2020.09/pkg/common/infrastructure/database/postgres"
)

const appName = "user"

type Config struct {
	Port       int `default:"8080"`
	DbUser     string
	DbPassword string
	DbHost     string `default:"localhost"`
	DbPort     int    `default:"5432"`
	DbName     string
}

func main() {
	var c Config
	err := envconfig.Process(appName, &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	client := connectToDb(c)
	defer client.Close()

	app := user.NewApplication(client)
	router := transport.NewRouter(app)

	startHttpServer(c, createHttpHandler(router))
}

func connectToDb(c Config) database.Client {
	const attempts = 5
	const timeout = 2 * time.Second

	dsn := postgres.NewDSN(c.DbHost, c.DbPort, c.DbUser, c.DbPassword, c.DbName)

	var db *sql.DB
	var err error
	for i := 0; i < attempts; i++ {
		log.Print("Try to connect to DB...")

		db, err = sql.Open("postgres", dsn.ToDSNString())
		if err == nil {
			break
		}
		time.Sleep(timeout)
	}

	if err != nil {
		log.Fatalf("Failed to connect to DB: %+v", err)
	}

	for i := 0; i < attempts; i++ {
		log.Print("DB is pinging...")
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(timeout)
	}

	if err != nil {
		_ = db.Close()
		log.Fatalf("DB ping failed %+v", err)
	}

	return database.NewClient(db)
}

func startHttpServer(c Config, h http.Handler) {
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(c.Port),
		Handler: h,
	}

	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		log.Printf("Starting HTTP server on port %d", c.Port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP Server error: %v\n", err)
		}
	}()

	sig := <-osSignalChan
	log.Printf("OS signal recieved: %+v", sig)

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %+v", err)
	}

	log.Print("Server successfully stopped")
}

func createHttpHandler(router *transport.Router) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/ready", readyHandler).Methods(http.MethodGet)

	ur := r.PathPrefix("/user").Subrouter()
	ur.HandleFunc("", router.PostUser).Methods(http.MethodPost)
	ur.HandleFunc("/{id}", router.GetUser).Methods(http.MethodGet)
	ur.HandleFunc("/{id}", router.DeleteUser).Methods(http.MethodDelete)
	ur.HandleFunc("/{id}", router.PutUser).Methods(http.MethodPut)

	r.Use(contentTypeJSONMiddleware, logMiddleware)
	return r
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

func contentTypeJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
