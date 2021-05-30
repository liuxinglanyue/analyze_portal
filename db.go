// db
package main

import (
	"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"sort"
)

type Info struct {
    SpCode      string
    District    string
    Isp 		string
	Time 		string
	MaxFlow 	string
}

func  search(hql string) []Info {
	db, err := sql.Open("mysql", "root:123456@tcp(10.222.138.155:8066)/portal?charset=utf8")
	checkErr(err)

	//查询数据
	rows, err := db.Query(hql)
	log.Println("sql:", hql)
	checkErr(err)
	
	disMap := map[int]string{}
	disMap[110000] = "北京市"
	disMap[120000] = "天津市"
	disMap[130000] = "河北省"
	disMap[140000] = "山西省"
	disMap[150000] = "内蒙古自治区"
	disMap[210000] = "辽宁省"
	disMap[220000] = "吉林省"
	disMap[230000] = "黑龙江省"
	disMap[310000] = "上海市"
	disMap[320000] = "江苏省"
	disMap[330000] = "浙江省"
	disMap[340000] = "安徽省"
	disMap[350000] = "福建省"
	disMap[360000] = "江西省"
	disMap[370000] = "山东省"
	disMap[410000] = "河南省"
	disMap[420000] = "湖北省"
	disMap[430000] = "湖南省"
	disMap[440000] = "广东省"
	disMap[450000] = "广西壮族自治区"
	disMap[460000] = "海南省"
	disMap[500000] = "重庆市"
	disMap[510000] = "四川省"
	disMap[520000] = "贵州省"
	disMap[530000] = "云南省"
	disMap[540000] = "西藏自治区"
	disMap[610000] = "陕西省"
	disMap[620000] = "甘肃省"
	disMap[630000] = "青海省"
	disMap[640000] = "宁夏回族自治区"
	disMap[650000] = "新疆维吾尔自治区"
	disMap[710000] = "台湾省"
	disMap[810000] = "香港特别行政区"
	disMap[820000] = "澳门特别行政区"
	
	ispMap := map[int]string{}
	ispMap[20001] = "中国电信"
	ispMap[20002] = "中国联通"
	ispMap[20003] = "中国移动"
	ispMap[20004] = "华数宽带"
	ispMap[20005] = "方正宽带"
	ispMap[20006] = "中国网通"
	ispMap[20007] = "鹏博士"
	ispMap[0] = "其他"

	
	//data := make(map[string]Info)
	data := map[Info]map[string]float64{}
	

	//district + isp + sp_code 的同一时间的flow合并

	for rows.Next() {
		var time string
		var flow float64
		var district_id int
		var isp_id int
		var sp_code string
		var district string
		var isp string
		err = rows.Scan(&time, &flow,&district_id,&isp_id,&sp_code)
		checkErr(err)
		
		district  = disMap[district_id]
		isp = ispMap[isp_id]
		//fmt.Println(district+"and"+isp)
		
		//var key = district +"--"+ isp + "--" + sp_code

		
		key := Info{SpCode:sp_code,District:district,Isp:isp}
		time_flow, ok := data[key]
		if !ok {
			
			m := map[string]float64{}
			m[time] = flow
			data[key] = m
			//mapSize++
		} else {
			
			old_flow,same := time_flow[time]
			if same {
				time_flow[time] = old_flow+flow
			}else {
				time_flow[time] = flow
			}
			
		}	
	}
	//flow合并后取最大的flow
	
	var resultData = make([]Info, len(data))
	var i int = 0
	//fmt.Println(i)
	for keys, values := range data {
		var maxFlow float64 = 0
		for key,value := range values {	
			if maxFlow < value {
				maxFlow = value
				keys.Time = key
			}
		}
		//1024.1024.300*8=39321600
		keys.MaxFlow = strconv.FormatFloat(maxFlow/39321600,'f', 2, 32)+"Mbps"
		resultData[i] = keys
		i++
		
	}
	sort.Sort(Bydis(resultData))
	db.Close()

	return resultData
	

}

type Bydis []Info

func (a Bydis) Len() int           { return len(a) }
func (a Bydis) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Bydis) Less(i, j int) bool { 
	if a[i].District == a[j].District {
		return a[i].Isp < a[j].Isp 
	}
	return a[i].District < a[j].District 
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}



