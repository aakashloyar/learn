package main

import (
    "text/template"
    "os"
)

func main() {
    
	tmpl := template.Must(template.New("li").Parse(`
	Users:
	{{ range .Users }}
	- {{ . }}
	{{ end }}
	`))

	data := map[string][]string{
		"Users": {"Aakash", "John", "Sara"},
	}

	tmpl.Execute(os.Stdout, data)
}
