package db

import (
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Action struct {
	MovieId		uint32
	UserName	string
	Rating		float32
	Comment		string
}

func (db *DB) GetUserActions( tUserName string ) []Action {

	allUserActions := make( []Action, 0 )

	sqldb := db.sqldb

	//query := fmt.Sprintf( "SELECT MovieId, UserName, Rating, Comment FROM Action WHERE UserName='%s'", tUserName )

	results, err := sqldb.Query( `SELECT MovieId, UserName, Rating, IFNULL(Comment, "") FROM Action WHERE UserName=?`, tUserName )
	if err != nil {
		panic( err )
	}

	for results.Next() {

		ac := Action{}

		if err = results.Scan( &ac.MovieId, &ac.UserName, &ac.Rating, &ac.Comment ); err != nil {
      panic(err)
    }

		allUserActions = append( allUserActions, ac )
	}

	return allUserActions
}

