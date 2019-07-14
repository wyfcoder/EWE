package RayRunDataStruct

//数据 二维数据集 名称标识 射线的频率用于排序
type RayRunData struct {
	Data []float64
	Name string
}

func (data RayRunData)GetLineName() string{
	return data.Name
}

func (data RayRunData)GetLineData() *[]float64{
	return &data.Data
}



