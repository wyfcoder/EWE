package RayRun

import (
	"github.com/WebForEME/AMethod/TimeTool"
	"github.com/WebForEME/sqlOperate"
)

//RayRun 程序的数据库函数

//插入一个数据 环境数据
const  insertIntoRayRun     = "insert into rayrun values($1,$2)"
const  selectRayRunText     = "select data from rayrun where id=$1"
const  selectIRIDATA        = "select year,month,day,data from iridata where id=$1"
const  deleteIRIDATA         = "delete from iridata"
const  insertIRIDATA         = "insert into iridata values($1,$2,$3,$4,$5)"
const (
	AM            = "02"
	PM            = "14.50"
	)
func insertRayRun(id string,text string){

	stmt, _ := sqlOperate.Db.Prepare(insertIntoRayRun)
	defer stmt.Close()
	stmt.QueryRow(id,text)
}
func SelectFromRayRun(id string,data *string){
	sqlOperate.Db.QueryRow(selectRayRunText,id).Scan(data)
}

//处理IRI数据根据信息返回一次数据
func GetIRIData(time string)(year string,month string,day string,text string){
	sqlOperate.Db.QueryRow(selectIRIDATA,time).Scan(&year,&month,&day,&text)
	return year,month,day,text
}

//更新数据
func UpDateIRIData(textAM string,textPM string){
	sqlOperate.Db.Exec(deleteIRIDATA)
	timeString :=TimeTool.TimeStringNow()
	stmt, _ := sqlOperate.Db.Prepare(insertIRIDATA)
	defer stmt.Close()
	stmt.QueryRow(AM,timeString.Year,timeString.Month,timeString.Day,textAM)
	stmt.QueryRow(PM,timeString.Year,timeString.Month,timeString.Day,textPM)
}
