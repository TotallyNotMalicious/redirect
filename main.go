package main

import (
	"log"
	"net/http"
)

func main() {

	redirect := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://example.com", 307) // change to the url that you would like to redirect to
	}

	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(":80", nil))
}
