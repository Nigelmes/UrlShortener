package main

import (
	"fmt"
	"net/http"
)

var store = NewUrlStore()

func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	url := r.FormValue("url")
	if url == "" {
		fmt.Fprint(w, AddForm)
		return
	}
	key := store.Put(url)

	fmt.Fprintf(w, "%s", key)
}

func Redirect(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/add", Add)
	http.HandleFunc("/", Redirect)
	if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
		panic(err)
	}
}
