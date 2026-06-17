package main
import (
	//"fmt"
	"text/template"
	"os"
)

func main() {
	tmp:=template.Must(template.New("Hello").Parse(`Hello, {{ .Name}}!`))
	// if(err!=nil) {
	// 	panic(err)
	// }
	data:=map[string]string{"Name":"Aakash"}
	tmp.Execute(os.Stdout,data);
}
//.Must is a helper that wraps a call to a function returning (*Template, error)
// and panics if the error is non-nil. It is intended for use in variable
// initializations such as
//
//	var tmpl = template.Must(template.New("name").Parse("text"))