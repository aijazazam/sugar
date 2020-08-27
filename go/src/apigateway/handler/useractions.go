package handler

import (
  //"fmt"
  "net/http"
  "io/ioutil"
  "crypto/tls"
)

func GetUserActions(w http.ResponseWriter, r *http.Request) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{ InsecureSkipVerify: true }

	username, _, err := DecodeToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
    return
	}

	req, err := http.NewRequest(http.MethodGet,"https://localhost:8003"+r.URL.Path, nil)
  if err != nil {
    panic( err )
  }

	q := req.URL.Query()
	q.Add("username", username)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

	blob, err := ioutil.ReadAll( resp.Body )
	if err != nil {
		panic(err)
	}

	w.Write( blob )
}

