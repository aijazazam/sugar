package db

import (
	"fmt"
	"strings"
	"database/sql"

	com "common"
	_ "github.com/go-sql-driver/mysql"
)

type MovieMin struct {
	MovieId		uint32
	Name			string
}

func (db *DB) GoScanMovieName( cMovieMin chan<- MovieMin ) {

	sqldb := db.sqldb

  results, err := sqldb.Query( `SELECT MovieId, Name FROM Movie` )
  if err != nil {
    panic( err )
  }

  for results.Next() {

    tMovieMin := MovieMin{}

    if err = results.Scan( &tMovieMin.MovieId, &tMovieMin.Name ); err != nil {
      panic(err)
    }

		cMovieMin <- tMovieMin
  }

	close(cMovieMin)

  return
}

func (db *DB) GetSearch( tMovieIds []uint32 ) []com.Movie {

	tResult	:= make( []com.Movie, 0 )
	sqldb		:= db.sqldb

	in := strings.Trim(strings.Join(strings.Split(fmt.Sprint(tMovieIds), " "), ","), "[]")

	query := fmt.Sprintf( `SELECT MovieId, Name, Genere, Description, Director, Actors, Year,
												 Duration, Rating, Votes, Revenue FROM Movie WHERE MovieId in (%s)`, in )

	row, err := sqldb.Query( query )

	if err != nil {
		panic( err )
	}

	for row.Next() {
		tM := com.Movie{}

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


