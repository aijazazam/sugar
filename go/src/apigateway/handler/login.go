package handler

import (
	"io"
	//"fmt"
	"net/http"
	//"io/ioutil"
	"encoding/json"

	auth "apigateway/authenticate"
)

type Credentials struct {
	UserName	string
	Password	string
}

func Login(w http.ResponseWriter, r *http.Request) {

	c := Credentials{}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(c.UserName) == 0 || len(c.Password) == 0 {
		w.WriteHeader( http.StatusUnprocessableEntity )
		io.WriteString(w, `{"error":"empty_credentials, please enter both username and password"}`)
		return
	}

	isAdmin, isAuthenticated := auth.Authentication( c.UserName, c.Password )

	if !isAuthenticated {
		w.WriteHeader(http.StatusUnauthorized)
    io.WriteString(w, `{"error":"Authentication Failed"}`)
    return
	}

	token, err := CreateToken( c.UserName, isAdmin )
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"token_generation_failed"}`)
		return
	}

	io.WriteString(w, `{"jwt":"`+token+`"}`)
	//w.Write([]byte(fmt.Sprintf("Authenticated and Successfully generated jwt: %s", tokenString)))

	return
}

