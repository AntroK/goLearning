package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

type Person struct {
	Name 	string
	Age		int
}

const temp1 = `{{.TotalCount}} issues:
{{range.Items}}-----------------------
Number: {{.Number}}
User:	{{.User.Login}}
Title:	{{.Title | printf "%.64s"}}
Age:	{{.CreateAt | daysAgo}} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

/*var report = template.Must(template.New("issuelist").Funcs(template.FuncMap{"daysAgo": daysAgo(time.Now())}).Parse(temp1))
*/
func PersonPrint(){
	p := Person{"liuqi", 23}
	temp, err := template.New("test").Parse("Name: {{.Name}}, Age:{{.Age}}")
	if err != nil{
		panic(err)
	}
	err = temp.Execute(os.Stdout, p)
	if err != nil{
		panic(err)
	}
	fmt.Println(temp)
}

func printHello(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("template/test.html")
	if err != nil{
		panic(err)
	}
	t1.Execute(w, "hello world")
}

func main(){
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/print", printHello)
	server.ListenAndServe()
	//PersonPrint()
	/*var data string
	if err := report.Execute(os.Stdout, data); err != nil{
		panic(err)
	}*/
}

