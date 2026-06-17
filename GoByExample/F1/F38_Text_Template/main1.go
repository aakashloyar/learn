package main

import (
    "text/template"
    "os"
)

func main() {
    // tmpl := template.Must(template.New("hello").Parse("Hello, {{.Name}}!"))
    // data := map[string]string{"Name": "Aakash"}

    // tmpl.Execute(os.Stdout, data)
	// tmpl := template.Must(template.New("cond").Parse(`
	// {{ if .IsAdmin }}
	// Welcome Admin!
	// {{ else }}
	// Welcome User!
	// {{ end }}
	//`))

	// data := map[string]bool{"IsAdmin": true}

	// tmpl.Execute(os.Stdout, data)

	tmpl := template.Must(template.New("list").Parse(`
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

// Breakdown:

// {{ .Name }} means: look up Name field in the data structure.

// . (dot) means “current data context”.