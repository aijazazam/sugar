package db

import (
	"fmt"
	"strings"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Movie struct {
	MovieId       uint32
  Name          string
  Genere        string
  Description   string
  Director      string
  Actors        string
  Year          uint32
  Duration      uint32
  Rating        float32
  Votes         uint32
  Revenue       float32
	UserRating		float32
	UserComment		string
}

func (db *DB) GetMoviesList( tMovieIds []uint32 ) []Movie {

	tResult	:= make( []Movie, 0 )
	sqldb		:= db.sqldb

	in := strings.Trim(strings.Join(strings.Split(fmt.Sprint(tMovieIds), " "), ","), "[]")

	query := fmt.Sprintf( `SELECT MovieId, Name, Genere, Description, Director, Actors, Year, Duration, Rating, Votes, Revenue
												 FROM Movie WHERE MovieId in (%s)`, in )

	row, err := sqldb.Query( query )

	if err != nil {
		panic( err )
	}

	for row.Next() {
		tM := Movie{}

		err := row.Scan( &tM.MovieId, &tM.Name, &tM.Genere, &tM.Description, &tM.Director,
	                   &tM.Actors, &tM.Year, &tM.Duration, &tM.Rating, &tM.Votes, &tM.Revenue )

		switch err {
			case sql.ErrNoRows:
				fmt.Println("No rows were returned!")
			case nil:
				//fmt.Println("Success")
			default:
				panic(err)
		}

		tResult = append( tResult, tM )
	}

	return tResult
}


