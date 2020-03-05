package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Antônio Ricart"
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang=pt-br>
	<head>
	<meta charset="UTF-8"/>
	<title>Hello World!</title>
	</head>
	<body>
	<h1> ` + name + `</h1>
	</body>
	</html>
	`)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Erro criando o arquivo", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
