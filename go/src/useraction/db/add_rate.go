package db

import (
	_ "github.com/go-sql-driver/mysql"
)

type Rate struct {
  MovieId   uint32
  UserName  string
  Rating    float32
}

func (db *DB) AddRate( tRate Rate ) {

	sqldb := db.sqldb

	if _, err := sqldb.Query( `INSERT INTO Action(MovieId, UserName, Rating) VALUES(?,?,?) ON DUPLICATE KEY UPDATE Rating=?`, tRate.MovieId, tRate.UserName, tRate.Rating, tRate.Rating ); err != nil {
		panic( err )
	}
}

