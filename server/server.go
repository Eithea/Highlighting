package main

import "fmt"
import "net/http"


func handler_home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home")
}

func handler_GET_static(w http.ResponseWriter, r *http.Request) {
	// do some static logic
	fmt.Fprint(w, "result of GET")
}

func handler_GET_dynamic(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		q = "nothing"
	}
	// do some dynamic logic
	fmt.Fprint(w, "result of GET\n")
	fmt.Fprint(w, "q is " + q)
}



func main() {
	http.HandleFunc("/", handler_home)

	http.HandleFunc("/GETtest_static", handler_GET_static)

	http.HandleFunc("/GETtest_dynamic", handler_GET_dynamic)

	http.ListenAndServe(":8080", nil)
}