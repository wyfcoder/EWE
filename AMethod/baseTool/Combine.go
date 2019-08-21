package baseTool


func CombineFloatList(receive *[]float64 , sender *[]float64){

	for i:=0 ; i< len(*sender) ; i++{
		(*receive) = append((*receive),(*sender)[i])
	}

}
