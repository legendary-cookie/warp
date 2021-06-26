package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"os"
	"time"
)

var mimetypes Mimeconf = getMimeconf()
var conf Config = getConf()

var m map[string]string


func serveRoot(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if (path == "/") {
		path = "index.html"
	}
	realpath := filepath.Join("/var/lib/warp", path)
	fi, err := os.Stat(realpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(realpath)
	fmt.Println(fi.Mode().IsDir())
	if fi.Mode().IsDir() {
		files, err := ioutil.ReadDir(realpath)
		if err != nil {
			fmt.Println(err)
		}
		filehtml := "<!DOCTYPE html>\n<head>\n\t<title>"+path+"</title>\n</head><body><div><h3>Files and Directories in "+path+"</h3></div>"
		for _, file := range files {
			filehtml = filehtml+"<div><a href=\""+path+"/"+file.Name()+"\">"+file.Name()+"</a></div>"
		}
		filehtml = filehtml+"</body>"
		fmt.Fprintf(w, filehtml)
		return
	}
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
	if conf.Templates {
		if !strings.Contains(string(dat), "###WARP_EXCLUDE") {
			year, month, _ := time.Now().Date()
			dat = []byte(strings.ReplaceAll(string(dat), "{{YEAR}}", strconv.Itoa(year)))
			dat = []byte(strings.ReplaceAll(string(dat), "{{MONTH}}", month.String()))
		}
			dat = []byte(strings.ReplaceAll(string(dat), "###WARP_EXCLUDE", ""))
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
	fmt.Println("Starting warp webserver...")
	m = make(map[string]string)
	for _, s := range mimetypes.Mime {
		split := strings.Split(s, "|")
		m[split[0]] = split[1]
	}
	addr := conf.Address
	port := conf.Port
	address := addr + ":" + port
	fmt.Println("Listening on: " +address)
	http.HandleFunc("/", serveRoot)
	log.Fatal(http.ListenAndServe(address, nil))
}
