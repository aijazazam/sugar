package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	com "common"
)

func (db *DB) UpdateMovieRating( tRate com.AvgRating ) error {

	sqldb := db.sqldb

	_, err := sqldb.Query( `Update Movie SET Rating=?,Votes=? WHERE MovieId=?`, tRate.Rating, tRate.Votes, tRate.MovieId )
	if err != nil {
		panic( err )
	}

	fmt.Printf( "Updated MovieId=%d with Rating=%v\n", tRate.MovieId, tRate.Rating )

	return err
}

