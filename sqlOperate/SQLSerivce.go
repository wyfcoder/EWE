package sqlOperate

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/WebForEME/AMethod/TextDeal"
	_ "github.com/lib/pq"
	"log"
)

//使用PostgreSQL,执行数据库链接，查询等操作

var Db *sql.DB

//链接数据库
func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err) //链接错误则直接退出即可
	}
	return
}

//加入客户的逻辑
func AddUsers(name string, account string, password string, findPasswordString string) (error, string) {
	statement := "insert into informationofusers values($1,$2,$3,$4) returning id"

	stmt, err := Db.Prepare(statement)

	defer stmt.Close()

	err = stmt.QueryRow(account, name, findPasswordString, password).Scan(&account)

	if err != nil {
		//找出错误的原因
		message := "Register wrong:"
		if Db.QueryRow(selectID, account).Scan(&account) == nil {
			message += "Account has been used."
		}
		if Db.QueryRow(selectName, name).Scan(&account) == nil {
			message += "Name has been used."
		}

		//不知名原因
		if message == ""{
			message += "Database wrongs ."
		}

		return err, message
	}
	//创建验证组
	statement = "insert into verificationCode values($1,$2) returning id"
	stmt, err = Db.Prepare(statement)
	//插入一个码
	err = stmt.QueryRow(account, "123").Scan(&account)
	return nil, ""
}

//登录
func Login(user string, password string) (User, error) {
	//user 可能是用户名，也能是账号
	u := User{
		Name:     "",
		Account:  "",
		Password: "",
		Sfp:      "",
	}
	if queryByName(user, &u) == nil {

	} else if queryByAccount(user, &u) == nil {

	} else {
		return u, errors.New("The account or the name is not exit.") //查找不到账号
	}

	if password != u.Password {
		return u, errors.New("The password is wrong.") //找回密码的String是错误的
	}

	return u, nil
}

//找回密码的逻辑
func FindBackPassword(user string, findPasswordString string) (User, error) {
	//user 可能是用户名，也能是账号
	u := User{
		Name:     "",
		Account:  "",
		Password: "",
		Sfp:      "",
	}
	if queryByName(user, &u) == nil {

	} else if queryByAccount(user, &u) == nil {

	} else {
		return u, errors.New("The account or the name is not exit.") //查找不到账号
	}

	if findPasswordString != u.Sfp {
		return u, errors.New("The String is wrong.") //找回密码的String是错误的
	}

	return u, nil
}

//通过名字查找，返回客户的信息
func queryByName(name string, user *User) error {
	return Db.QueryRow(selectByName, name).Scan(&user.Account, &user.Name, &user.Password, &user.Sfp)
}

//通过账号查找，保存的参数由kind给出
func queryByAccount(account string, user *User) error {
	return Db.QueryRow(selectByAccount, account).Scan(&user.Account, &user.Name, &user.Password, &user.Sfp)
}

//更新密码
func ResetPassword(account string, password string) {
	Db.Query("update informationofusers set password='" + password + "' where id='" + account + "'")
}

//更新验证码
func UpdateCode(account string, code string) {
	Db.Query("update verificationcode set code='" + code + "' where id='" + account + "'")
}

func UpdateCode2(account string, code string) {
	Db.Query("update ManagerVerificationcode set code='" + code + "' where id='" + account + "'")
}

//验证验证码
func CheckCode(account string, code string) bool {
	c := ""
	Db.QueryRow(selectCode, account).Scan(&c)
	if c != code {
		return false
	}
	return true
}

//验证Manger 验证码
func CheckMCode(account string, code string) bool {
	c := ""
	Db.QueryRow(selectCode2, account).Scan(&c)
	if c != code {
		return false
	}
	return true
}

//添加数据文件
func AddDataFile(name string, tag string, time string, describe *string, data *string, account string) int {
	if Db.QueryRow(selectByIdAndName, account, name).Scan(&account) == nil {
		return 1 //检查唯一性
	} else { //插入数据
		statement := "insert into datas values($1,$2,$3,$4,$5,$6) returning id"
		stmt, err := Db.Prepare(statement)
		if err != nil {
			return 2
		}
		defer stmt.Close()
		err = stmt.QueryRow(account, name, tag, time, *describe, *data).Scan(&account)
		if err != nil {
			return 2
		}
	}
	return 0
}

func DeleteData(account string, text string) {
	Db.Exec(deleteFile, account, text)
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


