package main

import "fmt"
import "net/http"


func handler_home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home")
}

func handler_path1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "1111")
}

func handler_path2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "2222")
}



func main() {
	http.HandleFunc("/", handler_home)

	http.HandleFunc("/path1", handler_path1)

	http.HandleFunc("/path2", handler_path2)

	http.ListenAndServe(":3000", nil)
}