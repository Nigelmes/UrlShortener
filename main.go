package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var store = NewUrlStore()

func Homepage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./index/homepage.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	if r.Method == "POST" {
		url := r.FormValue("url")
		if url != "" {
			key := store.Put(url)
			tmpl.Execute(w, "http://localhost/short/"+key)
			return
		}
	}
	tmpl.Execute(w, "")
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/short/"):]
	url := store.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func main() {
	http.Handle("/index/", http.StripPrefix("/index/", http.FileServer(http.Dir("./index/"))))
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/short/", Redirect)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
