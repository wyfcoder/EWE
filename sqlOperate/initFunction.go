package sqlOperate

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

//初始化必要数据库

//初始化 RayRun 数据库 //打开数据文件 取值 加入数据库
func InitRayRun(){
	fi, err := os.Open("sqlOperate/datas/iri数据.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	i:=0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i !=0 {
			insertRayRun(strconv.Itoa(i),string(a))
		}
		i++
	}
}
