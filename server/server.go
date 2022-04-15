package main

import (
	"net/http"
	"src/route"
)


func main() {
	http.HandleFunc("/", route.Handler_home)

	http.HandleFunc("/GETtest_static", route.Handler_GET_static)

	http.HandleFunc("/GETtest_dynamic", route.Handler_GET_dynamic)

	http.ListenAndServe(":8080", nil)
}