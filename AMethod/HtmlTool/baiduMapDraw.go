package HtmlTool

import "strconv"

//适配百度地图的数据
//{"lng":116.418261,"lat":39.921984,"count":50},
func DrawBaiduData(data *[]float64) string{
	html  := ""
	for i:=0 ; i<len(*data) ;i+=3{
		lng := strconv.FormatFloat((*data)[i], 'f', -1, 64)
		lat:= strconv.FormatFloat((*data)[i+1], 'f', -1, 64)
		count := strconv.FormatFloat((*data)[i+2],'f',-1,64)
		html += `{"lng" :`+lng + ","+`"lat":` + lat + ","+`"count":` + count + "}"
		if i != len(*data)-3 {
			html += ","
		}
	}
	return html
}