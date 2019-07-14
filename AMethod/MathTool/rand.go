package MathTool


import (
	"math/rand"
	"strconv"
	"time"
)

//产生一个随机码长度为16
func CreateVerificationCode() string  {
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	code:=""

	for i:=0;i<16;i++{
		code=code+strconv.Itoa((rand1.Int()%9))
	}
	return code
}

