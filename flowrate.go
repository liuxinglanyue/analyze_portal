// flowrate
package main

import (
	//"fmt"
	"html/template"
	//"log"
	"net/http"
	"strings"
)

func flowrate(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("flowrate.gtpl")
		//f1 := Friend{desc: ""}
		t.Execute(w, "")
	} else {
		r.ParseForm()

		if r.FormValue("start_time") == "" || r.FormValue("end_time") == "" {
			t, _ := template.ParseFiles("flowrate.gtpl")
			t.Execute(w, "时间不可为空!")
			return
		}

		dis_str := strings.Join(r.Form["dis"], ",")
		isp_str := strings.Join(r.Form["isp"], ",")
		//fmt.Println(dis_str, isp_str)

		var hql string
		//hql = "SELECT record_time, sum(flow_rate) FROM r_sector_flow_5 where record_time >= '" + r.FormValue("start_time") + "' and record_time < '" + r.FormValue("end_time") + "' "
		hql = "select record_time,sum(flow_rate),district_id,isp_id,sp_code from r_sector_flow_5 where record_time >= '" + r.FormValue("start_time") + "' and record_time < '" + r.FormValue("end_time") + "' "
		
		var app_dis string
		var app_isp string
		var app_sp string
		if strings.Count(dis_str, ",") != 33 && dis_str != "" {
			app_dis = " and district_id in (" + dis_str + ") "
		}
		if strings.Count(isp_str, ",") != 7 && isp_str != "" {
			app_isp = " and isp_id in (" + isp_str + ") "
		}
		if r.FormValue("code") != "" {
			app_sp = " and sp_code in ('" + strings.Replace(r.FormValue("code"), ",", "','", -1) + "') "
		}
		//fmt.Fprintf(w, search(hql+app_isp+app_dis+" group by record_time"))
		t, _ := template.ParseFiles("flowrate.gtpl")

		t.Execute(w, search(hql+app_isp+app_dis+app_sp+" group by district_id,isp_id,sp_code,record_time "))
	}
}

//SELECT record_time, sum(flow_rate) FROM r_sector_flow_5 where record_time >= '2014-10-26 00:00:00' and record_time < '2014-10-27 00:00:00' group by record_time
