package main

import (
	"bufio"
	"fmt"
	"os"
)

// json
type Movies struct {
	// 变量名大写，   `json:"name"`  json符号写上
	Name string		`json:"name"`
	Date string		`json:"date"`
}

func main(){
	NumToChinese()

}
func inputFromKeyBoardOne() int {
	var x int
	n, err := fmt.Scan(&x)
	if err != nil{
		panic(err)
	}
	fmt.Println(n)

	return x
}

func inputFromKeyBoardTwo() string {
	reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
	return res
}

func NumToChinese(){
	var data = inputFromKeyBoardTwo()

	var word = []string{"","十","百","千","万"}
	var length = len(data) - 1
	var res = ""
	var chs = []byte(data)
	for i:=0;i<length;i++{
		if chs[i] == '0' && chs[i+1] == '0'{
			continue
		}else if chs[i] == '0'{
			res += "零"
		}else{
			res += string(getChinese(chs[i])) + word[length - i- 1]
		}
	}
	fmt.Println(res)
}

func getChinese(i byte) string{
	var chinese = ""
	switch i {
	case 57:
		chinese = "九"
	case 56:
		chinese = "八"
	case 55:
		chinese = "七"
	case 54:
		chinese = "六"
	case 53:
		chinese = "五"
	case 52:
		chinese = "四"
	case 51:
		chinese = "三"
	case 50:
		chinese = "二"
	case 49:
		chinese = "一"
	case 48:
		chinese = "零"
	}
	return chinese
}

func jsonTest(){
	/*var m = Movies{Name: "liu", Date:"20220110"}
	data ,err := json.MarshalIndent(m, "", "\t")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n", data)*/
	//{
	//        "name": "liu",
	//        "date": "20220110"
	//}

	/*data ,err := json.MarshalIndent(m, "", "\t")
	  fmt.Printf("%s\n", data)
	{"name":"liu","date":"20220110"}
	*/
}