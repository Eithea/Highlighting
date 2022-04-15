package db

import "src/secret"

type dbinfo struct {
	user string
	pw string
	url string
	engine string
	dbname string
}

var DB = dbinfo{"root", secret.PW, secret.ADDR, "mysql", "test"}