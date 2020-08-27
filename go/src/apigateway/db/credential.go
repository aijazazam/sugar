package db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	com "common"
)

func (db *DB) GetCredentials( tUserName string ) com.Credential {

  if ( len(tUserName) == 0 ) {
    panic("Empty UserName")
  }

  sqldb := db.sqldb

  row := sqldb.QueryRow( `SELECT UserName, PasswordHash, PasswordSalt, IsAdmin, IsDisabled FROM Authenticate WHERE UserName=?;`, tUserName )

  tUser := com.Credential{}

  switch err := row.Scan( &tUser.UserName, &tUser.PasswordHash, &tUser.PasswordSalt, &tUser.IsAdmin, &tUser.IsDisabled); err {
    case sql.ErrNoRows:
      fmt.Println("No rows were returned!")
    case nil:
      fmt.Println("Success", tUser)
    default:
      panic(err)
  }

	return tUser
}

func (db *DB) GoScanCredentials( cAuth chan<- com.Credential ) {

	sqldb := db.sqldb

	results, err := sqldb.Query( `SELECT UserName FROM Authenticate` )
	if err != nil {
		panic( err )
	}

	for results.Next() {

		tUserName := ""

		if err = results.Scan( &tUserName ); err != nil {
			panic(err)
		}

		cAuth <- com.Credential{ UserName: tUserName }
	}

	close(cAuth)
}

