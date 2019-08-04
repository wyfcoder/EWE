package PSR

import "github.com/WebForEME/AMethod/CommandsTool"

type Information struct {
	commands string
	error    string
	commandsSearch []CommandsTool.SearchCommand
}

func (information Information)GetCommands() string{
	return information.commands
}

func (information Information)GetError() string{
	return information.error
}

func (information Information)GetList() []CommandsTool.SearchCommand{
	return information.commandsSearch
}
