package main

import (
	"mime"
	"net/http"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	mime.ParseMediaType(contentType)

	w.Header().Set("Content-Type", "text/plain")
	// w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world!"))
	// w.Write([]byte("This is an example."))
}

func main() {
	http.ListenAndServe(":8000", nil)
}
