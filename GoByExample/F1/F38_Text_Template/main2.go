package main
import (
	//"fmt"
	"text/template"
	"os"
)

func main() {
	tmp,err:=template.New("Hello").Parse(`Hello, {{ .Name}}!`)//if invalidsyntax returning error
	if(err!=nil) {
		panic(err)
	}
	data:=map[string]string{"Name":"Aakash"}
	tmp.Execute(os.Stdout,data);
}