package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var store = NewUrlStore()

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	key := store.Put(url)
	//store.Put(url)
	//http.Redirect(w, r, "/", http.StatusFound)
	fmt.Fprintf(w, "%s", key)
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./index/homepage.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	tmpl.Execute(w, nil)
}

//func Redirect(w http.ResponseWriter, r *http.Request) {
//	key := r.URL.Path[1:]
//	url := store.Get(key)
//	if url == "" {
//		http.NotFound(w, r)
//		return
//	}
//	http.Redirect(w, r, url, http.StatusFound)
//}

func main() {
	http.HandleFunc("/add", Add)
	http.HandleFunc("/", Homepage)
	//http.HandleFunc("/", Redirect)
	if err := http.ListenAndServe("0.0.0.0:45800", nil); err != nil {
		panic(err)
	}
}
