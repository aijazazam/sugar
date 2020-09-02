package handler

import (
	//"io"
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"encoding/json"
)

/*func Recommendation(w http.ResponseWriter, r *http.Request) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{ InsecureSkipVerify: true }

	resp, err := http.Get("https://localhost:8002/home")
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadAll( resp.Body )
	fmt.Println( string(data) )
	w.Write( data )

	return
}*/

func Recommendation(w http.ResponseWriter, r *http.Request) {

  movie_bytes := recommendation_backend( r )
fmt.Println(string(movie_bytes))
  var movie_body []interface{}
  if err := json.Unmarshal( movie_bytes, &movie_body ); err != nil {
    panic(err)
  }

  movieIds := extract_movies_list( movie_body )

  comments_bytes := comments_backend( movieIds )

	if comments_bytes != nil {
	  var comments_body []interface{}
		if err := json.Unmarshal(comments_bytes, &comments_body); err != nil {
			panic(err)
		}

		mComments := extract_comments( comments_body )

		movie_body = set_comments( movie_body, mComments )
	}

  bytes, err := json.Marshal( movie_body )
  if err != nil {
    panic(err)
  }

  w.Write( bytes )
}

func recommendation_backend( r *http.Request) []byte {

  http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{ InsecureSkipVerify: true }

  req, err := http.NewRequest(http.MethodGet,"https://localhost:8002"+r.URL.Path, nil)
  if err != nil {
    panic( err )
  }

  //req.URL.RawQuery = r.URL.Query().Encode()

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll( resp.Body )

  return body
}


