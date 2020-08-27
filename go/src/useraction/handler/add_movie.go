package handler

import (
	"net/http"
	//"io/ioutil"
	"encoding/json"

	com "common"
	sql "useraction/db"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {

	movie := com.Movie{}

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

	if len(movie.Name) == 0 {
		http.Error(w, "Movie Name cannot be empty", http.StatusBadRequest)
    return
	}

	db := sql.GetDB()

	if err := db.AddMovie( movie ); err != nil {
		panic( err )
	}

	return
}

