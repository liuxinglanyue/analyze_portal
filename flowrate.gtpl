<html>
<link rel="stylesheet" type="text/css" href="assets/jquery.multiselect.css" />
<link rel="stylesheet" type="text/css" href="assets/style.css" />
<link rel="stylesheet" type="text/css" href="assets/prettify.css" />
<link rel="stylesheet" type="text/css" href="assets/jquery-ui.css" />
<link rel="stylesheet" type="text/css" href="assets/jquery-ui-timepicker-addon.css" />
<script type="text/javascript" src="assets/jquery.js"></script>
<script type="text/javascript" src="assets/jquery-ui.min.js"></script>
<script type="text/javascript" src="assets/prettify.js"></script>
<script type="text/javascript" src="assets/jquery.multiselect.js"></script>
<script type="text/javascript" src="assets/jquery-ui-timepicker-addon.js"></script>
<script type="text/javascript">
$(function(){
	$("select").multiselect();
	$("#datetime1").datetimepicker({ dateFormat: 'yy-mm-dd '});
	$("#datetime2").datetimepicker({ dateFormat: 'yy-mm-dd '});

	Date.prototype.Format = function(fmt){ //author: meizz  
	 	var o = {  
	   "M+" : this.getMonth()+1,                 //月份  
	   "d+" : this.getDate(),                    //日  
	   "h+" : this.getHours(),                   //小时  
	   "m+" : this.getMinutes(),                 //分  
	   "s+" : this.getSeconds(),                 //秒  
	   "q+" : Math.floor((this.getMonth()+3)/3), //季度  
	   "S"  : this.getMilliseconds()             //毫秒  
  	}; 
	  if(/(y+)/.test(fmt))  
	    fmt=fmt.replace(RegExp.$1, (this.getFullYear()+"").substr(4 - RegExp.$1.length));  
	  for(var k in o)  
	    if(new RegExp("("+ k +")").test(fmt))  
	  fmt = fmt.replace(RegExp.$1, (RegExp.$1.length==1) ? (o[k]) : (("00"+ o[k]).substr((""+ o[k]).length)));  
	  return fmt;  
	}
	var endDate = new Date();
	var minutes = endDate.getMinutes();
	minutes = 5*(parseInt(minutes/5));
	console.log(minutes);
	endDate.setMinutes(minutes);
	var startDate = new Date(endDate.getTime()-24*60*60*1000);

	$("#datetime1").val(startDate.Format("yyyy-MM-dd hh:mm"));
	$("#datetime2").val(endDate.Format("yyyy-MM-dd hh:mm"));

});
</script>
<head>
<title>分析带宽数据</title>
</head>
<body>
<form action="/flowrate" method="post">
    Sp Code:<input type="text" name="code">
	
    省份:<select title="Basic example" multiple="multiple" name="dis" size="5">
			<option value="110000">北京市</option>
			<option value="120000">天津市</option>
			<option value="130000">河北省</option>
			<option value="140000">山西省</option>
			<option value="150000">内蒙古自治区</option>
			<option value="210000">辽宁省</option>
			<option value="220000">吉林省</option>
			<option value="230000">黑龙江省</option>
			<option value="310000">上海市</option>
			<option value="320000">江苏省</option>
			<option value="330000">浙江省</option>
			<option value="340000">安徽省</option>
			<option value="350000">福建省</option>
			<option value="360000">江西省</option>
			<option value="370000">山东省</option>
			<option value="410000">河南省</option>
			<option value="420000">湖北省</option>
			<option value="430000">湖南省</option>
			<option value="440000">广东省</option>
			<option value="450000">广西壮族自治区</option>
			<option value="460000">海南省</option>
			<option value="500000">重庆市</option>
			<option value="510000">四川省</option>
			<option value="520000">贵州省</option>
			<option value="530000">云南省</option>
			<option value="540000">西藏自治区</option>
			<option value="610000">陕西省</option>
			<option value="620000">甘肃省</option>
			<option value="630000">青海省</option>
			<option value="640000">宁夏回族自治区</option>
			<option value="650000">新疆维吾尔自治区</option>
			<option value="710000">台湾省</option>
			<option value="810000">香港特别行政区</option>
			<option value="820000">澳门特别行政区</option>
		</select>
	
	运营商:<select title="Basic example" multiple="multiple" name="isp" size="5">
			<option value="20001">中国电信</option>
			<option value="20002">中国联通</option>
			<option value="20003">中国移动</option>
			<option value="20004">华数宽带</option>
			<option value="20005">方正宽带</option>
			<option value="20006">中国网通</option>
			<option value="20007">鹏博士</option>
			<option value="0">其他</option>
		</select>
	</br>
	</br>
	开始时间:<input id="datetime1" type="text" name="start_time">
	结束时间:<input id="datetime2" type="text" name="end_time">
    <input type="submit" value="查询">
</form>
</br>





<table  border="1">
  <tr>
    <th width="10%">省份</th>
    <th width="10%">运营商</th>
	<th width="10%">sp</th>
	<th width="10%">最大带宽</th>
	<th width="20%">时间</th>
  </tr>
{{range .}}

  <tr>
    <td>{{ .District}}</td>
    <td>{{ .Isp}}</td>
    <td>{{ .SpCode}}</td>
	<td>{{ .MaxFlow}}</td>
    <td>{{ .Time}}</td>
  </tr>
{{end}}
</table>


</body>
</html>