// assets
package main

import (
	//"fmt"
	"html/template"
	//"log"
	"net/http"
	"strings"
)

var staticMap map[string]string

type Mux struct {
}

func AddstaticMap(webdir, localdir string) {
	staticMap[webdir] = localdir
}

func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sli := strings.Split(r.URL.Path, "/")
	prefix := "/" + sli[1]
	if localdir, ok := staticMap[prefix]; ok != false {
		file := localdir + r.URL.Path[len(prefix):]
		http.ServeFile(w, r, file)
		return
	}

	if r.URL.Path == "/show" {
		Show(w, r)
		return
	}

	if r.URL.Path == "/flowrate" {
		flowrate(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func Show(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}
