package makeTemplate

//绘图组件颜色选择组件
func MakeSelectComponent(plot *[]DataPlot) string{
	selectScript := `<select id="selectLine">`
	for i:=0 ; i < len(*plot) ;i++{
		selectScript += `<option>`+(*plot)[i].GetLineName()+`</option>`
	}
	selectScript +=`</select>`
	return selectScript
}
