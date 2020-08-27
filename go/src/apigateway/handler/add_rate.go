package handler

import (
  //"fmt"
	"bytes"
  "net/http"
  "io/ioutil"
  "crypto/tls"
  "encoding/json"
)

type Rate struct {
  MovieId   uint32
  UserName  string
  Rating    float32
}

func AddRates(w http.ResponseWriter, r *http.Request) {

	rate := Rate{}

	username, _, err := DecodeToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
    return
	}

	if err := json.NewDecoder(r.Body).Decode(&rate); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

	if rate.MovieId <= 0 || rate.Rating < 0.0 || rate.Rating > 10.0 {
    http.Error(w, "Invalid MovieId or Rating( 0.0 to 10.0 )", http.StatusBadRequest)
    return
  }

	if len(rate.UserName) != 0 {
    http.Error(w, "You are not allowed to enter username, its picked up from jwt", http.StatusBadRequest)
    return
  }

	rate.UserName = username

	blob := add_rate_backend( rate, r )

	w.Write( blob )
}

func add_rate_backend( rate Rate, r *http.Request ) []byte {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{ InsecureSkipVerify: true }

	blob, err := json.Marshal( &rate )
	if err != nil {
		panic( err )
	}

	req, err := http.NewRequest(http.MethodPost,"https://localhost:8003"+r.URL.Path, bytes.NewBuffer( blob ))
  if err != nil {
    panic( err )
  }

	req.URL.RawQuery = r.URL.Query().Encode()

	resp, err := http.DefaultClient.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

	blob, _ = ioutil.ReadAll( resp.Body )

	return blob
}

