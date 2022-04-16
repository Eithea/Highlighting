package route

import (
	"fmt"
	"net/http"
	"src/json"
	"src/db"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(json.Struct_test1("home", 111)))
}

func Test1(w http.ResponseWriter, r *http.Request) {
	// do some static logic
	fmt.Fprint(w, string(json.Struct_test1("test", 123)))
}

func Test2(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("query")
	if q == "" {
		q = "nothing"
	}
	// do some dynamic logic
	db.Dbtest(q, w)
	fmt.Fprint(w, string(json.Struct_test2(q, 123)))
}