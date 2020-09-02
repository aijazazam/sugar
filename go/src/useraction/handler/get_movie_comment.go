package handler

import (
	"io"
	"fmt"
	"net/http"
	//"io/ioutil"
	"encoding/json"

	sql "useraction/db"
)

func GetMovieComments(w http.ResponseWriter, r *http.Request) {

	movieIds, ok := r.URL.Query()["movieid"]

	if len(movieIds) == 0 || !ok {
		http.Error(w, "list of movie requred", http.StatusBadRequest)
    return
	}

	if len(movieIds[0]) == 0 {
		return
	}

	db := sql.GetDB()
	movie_comments := db.GetMovieComments( movieIds[0] )

	blob, err := json.Marshal( movie_comments )
  if err != nil {
    fmt.Printf("Failed to Marchal json with err=", err)
    return
  }

  io.WriteString( w, string(blob) )

	return
}

