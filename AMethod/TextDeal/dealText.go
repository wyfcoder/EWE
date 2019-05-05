package TextDeal

import (
	"errors"
	"math"
	"strconv"
)

//处理一段字符输入将其变为更加规整的数据流,其中要解决一些特殊的数字表达形式
//返回 数据是否正确 数据量 规整的字符数
func DealText(text *string) (bool, int, string) {
	count := 0
	datas := ""
	number:=0.0
	isContinue :=false
	err:=errors.New("")
	//遇到换行，空格时自动的忽略 特殊计数法明确转换
	for i := 0; i < len(*text); i++{
		char:=(*text)[i]
		if isNumber(char){
			isContinue=true
		    i,number,err=readANumber(text,i)
			if err != nil{
				return false,0,""
			}
		}
		if  isContinue{

			if i<len(*text)-1 {
				char = (*text)[i]
				if char == 'E' || char == 'e' { //科学计数法计数
					i++
					char = (*text)[i]
					i++
					number2 := 0.0
					i, number2, err = readANumber(text, i)
					if err != nil {
						return false, 0, ""
					}
					if char == '+' {
						number = number * math.Pow(10, number2)
					} else {
						number = number * math.Pow(10, -number2)
					}
				}
			}
			count++
			datas +=strconv.FormatFloat(number,'f',-1,64)+" "
			isContinue=false
		}
	}

	return true,count,datas
}

//处理文本使其返回一个数组

func DealText2(text *string,data *[]float64){
	s:=[]byte{}
	for i := 0; i < len(*text); i++ {
		if (*text)[i] != ' ' {
			s = append(s, (*text)[i])
			if i != len(*text)-1 {
				continue
			}
		}
		s2 := string(s)
		number, _ := strconv.ParseFloat(s2, 64)
		(*data) = append((*data), number)
		s = s[0:0]
	}
}

//返回一个数字
func readANumber(text *string,start int) (int,float64,error){
	s := []byte{}
	numberText:= ""
	for ;start<len(*text);start++{
		if  isNumber((*text)[start]){
			s = append(s, (*text)[start])
		}else{
			break
		}
	}
	numberText = string(s)
	number, err := strconv.ParseFloat(numberText, 64)

	if err!=nil{
		return start,0.0,err
	}else{
		return start,number,nil
	}
}

func isNumber(char byte) bool{

	if (char >= '0' && char <='9') || char == '.'{
		return true
	}else{
		return false
	}

}

