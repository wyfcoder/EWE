package CommandsTool

type SearchCommand struct {
	ID string
	Date string
	Name string
}


func (command SearchCommand) GetID() string{
	return command.ID
}

func (command SearchCommand) GetDate() string{
	return command.Date
}

func  (command SearchCommand) GetName() string  {
	return command.Name
}