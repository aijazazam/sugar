package main

import (
	"os"
	//"fmt"
	"log"
	"time"
	"net/http"
	"io/ioutil"
	"crypto/tls"

	"search/kafka"
	"search/handler"
)

const (
	PUBLIC_CERT = "/src/common/cert/public.crt"
	PRIVATE_KEY = "/src/common/cert/private.key"
)

func main() {
	server := &http.Server{
		Addr:         ":8001",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    tlsConfig(),
	}

	go kafka.GoMovieConsumer(0)
	go kafka.GoRatingConsumer(0)

	http.HandleFunc("/search", handler.Search)

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

