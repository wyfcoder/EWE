package makeTemplate

//适配于输入的数据
//能够访问到其数据和曲线名称即可
type DataPlot interface {
	GetLineName() string
	GetLineData() *[]float64
}
