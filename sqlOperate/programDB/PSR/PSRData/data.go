package PSRData

import "github.com/WebForEME/AMethod/CommandsTool"

//数据结构用于web显示
type Information struct {
	Commands string
	Error    string
	CommandsSearch []CommandsTool.SearchCommand
}

func (information Information)GetCommands() string{
	return information.Commands
}

func (information Information)GetError() string{
	return information.Error
}

func (information Information)GetList() []CommandsTool.SearchCommand{
	return information.CommandsSearch
}

//数据，用于绘制图像

type PSRDataR struct {
	Id string
	Date string
	Name string
	Data string
}