package sqlOperate

import (
	"github.com/WebForEME/AMethod/TextDeal"
)

//TODO：错误的处理
//删除绘图数据文件
func DeleteData(account string, text string) {
	Db.Exec(deleteFile, account, text)
}

//完成数据文件的临时生产,查询结果只有一条数据
func DownloadData(account string , text string) string{

	rows, err := Db.Query("SELECT information FROM datas where id='" + account + "' and fileName ='"+text+"'")
	if err != nil{
		return ""
	}
	fileContent  := ""
	rows.Next()
	rows.Scan(&fileContent)
	return fileContent
}


//提取用户的数据 --不包含具体图像数据
func Datas(account string) (datas []Data, err error) {
	rows, err := Db.Query("SELECT Id,fileName, tag, time, describe FROM datas where id='" + account + "'")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Data{}
		if err = rows.Scan(&conv.Id, &conv.Name, &conv.Tag, &conv.Time, &conv.Describe); err != nil {
			return
		}
		datas = append(datas, conv)
	}
	rows.Close()
	return
}

//提取用户数据-- 保护具体的图像数据
func DataTables(account string,data *[]DataPlot) error{
	rows, err := Db.Query("SELECT fileName, tag, information FROM datas where id='" + account + "'")
	if err != nil{
		return err
	}
	for rows.Next() {
		conv := DataPlot{}
		information := ""
		if err = rows.Scan(&conv.Name, &conv.Tag, &information); err != nil {
			return err
		}
		TextDeal.DealText2(&information,&conv.Datas)
		(*data) = append((*data), conv)
	}
	return nil
}

//数据序列化
func GetValue(id string, name string) []float64 {
	data, err := Db.Query("SELECT information FROM datas where id='" + id + "'and fileName='" + name + "'")
	text := ""
	datas := []float64{}
	if err == nil {
		data.Next()
		data.Scan(&text)
		TextDeal.DealText2(&text,&datas)
		return datas
	} else {
		return nil
	}
}