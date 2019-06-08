package sqlOperate

import (
	"errors"
	"time"
)

//Manager login

func ManagerLogin(user string, password string) (User, error) {
	u := User{
		Name:     "",
		Account:  "",
		Password: "",
		Sfp:      "",
	}
	if querayManger(user, &u) == nil {

	} else if querayManger2(user, &u) == nil {

	} else {
		return u, errors.New("The name is not exit.") //查找不到账号
	}
	if u.Password != password {
		return u, errors.New("The password is not exit.") //查找不到账号
	}
	return u, nil
}

func querayManger(userName string, information *User) error {
	return Db.QueryRow(selectManagerName, userName).Scan(&information.Name, &information.Account, &information.Password)
}

func querayManger2(userName string, information *User) error {
	return Db.QueryRow(selectManagerAccount, userName).Scan(&information.Name, &information.Account, &information.Password)
}

func AddNotice(name string, content string) {
	stmt, _ := Db.Prepare(insertIntoNotice)
	defer stmt.Close()
	stmt.QueryRow(name, content, time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
}

func DeleteNotice(title string,date string){
	Db.Exec(deleteNotice, date, title)
}

func Notices() (notices []Notice) {
	rows, err := Db.Query(selectNotice)

	for rows.Next() {
		conv := Notice{}
		if err = rows.Scan(&conv.Title, &conv.Content, &conv.Date); err != nil {
			return
		}
		notices = append(notices, conv)
	}
	rows.Close()
	return
}

func AddFileM(title string,tag string,path string){
	stmt, _ := Db.Prepare(insertIntoFileM)
	defer stmt.Close()
	stmt.QueryRow(title,tag,path,time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
}

func DeleteFileM(title string,date string){
	Db.Exec(deleteFileM, title, date)
}

func Files() (files []FileInformation){
	rows, err := Db.Query(selectFile)
	for rows.Next() {
		conv := FileInformation{}
		if err = rows.Scan(&conv.Title, &conv.Path, &conv.Date,&conv.Tag); err != nil {
			return
		}
		files = append(files, conv)
	}
	rows.Close()
	return
}