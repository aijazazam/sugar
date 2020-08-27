package db

import (
	_ "github.com/go-sql-driver/mysql"

	com "common"
)

func (db *DB) GetRecommend( tMinRating float32, tLimit int ) []com.Movie {

	sqldb := db.sqldb

	tRecommend := make( []com.Movie, 0 )

	results, err := sqldb.Query( `SELECT MovieId, Name, Genere, Description, Director, Actors, Year, Duration, Rating,
																Votes, Revenue FROM Movie WHERE Rating >= ? ORDER BY Rating DESC Limit ?`, tMinRating, tLimit )
	if err != nil {
		panic( err )
	}

	for results.Next() {

		tMovie := com.Movie{}

		if err = results.Scan( &tMovie.MovieId, &tMovie.Name, &tMovie.Genere, &tMovie.Description, &tMovie.Director,
														&tMovie.Actors, &tMovie.Year, &tMovie.Duration, &tMovie.Rating, &tMovie.Votes, &tMovie.Revenue ); err != nil {
			panic(err)
		}

		tRecommend = append( tRecommend, tMovie )
	}

	return tRecommend
}

