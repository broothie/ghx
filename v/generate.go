//go:build generate

package main

import (
	_ "embed"
	"os"
	"strings"
	"text/template"

	"github.com/samber/lo"
)

//go:embed tags.go.tmpl
var tagsTemplate string

//go:generate go run generate.go
func main() {
	generateTags()
}

func generateTags() {
	tagsBytes, err := os.ReadFile("tags.txt")
	if err != nil {
		panic(err)
	}

	tags := lo.Filter(strings.Split(string(tagsBytes), "\n"), func(s string, _ int) bool { return s != "" })
	tmpl := template.Must(template.New("tags.go.tmpl").Funcs(funcMap()).Parse(tagsTemplate))

	file, err := os.OpenFile("tags.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	if err := tmpl.Execute(file, tags); err != nil {
		panic(err)
	}
}

func funcMap() template.FuncMap {
	return template.FuncMap{"toLower": strings.ToLower}
}
