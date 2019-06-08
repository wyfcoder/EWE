package Compile

//在指令集中间莫个位置 插入一段新的指令集
func Combine(main *[]Instruct,insert *[]Instruct,i int){
	new :=[]Instruct{}
	for k:=0 ; k<i ;k++{
		new=append(new,(*main)[k])
	}
	for k:=0;k<len(*insert);k++{
		new=append(new,(*insert)[k])
	}
	for k:=i+1;k<len(*main);k++{
		new=append(new,(*main)[k])
	}
	*main=new
}
