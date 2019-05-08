package HtmlTool

import "strconv"

//把数据转化为 string 合适于网页显示
func DrawData( data *[]float64) string{

	html :=""
	for i := 0; i < len(*data); i += 2 {
		x := strconv.FormatFloat((*data)[i], 'f', -1, 64)
		y := strconv.FormatFloat((*data)[i+1], 'f', -1, 64)
		html += "[" + x + "," + y + "]"
		if i != len(*data)-2 {
			html += ","
		}
	}

	return html
}