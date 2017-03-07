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
	fmt.Fprint(w, scrape.Scrape("http://www.drudgereport.com"))
}

func main() {
	http.HandleFunc("/status", status)
	http.HandleFunc("/scrape", scraper)
	http.ListenAndServe(":8080", nil)
}
