package main

import (
	"os"
	//"fmt"
	"log"
	"time"
	"net/http"
	"io/ioutil"
	"crypto/tls"

	kafka "useraction/kafka"
	"useraction/handler"
)

const (
	PUBLIC_CERT = "/src/common/cert/public.crt"
	PRIVATE_KEY = "/src/common/cert/private.key"
)

// http://www.inanzzz.com/index.php/post/9ats/http2-and-tls-client-and-server-example-with-golang
func main() {

	go kafka.GoProducer()

	server := &http.Server{
		Addr:         ":8003",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    tlsConfig(),
	}

	// GET
	http.HandleFunc("/useraction",	handler.GetUserActions)
	http.HandleFunc("/comment/movie",	handler.GetMovieComments)

	// POST
	http.HandleFunc("/add/rate", handler.AddRates)
	http.HandleFunc("/add/comment", handler.AddComments)
	http.HandleFunc("/add/movie", handler.AddMovie)

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}

func tlsConfig() *tls.Config {

	certFile := os.Getenv("GOPATH") + PUBLIC_CERT
	privateKey := os.Getenv("GOPATH") + PRIVATE_KEY

	crt, err := ioutil.ReadFile( certFile )
	if err != nil {
		log.Fatal(err)
	}

	key, err := ioutil.ReadFile( privateKey )
	if err != nil {
		log.Fatal(err)
	}

	cert, err := tls.X509KeyPair(crt, key)
	if err != nil {
		log.Fatal(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
	}
}

