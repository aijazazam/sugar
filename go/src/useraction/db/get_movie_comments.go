package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func (db *DB) GetMovieComments( tMovies string ) []Comment {

	allComments := make( []Comment, 0 )

	sqldb := db.sqldb

	query := fmt.Sprintf( "SELECT MovieId, UserName, Comment FROM Action WHERE Comment IS NOT NULL AND MovieId in (%s);", tMovies )

	results, err := sqldb.Query( query )
	if err != nil {
		panic( err )
	}

	for results.Next() {

		comment := Comment{}

		if err = results.Scan( &comment.MovieId, &comment.UserName, &comment.Comment ); err != nil {
      panic(err)
    }

		allComments = append( allComments, comment )
	}

	return allComments
}

