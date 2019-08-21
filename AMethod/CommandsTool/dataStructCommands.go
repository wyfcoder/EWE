package CommandsTool

// 0 1 2 3 三种模式
//代表 (a,a,a) (a,*,a) (a,a,*) (a,*,*)
type SearchCommand struct {
	ID string
	Date string
	Name string
	Type int
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