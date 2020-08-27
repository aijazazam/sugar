package handler

import (
	"net/http"
	//"io/ioutil"
	"encoding/json"

	sql "useraction/db"
)

/*type Comment struct {
	MovieId		uint32
	UserName	string
	Comment		string
}*/

func AddComments(w http.ResponseWriter, r *http.Request) {

	comment := sql.Comment{}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

	if comment.MovieId <= 0 || len(comment.UserName) == 0 {
		http.Error(w, "Invalid MovieId or UserName", http.StatusBadRequest)
    return
	}

	db := sql.GetDB()

	db.AddComment( comment )

	return
}

