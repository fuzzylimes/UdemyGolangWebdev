package main

import (
	"fmt"
)

func main() {
	name := "fuzzylimes"
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hell World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>`
	fmt.Println(tpl)
}
