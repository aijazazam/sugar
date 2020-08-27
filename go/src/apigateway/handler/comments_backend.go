package handler

import (
	"fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "crypto/tls"
)

func comments_backend( movieIds []uint32 ) []byte {

  http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{ InsecureSkipVerify: true }

  req, err := http.NewRequest(http.MethodGet,"https://localhost:8003/comment/movie", nil)
  if err != nil {
    panic( err )
  }

  movieIds_string := strings.Trim(strings.Join(strings.Split(fmt.Sprint(movieIds), " "), ","), "[]")

  q := req.URL.Query()
  q.Add("movieid", movieIds_string)

  req.URL.RawQuery = q.Encode()

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll( resp.Body )

  return body
}


