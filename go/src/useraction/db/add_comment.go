package db

import (
	_ "github.com/go-sql-driver/mysql"
)

type Comment struct {
  MovieId   uint32
  UserName  string
  Comment   string
}

func (db *DB) AddComment( tCom Comment ) {

	sqldb := db.sqldb

	if _, err := sqldb.Query( `INSERT INTO Action(MovieId, UserName, Comment) VALUES(?,?,?) ON DUPLICATE KEY UPDATE Comment=?`,
														 tCom.MovieId, tCom.UserName, tCom.Comment, tCom.Comment ); err != nil {
		panic( err )
	}
}

