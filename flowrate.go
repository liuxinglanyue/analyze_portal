// flowrate
package main

import (
	//"fmt"
	"html/template"
	//"log"
	"net/http"
	//"strings"
)

func flowrate(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("flowrate.gtpl")
		t.Execute(w, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		//fmt.Println("username:", r.Form["username"])
		//fmt.Println("password:", r.Form["password"])
	}
}
