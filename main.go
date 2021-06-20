package main

import (
	"log"
	"net/http"
	"path/filepath"
	"fmt"
	"strings"
	"io/ioutil"
)

var mimetypes Mimeconf = getMimeconf()


var m map[string]string



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
	if (!strings.HasSuffix(realpath, ".html")) {
		base := filepath.Base(realpath)
		ext := strings.ReplaceAll(filepath.Ext(base), ".", "")
		if val, ok := m[ext]; ok {
			w.Header().Set("Content-Type", val)
		}
	}
  fmt.Fprintf(w, string(dat))
}

func main() {
	m = make(map[string]string)
	for _, s := range mimetypes.Mime {
		split := strings.Split(s, "|")
		m[split[0]] = split[1]
	}
	conf := getConf()
	addr := conf.Address
	port := conf.Port
	address := addr + ":" + port
	http.HandleFunc("/", serveRoot)
	log.Fatal(http.ListenAndServe(address, nil))
}
