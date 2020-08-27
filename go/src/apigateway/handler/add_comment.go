package handler

import (
  //"fmt"
	"bytes"
  "net/http"
  "io/ioutil"
  "crypto/tls"
  "encoding/json"
)

type Comment struct {
  MovieId   uint32
  UserName  string
  Comment   string
}

func AddComments(w http.ResponseWriter, r *http.Request) {

	comment := Comment{}

	username, _, err := DecodeToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
    return
	}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

	if comment.MovieId <= 0 || len(comment.Comment) == 0 {
    http.Error(w, "Invalid MovieId or empty Comment", http.StatusBadRequest)
    return
  }

	if len(comment.UserName) != 0 {
    http.Error(w, "You are not allowed to enter username, its picked up from jwt", http.StatusBadRequest)
    return
  }

	comment.UserName = username

	blob := add_comment_backend( comment, r )

	w.Write( blob )
}

func add_comment_backend( comment Comment, r *http.Request ) []byte {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{ InsecureSkipVerify: true }

	blob, err := json.Marshal( &comment )
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

