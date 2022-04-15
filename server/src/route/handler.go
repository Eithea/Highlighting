package route

import (
	"fmt"
	"net/http"
	"src/json"
)

func Handler_home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(json.Struct_testdata("home", 111)))
}

func Handler_GET_static(w http.ResponseWriter, r *http.Request) {
	// do some static logic
	fmt.Fprint(w, string(json.Struct_testdata("test", 123)))
}

func Handler_GET_dynamic(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("query")
	if q == "" {
		q = "nothing"
	}
	// do some dynamic logic
	fmt.Fprint(w, string(json.Struct_testdata(q, 123)))
}
