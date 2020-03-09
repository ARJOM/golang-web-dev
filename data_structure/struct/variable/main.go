package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sabio struct {
	Nome string
	Lema string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buda := sabio{
		Nome: "Buda",
		Lema: "A crença de não crenças",
	}

	err := tpl.Execute(os.Stdout, buda)
	if err != nil {
		log.Fatal(err)
	}
}
