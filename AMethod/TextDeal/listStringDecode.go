package TextDeal

import "strconv"

//解码 类似 _123_21_1_2的字符 生成 int 数组
//保证输入的串没有错误
func DecodeListString2i(listString string,list *[]int){
	i :=1
	numberString := ""
	number := 0
	for i<len(listString){
		for listString[i] != ' '{
			numberString += string(listString[i])
			i++
			if i >= len(listString){
				break
			}
		}
		number,_ = strconv.Atoi(numberString)
		(*list) = append((*list),number)
		numberString = ""
		i++
	}
}

func DecodeListString2s(listString string,list* []string){
	i :=1
	textString := ""
	for i<len(listString) {
		for listString[i] != ' ' {
			textString += string(listString[i])
			i++
			if i >= len(listString) {
				break
			}
		}
		(*list) = append((*list), textString)
		textString = ""
		i++
	}
}
