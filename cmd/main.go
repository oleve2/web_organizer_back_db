package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"webapp3/cmd/app"
	backendServ "webapp3/pkg/backend"

	"github.com/go-chi/chi"
)

const (
	defaultPort       = "9999"
	defaultHost       = "0.0.0.0"
	defaultSqlitePath = "../db/prd.db" /*запускаем из корня, поэтому путь именно такой*/
)

func main() {
	// read env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host, ok := os.LookupEnv("APP_HOST")
	if !ok {
		host = defaultHost
	}
	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = defaultPort
	}
	sqlitePath, ok := os.LookupEnv("APP_SQLITE_PATH")
	if !ok {
		sqlitePath = defaultSqlitePath
	}

	// start server
	if err := execute(net.JoinHostPort(host, port), sqlitePath); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func execute(addr string, sqlitePath string) error {
	// роутер
	mux := chi.NewRouter()

	// backend service
	backendSvc := backendServ.NewService(sqlitePath)

	// backend http server
	application := app.NewServer(
		mux,
		backendSvc,
	)

	// init app
	err := application.Init()
	if err != nil {
		log.Println(err)
		return err
	}
	server := &http.Server{
		Addr:    addr,
		Handler: application,
	}

	fmt.Printf("server started on http://%s\n", addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
