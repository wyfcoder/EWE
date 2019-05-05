package sqlOperate

//RayRun 程序的数据库函数

//插入一个数据
const  insertIntoRayRun     = "insert into rayrun values($1,$2)"
const  selectRayRunText     = "select data from rayrun where id=$1"

func insertRayRun(id string,text string){
	stmt, _ := Db.Prepare(insertIntoRayRun)
	defer stmt.Close()
	stmt.QueryRow(id,text)
}
func SelectFromRayRun(id string,data *string){
	Db.QueryRow(selectRayRunText,id).Scan(data)
}
