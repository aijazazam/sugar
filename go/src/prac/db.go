package main

import (
	"fmt"
  "database/sql"

	com "common"
  _ "github.com/go-sql-driver/mysql"
)

type DB struct {
  sqldb *sql.DB
}

var (
  db *DB = nil
)

func main() {

	db := GetDB()

  sqldb := db.sqldb

	results, err := sqldb.Query( `SELECT MovieId, Name, Genere, Description, Director, Actors, Year, Duration, Rating,
                                Votes, Revenue, CommentsCount FROM Movie WHERE Name Like ? ORDER BY Rating DESC Limit ?`, "ar?", 10 )
  if err != nil {
    panic( err )
  }

  for results.Next() {

    tMovie := com.Movie{}

    if err = results.Scan( &tMovie.MovieId, &tMovie.Name, &tMovie.Genere, &tMovie.Description, &tMovie.Director,
                            &tMovie.Actors, &tMovie.Year, &tMovie.Duration, &tMovie.Rating, &tMovie.Votes,
                            &tMovie.Revenue, &tMovie.CommentsCount ); err != nil {
      panic(err)
    }

		fmt.Println( tMovie )
	}

	fmt.Println("main shutdown")
}

func GetDB() (db *DB) {

  dbDriver  := "mysql"
  dbUser    := "root"
  dbPass    := "sugarbox"
  dbName    := "useraction"

  if( db != nil ) {
    return db
  }

  sqldb, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
  if err != nil {
    panic(err.Error())
  }

  db = &DB{ sqldb }

  return db
}

