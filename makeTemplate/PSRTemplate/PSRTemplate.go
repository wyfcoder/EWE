package PSRTemplate

import (
	"github.com/WebForEME/AMethod/HtmlTool"
	"github.com/WebForEME/AMethod/TextDeal"
	"github.com/WebForEME/AMethod/baseTool"
	"github.com/WebForEME/sqlOperate/programDB/PSR/PSRData"
	"strconv"
)

//制作模板
func MakePSRTemplate( data *[]PSRData.PSRDataR) string{

	html := HtmlHead

	makeList(&html,data)

	html += HtmlMapPre

	makeDataList(&html,data)

	html += HtmlTail

	return html
}

func makeList(html *string, data *[]PSRData.PSRDataR){
	for i:= 0; i< len(*data); i++{
		(*html) += `<div class="panel panel-default">`
		(*html) += ` <div class="panel-heading">`
		(*html) += `<span class="lead">ID :` + (*data)[i].Id + `</span>`
		(*html) += `</div>`
		(*html) += `   <div class="panel-body">
           <div class="pull-right">
               <input type="checkbox" name="checkbox" id= `+ strconv.Itoa(i)+`>
           </div>`
		(*html) +="Date : " + (*data)[i].Date + " FileName : " + (*data)[i].Name
		(*html) += `</div>
                    </div>
`
	}
}

//制作数据序列 填充数组
// var points =[
func makeDataList(html *string, data *[]PSRData.PSRDataR){

	number := []float64{}
	index  := []int{}

	for i:=0 ; i< len(*data) ; i++{
		//把数据转化为数组
		dataL := []float64{}
		TextDeal.DealText2(&((*data)[i]).Data,&dataL)

		//处理错误 文件数据格式错误
		if len(*data) % 3 !=0 {
			//写一个具体的文件错误内容，并保存至cookie，之后直接显示即可。
			return
		}
		baseTool.CombineFloatList(&number,&dataL)
		if i == 0 {
			index = append(index, len(dataL)/3 - 1)
		}else{
			index = append(index,len(dataL)/3 + index[i-1] - 1)
		}
	}

	//设置必要的参数

	if len(number) >= 3 {
		(*html) += `var point = new BMap.Point(` + strconv.FormatFloat(number[0], 'f', -1, 64) + `,` + strconv.FormatFloat(number[1], 'f', -1, 64) + `);`
	}else{
		(*html) += `var point = new BMap.Point(108.838103,34.130109);`
	}
	(*html) += `map.centerAndZoom(point, 15);`
	(*html) += `var points =[`

	(*html) += HtmlTool.DrawBaiduData(&number)

	(*html) += `];`

	(*html) += `var list =[`

	(*html) += HtmlTool.DrawList(&index)

	(*html) += `];`

	(*html) += htmlMapScriptPre

	max := 0.0

	for i:=0 ; i< len(number) ;i+=3 {
		if number[i+2] > max{
			max = number[i+2]
		}
	}

	(*html) += `heatmapOverlay.setDataSet({data:points,max:`+strconv.FormatFloat(max,'E',-1,64)+`});`


	(*html) += htmlMapScriptNext
}