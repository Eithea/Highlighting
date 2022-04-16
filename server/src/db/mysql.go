package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Open() interface{} {
	q := testDB.user + ":" + testDB.pw + "@tcp(" + testDB.url + ")/" + testDB.dbname
	db, err := sql.Open(testDB.engine, q)
	if err != nil {
        panic(err)
    }
	return db
}

