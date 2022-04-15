package main

import (
	"net/http"
	"src/route"
)


func main() {
	http.HandleFunc("/", route.Home)
	http.HandleFunc("/t1", route.Test1)
	http.HandleFunc("/t2", route.Test2)

	http.ListenAndServe(":8080", nil)
}