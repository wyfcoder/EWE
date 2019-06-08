package sqlOperate

const (
	host              = "localhost"
	port              = 5432
	user              = "postgres"
	password          = "wyf"
	dbname            = "ewe"

	//验证存在性的方式
	selectID          = "select id from informationofusers where id=$1"
	selectName        = "select id from informationofusers where name=$1"

	selectByName      = "select id,name,password,sfp from informationofusers where name=$1"
	selectByAccount   = "select id,name,password,sfp from informationofusers where id=$1"
	selectCode        = "select code from verificationcode where id=$1"
	selectByIdAndName = "select  id from datas where id=$1 and fileName=$2"
	deleteFile        = "delete from datas where id=$1 and fileName=$2"

	selectManagerName    = "select name,account,password from Manager where name=$1"
	selectManagerAccount = "select name,account,password from Manager where account=$1"
	selectCode2          = "select code from ManagerVerificationcode where id=$1"
	selectNotice         = "select Title,Content,ptime FROM notice"
	deleteNotice         = "delete from Notice where pTime=$1 and title=$2"
	insertIntoNotice     = "insert into Notice values($1,$2,$3)"
	insertIntoFileM      = "insert into filesm values($1,$2,$3,$4)"
	selectFile           = "select title,tag,ptime,path FROM filesm"
	deleteFileM          = "delete from filesm where title=$1 and ptime=$2"
)
