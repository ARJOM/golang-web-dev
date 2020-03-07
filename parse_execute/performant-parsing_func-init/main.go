package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	check(err)

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	check(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	check(err)

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	check(err)
}
