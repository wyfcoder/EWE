package PSRDB

import (
	"database/sql"
	"errors"
	"github.com/WebForEME/AMethod/CommandsTool"
	"github.com/WebForEME/AMethod/TextDeal"
	"github.com/WebForEME/sqlOperate"
	"github.com/WebForEME/sqlOperate/programDB/PSR/PSRData"
	"strings"
)

//为PSR程序提供查询服务
//输入的数据为指令集，保存结果的指针
// 0 1 2 3 三种模式
//代表 (a,a,a) (a,*,a) (a,a,*) (a,*,*)
func GetPSRValue( commands []CommandsTool.SearchCommand, information * PSRData.Information) int{

	for i:=0;i < len(commands) ; i++ {

		rows := &sql.Rows{}

		err  := errors.New("")

		switch commands[i].Type {
		case 0:
			rows,err = sqlOperate.Db.Query(selectType0, commands[i].ID, commands[i].Date, commands[i].Name)
		case 1:
			rows,err = sqlOperate.Db.Query(selectType1, commands[i].ID, commands[i].Name)
		case 2:
			rows,err = sqlOperate.Db.Query(selectType2, commands[i].ID, commands[i].Date)
		case 3:
			rows,err = sqlOperate.Db.Query(selectType3, commands[i].ID)
		}

		if err != nil{
			return i
		}

		for rows.Next() {
			conv := CommandsTool.SearchCommand{}
			rows.Scan(&conv.ID, &conv.Date, &conv.Name)
			isContinue := false
			for i:=0 ; i< len(information.CommandsSearch) ; i++{
				if information.CommandsSearch[i].ID == conv.ID && information.CommandsSearch[i].Date == conv.Date && information.CommandsSearch[i].Name == conv.Name{
					isContinue = true
					break
				}
			}
			if isContinue{
				continue
			}
			information.CommandsSearch = append(information.CommandsSearch, conv)
		}
		rows.Close()
	}
	return -1
}


func SetValue(item []string,data *[]PSRData.PSRDataR){

	for i:=0 ; i<len(item) ; i++{
		psrData := PSRData.PSRDataR{}
		itemD   := strings.FieldsFunc(item[i],TextDeal.SpiltVertical)

		psrData.Id =itemD[0]
		psrData.Date=itemD[1]
		psrData.Name = itemD[2]
		sqlOperate.Db.QueryRow(selectRealData,psrData.Id,psrData.Date,psrData.Name).Scan(&psrData.Data)
		(*data) = append((*data),psrData)
	}

}

func DownloadFile(id string , date string , name string) string{
	data := ""
	sqlOperate.Db.QueryRow(selectRealData,id,date,name).Scan(&data)
	return data
}
