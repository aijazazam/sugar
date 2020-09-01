package handler

import (
	"net/http"
	//"io/ioutil"
	"encoding/json"

	kafka "useraction/kafka"
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

	avgRating := db.GetAvgMovieRating( rate.MovieId )

	db.UpdateMovieRating( avgRating )

	kafka.SendAvgRating <- avgRating

	return
}

