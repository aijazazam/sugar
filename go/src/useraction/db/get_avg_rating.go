package db

import (
	com "common"
	_ "github.com/go-sql-driver/mysql"
)

func (db *DB) GetAvgMovieRating( tMovieId uint32 ) com.AvgRating {

	all_ratings := make( []float32, 0 )

	sqldb := db.sqldb

	results, err := sqldb.Query( "SELECT Rating FROM Action WHERE MovieId=?", tMovieId )
	if err != nil {
		panic( err )
	}

	for results.Next() {

		var rating float32

		if err = results.Scan( &rating ); err != nil {
      panic(err)
    }

		all_ratings = append( all_ratings, rating )
	}

	avg_movie_rating := com.AvgRating{}

	avg_movie_rating.MovieId	= tMovieId
	avg_movie_rating.Votes		= (uint32) (len(all_ratings))
	avg_movie_rating.Rating		= 0

	for _, r := range all_ratings {
		avg_movie_rating.Rating += r
	}

	avg_movie_rating.Rating /= (float32) (avg_movie_rating.Votes)

	return avg_movie_rating
}

