package db

import (
	"fmt"
	"database/sql"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"src/secret"
)

func Dbtest(q string, w http.ResponseWriter) {
	conn := "root:" + secret.PW + "@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", conn)
	if err != nil {
        panic(err)
    }
	defer db.Close()

	_, err = db.Exec("insert into man(name, addr) values(?, ?)", q, "서울")
	if err != nil {
        panic(err)
    }

	rows, err := db.Query("select name, addr from man")
	if err != nil {
        panic(err)
    }
	for rows.Next() {
		var name string
		var addr string
		rows.Scan(&name, &addr)
		fmt.Fprint(w, "name : " + name + " addr : " + addr + "\n")
	}
	db.Close()
}

