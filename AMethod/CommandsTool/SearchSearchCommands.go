package CommandsTool

import (
	"github.com/WebForEME/AMethod/TextDeal"
	"github.com/WebForEME/AMethod/TimeTool"
	"strings"
)

//对传入的参数进行分析
func SearchCommands(commands []string) (bool,int,[]SearchCommand){
	searchCommands := []SearchCommand{}
	for index , command :=   range  commands{
		ok,searchCommand :=searchCommandTrue(command)
		if !ok {
			return false,index,searchCommands
		}
		searchCommands= append(searchCommands,searchCommand)
	}
	return true,-1,searchCommands
}

//对一条指令进行检查
//ID,date,name 10,20,30  10,*,*
func searchCommandTrue(command string) (bool,SearchCommand){
	//id 检查
	commandData := SearchCommand{}

	list := strings.FieldsFunc(command , TextDeal.SpiltSpace)

	if len(list) != 3{
		return false,commandData
	}
	//对三个元素进行判断
	if list[0] == " " || list[0] == "*"{
		return false,commandData
	}
	_,ok := TimeTool.StringToDate(list[1])
	if list[1] != "*" &&  !ok{
		return false,commandData
	}

	if list[2]!="*" && len(list[2])==0{
		return false,commandData
	}
	commandData.ID=list[0]
	commandData.Date=list[1]
	commandData.Name=list[2]
	return true,commandData
}