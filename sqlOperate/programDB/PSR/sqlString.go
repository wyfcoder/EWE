package PSRDB

const
(
	preSql           = "select id,time,name from PSRData where"
	selectType0      = preSql +" id=$1 and time=$2 and name= $3"
	selectType1      = preSql +" id=$1 and name=$2"
	selectType2      = preSql +" id=$1 and time=$2"
	selectType3      = preSql +" id=$1"

	selectRealData   = "select data from PSRData where id=$1 and time=$2 and name= $3 "
)