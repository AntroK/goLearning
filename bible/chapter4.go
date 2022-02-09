package main

import (
	"encoding/json"
	"fmt"
)

// json
type Movies struct {
	// 变量名大写，   `json:"name"`  json符号写上
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
