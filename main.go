package main

import (
	"log"
	"net/http"
)

func setContentType(contentType string, h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		h.ServeHTTP(w, r)
	}
}

func main() {
	handler := http.FileServer(http.Dir("data"))
	http.Handle("/", setContentType("video/mp4", handler))
	log.Printf("listening on port %s\n", ":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
