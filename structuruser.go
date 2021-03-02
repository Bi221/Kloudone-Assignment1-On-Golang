package main

import "fmt"

type student struct {
	Name string
	roll int
	mark string
}

func main() {
	var v1 student
	v1.Name = "bijendra"
	v1.roll = 12
	v1.mark = "98"
	fmt.Println(v1)
}
