package main

import "fmt"

func main() {
	name := "Antônio Ricart"

	tpl := `
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
	`

	fmt.Println(tpl)
}
