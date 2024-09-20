package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello World")
// 	})

// 	http.ListenAndServe(":3000", nil)
// }
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "Wolrd"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}

//	func main() {
//		http.HandleFunc("/bar", barHandler)
//		http.ListenAndServe(":3000", nil)
//	}
// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello world")
// 	})
// 	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello Bar")
// 	})

//		http.ListenAndServe(":3000", mux)
//	}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

type Student struct {
	Name  string
	Age   int
	Score int
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa", 16, 87}
	data, _ := json.Marshal(student)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
