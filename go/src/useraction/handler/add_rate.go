package handler

import (
	"net/http"
	//"io/ioutil"
	"encoding/json"

	sql "useraction/db"
)

type Rate struct {
	MovieId		uint32
	UserName	string
	Rating		float32
}

func AddRates(w http.ResponseWriter, r *http.Request) {

	rate := sql.Rate{}

	if err := json.NewDecoder(r.Body).Decode(&rate); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

	if rate.MovieId <= 0 || len(rate.UserName) == 0 {
		http.Error(w, "Invalid MovieId or UserName", http.StatusBadRequest)
    return
	}

	db := sql.GetDB()

	db.AddRate( rate )

	// Once a logged in used rates a movie that should effect the overall avg movie rating, but since we don't have the data of all users who rated a movie in pre-populated data,
	// Calculating the updated avg movie rating becomes difficult, So for simplicity we avoid updating the avg movie rating even if we hit our "/rate" api.

	return
}

