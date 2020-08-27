package handler

import (
	"io"
	"fmt"
	"net/http"
	//"io/ioutil"
	"encoding/json"

	sql "useraction/db"
)

var EMPTY struct{}

func GetUserActions(w http.ResponseWriter, r *http.Request) {

	username, ok := r.URL.Query()["username"]

	if len(username) == 0 || !ok {
		http.Error(w, "list of movie requred", http.StatusBadRequest)
    return
	}

	db := sql.GetDB()
	user_actions := db.GetUserActions( username[0] )

	user_action_map := make( map[uint32]sql.Action )
	movieId_list := make( []uint32, 0 )

	for _, ua := range user_actions {
		movieId_list = append( movieId_list, ua.MovieId )
		user_action_map[ua.MovieId] = ua
	}

	movies := db.GetMoviesList( movieId_list )

	for i, movie := range movies {
		if action, ok := user_action_map[movie.MovieId]; ok {
			movies[i].UserRating	=	action.Rating
			movies[i].UserComment = action.Comment
		}
	}

	blob, err := json.Marshal( movies )
  if err != nil {
    fmt.Printf("Failed to Marchal json with err=", err)
    return
  }

  io.WriteString( w, string(blob) )

	return
}

func appendComments( ) {
}


