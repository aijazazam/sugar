package handler

import (
  //"fmt"
	"bytes"
  "net/http"
  "io/ioutil"
  "crypto/tls"
  "encoding/json"

	com "common"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {

	movie := com.Movie{}

	_, isAdmin, err := DecodeToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
    return
	}

	if !isAdmin {
		http.Error(w, "Only admin can add movie", http.StatusUnauthorized)
    return
	}

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

	if movie.MovieId != 0 || len(movie.Name) == 0 {
    http.Error(w, "MovieId should not be entered and Name is mandatory", http.StatusBadRequest)
    return
  }

	blob := add_movie_backend( movie, r )

	w.Write( blob )
}

func add_movie_backend( movie com.Movie, r *http.Request ) []byte {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{ InsecureSkipVerify: true }

	blob, err := json.Marshal( &movie )
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

