package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	com "common"
)

func (db *DB) InsertMovie( tMovie com.Movie ) error {

	sqldb := db.sqldb

	fmt.Println("Kafka Message Poll, movie about to be inserted is")
	fmt.Println( tMovie )

	_, err := sqldb.Query( `REPLACE INTO Movie VALUES(?,?,?,?,?,?,?,?,?,?,?)`, tMovie.MovieId, tMovie.Name, tMovie.Genere, tMovie.Description, tMovie.Director,
													tMovie.Actors, tMovie.Year, tMovie.Duration, tMovie.Rating, tMovie.Votes, tMovie.Revenue )
	if err != nil {
		panic( err )
	}

	fmt.Println( "Recommend DB inserted ", tMovie.Name )

	return err
}

