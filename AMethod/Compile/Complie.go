package Compile

import (
	"errors"
)
//编译系统

//输入的是一段文本，从中解析出每一条指令的
// E_IRI() E_IRI( 2016-01-10 ,100 , 200 )

//每条指令 由头部和body构成
//头部代表其名称 body代表其参数
//在解析时采用从后向前 依次把每个参数找到
//body解析的开始为)结束(
//head为解析开始为(结束为(后的一段字符

//约定

//开头 @

//head -> 由字母和下滑线构成 直到非约定出现

//body -> 每个参数之间由逗号隔开，保证连续的一段，中间可以有任意数量的空格

//参数:不加约束，可由字母和数字构成
func Compile(text *string)(error,[]Instruct,string) {

	instructs:= []Instruct{}

	instruct:=Instruct{}

	err :=errors.New("Command is wrong ")

    hasStart :=false

	i:=0

	for i=len(*text)-1;i >=0;i--{

		char :=(*text)[i]
		//开始一条指令
		if char == ')' {
			i--
			for char == ')' && i >=0 {
				err, i = readBody(text, i, &instruct)
				if err != nil {
					return err, instructs,MakeWrongText(text,i)
				}
				err, i = readHead(text, i, &instruct)
				if err != nil {
					return err, instructs,MakeWrongText(text,i)
				}
				instructs = append(instructs, instruct)
				instruct.Body=nil
				if i >= 0 {
					char = (*text)[i]
				}else {
					break
				}
				i--
			}
		}

		if char == '@'{
			hasStart=true
			break
		}
	}

	if hasStart {
		ReverseInstructs(&instructs)
	}else{
		err =errors.New("Don't have @ .")
		return err,nil,MakeWrongNull(text)
	}
	return nil,instructs,MakeReturnText(text,i)
}

//读取body内容 结束：( ,开头: 2,1)
//对于 ( 1 1 ,1) 报错
func readBody(text *string,start int,instruct *Instruct)(error,int){

	isEnd:=false

	param := ""

	err :=errors.New("It has no end.")

	for ;start >= 0 && !isEnd;start-- {

		err,start,param =readContinue(text,start,true)

		if err !=nil{
			err:=errors.New(" Wrong  at body near "+param)
			return err,start
		}
		if param !="" {
			(*instruct).Body = append((*instruct).Body, param)
		}
		char :=(*text)[start]
		if char == '('{
			isEnd=true
		}
	}
	ReverseStringArray(&instruct.Body)
	return nil,start
}

//读取head内容 读连续的一段
func readHead(text *string,start int,instruct *Instruct) (error,int){

	err,end,head :=readContinue(text,start,false)

	if err != nil{
		err:=errors.New("Wrong at head near "+head)
		return err , end
	}

	(*instruct).Head = head

	return nil,end
}
//读连续的一段 返回 翻转之后的字符 以及是否有误
func readContinue(text *string , start int , isBody bool) (error,int,string){
	isSpace:=false
	param:=""
	err :=errors.New("It has no end.")

	for ;start >= 0 ;start-- {

		char :=(*text)[start]
		if char == ',' || char == '(' || char == ')' || char == '@'{
			if param == "" {                             // 出现了 ( , , )

			    if char != '(' {
					return err, start, param
				}else {
					return nil,start,param
				}
			}
			ReverseString(&param)            //翻转 //收集一条指令
			break
		} else if char == ' ' ||  char== '\n' || char== '\r'{
			if !isSpace && param != "" {
				isSpace = true
			}
		}else{
			if isSpace{
				if !isBody {
					ReverseString(&param)
					break
				} else {
					return err,start,param
				}  // 出现了 (2 2 , 1 2 ,  2 2 )
			}
			param += string(char)
		}
	}

	return nil,start,param
}

//翻转字符数组，输入字符数组
func ReverseStringArray(params * []string){
	 end := 0
	 size:= len(*params)
	 if  size <= 1{
		 return
	 }

	 if size % 2 == 0{
	 	end = size /2
	 }else{
	 	end = (size-1) / 2
	 }

	 for i:=0 ; i< end ; i++ {
	 	 temp := (*params)[i]
		 (*params)[i]=(*params)[size-1-i]
		 (*params)[size-1-i]=temp
	 }

}

//翻转字符串，输入字符串
func ReverseString(param *string){
	end := 0
	size:= len(*param)
	rev:=""

	for end =size-1 ; end >= 0 ;end --{
		rev+= string((*param)[end])
	}
	*param=rev
}

//翻转指令
func ReverseInstructs(instructs *[]Instruct){
	end := 0
	size:= len(*instructs)
	if  size <= 1{
		return
	}

	if size % 2 == 0{
		end = size /2
	}else{
		end = (size-1) / 2
	}

	for i:=0 ; i< end ; i++ {
		temp := (*instructs)[i]
		(*instructs)[i]=(*instructs)[size-1-i]
		(*instructs)[size-1-i]=temp
	}
}

//制作返回文本
func MakeReturnText(text *string, start int) string{
	txt:=""

	size := len(*text)

	for start < size{
		txt += string((*text)[start])
		start++
	}
	if (*text)[size-1] != '\n'{
		txt += string('\r') + string('\n')
	}
	return txt
}

//黏贴 两段 信息 组合为错误信息
func CombineText(pre *string,wrong *string){
	(*pre) = (*pre) + string('*')+string('\n')+(*wrong)
}

//制作错误信息文本
func MakeWrongText(text *string, start int) string{
	wrong :=""
	//首先还原用户的有效输入
	head :=start

	size :=len(*text)

	for head >= 0{
		if (*text)[head] == '@'{
			break
		}
		head --
	}

	for head < size{
		wrong += string((*text)[head])
		head ++
	}

	if (*text)[size-1] != '\n'{
		wrong += string('\r') + string('\n')
	}
	wrong += "*"
	wrong += string('\r') + string('\n')
	//定位出错误的行,并返回
	head = start

	for head  >= 0{
		if (*text)[head] == '\n'{
			break
		}
		head --
	}
	head++
	wrong += "wrong at "
	for head < size{
		if (*text)[head] == '\n'{
			break
		}
		wrong += string((*text)[head])
		head++
	}
	return wrong
}

//空开始错误文本
func MakeWrongNull(text *string) string{
	return (*text)+string('*')+string('\n')+"Don't have @ ."
}
//制作还原指令指令
func ReturnString(instruct Instruct)string{

	text := instruct.Head

	text+="("

	for i:=0;i< len(instruct.Body);i++{
		text += instruct.Body[i]
		if i!=len(instruct.Body)-1{
			text+=","
		}else{
			text+=")"
		}
	}
	return text
}


//提取 1~2~3 类似的指令
func GetParams(text string) ([]string,bool){
	param:=""
	list:=[]string{}
	for i:=0 ; i<len(text) ;i++{
		if text[i] == '~'{
			if param == ""{
				return nil,false
			}
			list=append(list,param)
			param=""
			continue
		}
		param+=string(text[i])
	}
	list=append(list,param)
	return list,true
}
