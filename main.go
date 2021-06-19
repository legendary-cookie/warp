package main

import (
	"log"
	"net/http"
	"path/filepath"
	"fmt"
	"strings"
	"io/ioutil"
)

func serveRoot(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if (path == "/") {
		path = "index.html"
	}
	realpath := filepath.Join("/var/lib/warp", path)
	if (!Exists(realpath)) {
		if (!Exists(realpath+".html" )) {
			fmt.Fprintf(w, "<h1>404 Not found</h1>\n<p>Path: %v</p>", path)	
			return;
		} else {
			realpath = realpath + ".html"
	
		}
	}
	dat, err := ioutil.ReadFile(realpath)
  if (err != nil) {
		fmt.Fprintf(w, "ERR WHILE READING FILE")
	}
	if (strings.HasSuffix(realpath, ".css")) {
		w.Header().Set("Content-Type", "text/css")
	}
  fmt.Fprintf(w, string(dat))
}

func main() {
	addr := "0.0.0.0:8080"
	http.HandleFunc("/", serveRoot)
	log.Fatal(http.ListenAndServe(addr, nil))
}
