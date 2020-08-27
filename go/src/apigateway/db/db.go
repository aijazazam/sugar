package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	sqldb *sql.DB
}

var (
	db *DB = nil
)

func GetDB() (db *DB) {

	dbDriver	:= "mysql"
	dbUser		:= "root"
	dbPass		:= "sugarbox"
	dbName		:= "apigateway"

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

func (db *DB) CloseDB() {

	if( db != nil ) {
		db.sqldb.Close()
		db.sqldb = nil
		db = nil
	}
}

