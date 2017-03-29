package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	addr     = flag.String("addr", ":8080", "http service address")
	filename string
)

func readFileIfModified(lastMod time.Time) ([]byte, time.Time, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, lastMod, err
	}
	if !fi.ModTime().After(lastMod) {
		return nil, lastMod, nil
	}
	p, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fi.ModTime(), err
	}
	return p, fi.ModTime(), nil
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p, lastMod, err := readFileIfModified(time.Time{})
	if err != nil {
		p = []byte(err.Error())
		lastMod = time.Unix(0, 0)
	}

	var v = struct {
		Host      string
		Data      string
		LastMod   string
		Artifacts Artifacts
	}{
		r.Host,
		string(p),
		lastMod.Format("15:04:05 Mon Jan _2 2006"),
		newArtifacts(p),
	}

	err = homeTempl.Execute(w, &v)
	if err != nil {
		http.Error(w, "Failed to execute template", 500)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("filename not specified")
	}
	filename = flag.Args()[0]
	http.HandleFunc("/", serveHome)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
