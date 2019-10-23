package main

import (
	"log"
	"net/http"
	"path"
	"strings"
)

func main() {
	http.HandleFunc("/", serveFile)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("path: %s\n", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/")
	p = path.Clean(p)
	log.Printf("path: %s\n", p)
	if strings.HasPrefix(p, "..") {
		return
	}
	http.ServeFile(w, r, p)
}
