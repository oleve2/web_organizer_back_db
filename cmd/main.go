package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"webapp3/cmd/app"
	backendServ "webapp3/pkg/backend"
	updownServ "webapp3/pkg/updownload"

	"github.com/go-chi/chi"
)

const (
	defaultPort       = "9999"
	defaultHost       = "0.0.0.0"
	defaultSqlitePath = "../db/prd.db" /*запускаем из корня, поэтому путь именно такой*/
	//
	defaultStaticFilePath = "./static/files" // путь к папке раздачи статики
	defaultStaticURL      = "/f"             // url
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

	// upload/download service
	updownServ := updownServ.NewService(defaultStaticFilePath, defaultStaticURL)
	updownServ.InitCheck()

	// dir for serving static files
	filesDir := http.Dir(defaultStaticFilePath)
	FileServer(mux, defaultStaticURL, filesDir) // параметры в const

	// backend http server
	application := app.NewServer(
		mux,
		backendSvc,
		updownServ,
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

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
