package db

import (
	"fmt"
	"useraction/kafka"

	com "common"
	_ "github.com/go-sql-driver/mysql"
)

func (db *DB) AddMovie( m com.Movie ) error {

	sqldb := db.sqldb

	stmt, err := sqldb.Prepare( `INSERT INTO Movie(Name, Genere, Description, Director, Actors, Year,
															 Duration, Rating, Votes, Revenue ) VALUES(?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		panic( err )
	}

	res, err := stmt.Exec( m.Name, m.Genere, m.Description, m.Director, m.Actors, m.Year, m.Duration, m.Rating, m.Votes, m.Revenue )
  if err != nil {
    panic( err )
  }

	lastInserted, err := res.LastInsertId()
  if err != nil {
    panic( err )
  }

	m.MovieId = uint32(lastInserted)

	kafka.SendMovie <- m

	fmt.Println( "useraction Movie inserted ", m )

	return err
}

