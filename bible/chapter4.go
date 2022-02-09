package main

import (
	"encoding/json"
	"fmt"
)

// json
type Movies struct {
	Name string		`json:"name"`
	Date string		`json:"date"`
}

func main(){
	var m = Movies{Name: "liu", Date:"20220110"}
	data ,err := json.MarshalIndent(m, "", "\t")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n", data)
	//{
	//        "name": "liu",
	//        "date": "20220110"
	//}

	/*data ,err := json.MarshalIndent(m, "", "\t")
	  fmt.Printf("%s\n", data)
	{"name":"liu","date":"20220110"}
	*/
}
