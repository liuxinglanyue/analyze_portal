// db
package main

import (
	"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func search(hql string) string {
	db, err := sql.Open("mysql", "root:123456@tcp(10.222.138.155:8066)/portal?charset=utf8")
	checkErr(err)

	//查询数据
	rows, err := db.Query(hql)
	log.Println("sql:", hql)
	checkErr(err)

	m := map[string]float64{}
	var max_flow float64
	var max_time string

	for rows.Next() {
		var time string
		var flow float64
		err = rows.Scan(&time, &flow)
		checkErr(err)

		var old_flow float64
		old_flow = m[time]
		if old_flow == 0 {
			m[time] = flow
		} else {
			m[time] = old_flow + flow
		}

		if m[time] > max_flow {
			max_flow = m[time]
			max_time = time
		}
	}

	db.Close()
	//fmt.Println(max_time, max_flow)

	var sum float64
	sum = max_flow / 1024 / 1024 * 8

	return "时间：" + max_time + "  带宽：" + strconv.FormatFloat(sum, 'f', 2, 32) + "Mbps"
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
