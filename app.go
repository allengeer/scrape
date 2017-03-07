package main

import (
	"fmt"
	"net/http"
	"github.com/allengeer/scrape/lib"
)

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"RUNNING")
}
func scraper(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("url")
	if url == "" {
		r.ParseForm()
		url = r.FormValue("url")
	}
	fmt.Fprint(w, scrape.Scrape(url))
}

func main() {
	http.HandleFunc("/status", status)
	http.HandleFunc("/scrape", scraper)
	http.ListenAndServe(":5000", nil)
}
