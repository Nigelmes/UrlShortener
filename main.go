package main

import "net/http"

func Add(w http.ResponseWriter, r *http.Request) {

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
