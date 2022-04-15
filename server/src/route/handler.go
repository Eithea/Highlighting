package route

import (
	"fmt"
	"net/http"
)

func Handler_home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home")
}

func Handler_GET_static(w http.ResponseWriter, r *http.Request) {
	// do some static logic
	fmt.Fprint(w, "result of GET")
}

func Handler_GET_dynamic(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		q = "nothing"
	}
	// do some dynamic logic
	fmt.Fprint(w, "result of GET\n")
	fmt.Fprint(w, "q is " + q)
}
