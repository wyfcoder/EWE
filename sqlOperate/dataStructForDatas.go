package sqlOperate

//==========================================================
//用于绘制文件的数据描述 存入数据库的数据表
type Data struct {
	Id       string
	Name     string
	Tag      string
	Time     string
	Describe string
}

func (d *Data) GetName()string{
	return d.Name
}

func (d *Data)GetTag()  string{
	return d.Tag
}

func (d *Data)GetTime()string{
	return d.Time
}

func (d *Data)GetDescribe() string{
	return d.Describe
}

func (d* Data)GetId() string{
	return d.Id
}

//=========================================
//用于绘制曲线的实际绘制模块
type DataPlot struct {
	Name string
	Tag  string
	Datas []float64
}

func (data DataPlot)GetLineName() string{
	return data.Name
}

func (data DataPlot)GetLineData() *[]float64  {
	return &data.Datas
}