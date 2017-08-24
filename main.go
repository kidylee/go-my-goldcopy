package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var raddr = flag.String("raddr", "localhost:6379", "redis server address")
var log = logrus.New()

func init() {

	flag.Parse()

	// Log as JSON instead of the default ASCII formatter.
	log.Formatter = new(logrus.JSONFormatter)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.Out = os.Stdout

	// Only log the warning severity or above.
	log.Level = logrus.DebugLevel
}

func main() {

	//router init
	router := mux.NewRouter()
	router.HandleFunc("/", echo)
	router.HandleFunc("/category/{key}", getCategory).Methods("GET")
	router.HandleFunc("/category", createCategory).Methods("POST")

	//middleware init
	n := negroni.Classic()
	n.UseHandler(router)

	//server start
	log.Fatal(http.ListenAndServe(*addr, n))
}
