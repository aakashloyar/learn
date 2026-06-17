package main

import (
    "text/template"
    "os"
)

func main() {
    
	tmpl := template.Must(template.New("cond").Parse(`
	{{ if .IsAdmin }}
	Welcome Admin!
	{{ else }}
	Welcome User!
	{{ end }}
	`))

	data := map[string]bool{"IsAdmin": true}

	tmpl.Execute(os.Stdout, data)
}

// Breakdown:

// {{ .Name }} means: look up Name field in the data structure.

// . (dot) means “current data context”.