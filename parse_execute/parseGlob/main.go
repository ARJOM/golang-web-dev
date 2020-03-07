package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	tpl, err := template.ParseGlob("template/*.gohtml")
	check(err)

	err = tpl.Execute(os.Stdout, nil)
	check(err)

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	check(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	check(err)

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	check(err)

}
