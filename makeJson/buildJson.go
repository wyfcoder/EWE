package makeJson

import (
	"encoding/json"
	"os"
)

type HeatMap struct {
  Lng   float64           `json:"lng"`
  Lat   float64           `json:"lat"`
  Count int               `json:"count"`
}

type D2 struct {
	x float64          `json:"x"`
	y float64          `json:"y"`
}

func MakeHeatJson(heatMap[] HeatMap,id string)   (error,* os.File){
	jsonFile,err:=os.Create("Data/heatMap/"+id+".json")
	encoder:= json.NewEncoder(jsonFile)
	if err==nil{
			     err=encoder.Encode(&(heatMap))
			}
	return err,jsonFile
}


