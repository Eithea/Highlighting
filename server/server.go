package main

import "fmt"
import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Main")
	})

	http.HandleFunc("/path1", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "path1")
	})

	http.HandleFunc("/path2", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "path2")
	})

	http.ListenAndServe(":3000", nil)
}