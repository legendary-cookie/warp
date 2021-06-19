package main

import (
	"log"
	"net/http"
	"path/filepath"
	"fmt"
	"io/ioutil"
)

func serveRoot(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if (path == "/") {
		path = "index.html"
	}
	realpath := filepath.Join("/var/lib/warp", path)
	if (!Exists(realpath)) { 
			// TODO: 404.html
			fmt.Fprintf(w, "404 Not found\nPath: %v", path)	
			return;
	}
	dat, err := ioutil.ReadFile(realpath)
  if (err != nil) {
		fmt.Fprintf(w, "ERR WHILE READING FILE")
	}
  fmt.Fprintf(w, string(dat))
}

func main() {
	addr := "0.0.0.0:8080"
	http.HandleFunc("/", serveRoot)
	log.Fatal(http.ListenAndServe(addr, nil))
}
