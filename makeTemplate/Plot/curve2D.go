package Plot

import (
	"github.com/WebForEME/makeTemplate"
	"github.com/WebForEME/sqlOperate"
)

func Curve2D(data *[]sqlOperate.DataPlot) string{

	linesD :=  []makeTemplate.DataPlot{}
	for i:=0 ; i<len(*data) ;i++{
		linesD = append(linesD,(*data)[i])
	}
	dataString := makeTemplate.DataDraw2D(&linesD)
	selectComponent := makeTemplate.MakeSelectComponent(&linesD)
	return Head + Navbar + BodyPre+selectComponent + BodyLast + dataString + Tail

}
