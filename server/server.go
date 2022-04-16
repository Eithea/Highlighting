// package main

// import (
// 	"net/http"
// 	"src/route"
// )


// func main() {
// 	http.HandleFunc("/", route.Home)
// 	http.HandleFunc("/t1", route.Test1)
// 	http.HandleFunc("/t2", route.Test2)

// 	http.ListenAndServe(":8080", nil)
// }


package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"src/secret"
)

func main() {
	conn := "root:" + secret.PW + "@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", conn)
	if err != nil {
        panic(err)
    }
	defer db.Close()

	res, err := db.Exec("insert into man(name, addr) values(?, ?)", "김뉴비", "서울")
	if err != nil {
        panic(err)
    }
	n, err := res.RowsAffected()
    if n == 1 {
        fmt.Println("1 row inserted.")
    }

	rows, err := db.Query("select name, addr from man")
	if err != nil {
        panic(err)
    }
	for rows.Next() {
		var name string
		var addr string
		rows.Scan(&name, &addr)
		fmt.Printf("name %s, addr %s\n", name, addr)
	}

	db.Close()
}

