package sqlOperate

//把时间,数据,用户名存入到数据库中
const insertIntoFeedback     = "insert into feedback values($1,$2,$3) returning account"

const deleteFeedback         ="delete from feedback where time=$1 and account=$2"

const selectFeedback       = "select time,message,account from feedback"

func InsertIntoFeedbackTable(account string,message string,time string) error{

	stmt, err := Db.Prepare(insertIntoFeedback)

	if err!= nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(account, time, message).Scan(&account)

	return err
}


func DeleteFeedback(account string, time string){

	Db.Exec(deleteFeedback, time, account)
}
