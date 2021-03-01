package main

import (
	"encoding/json"
	"fmt"
)

type bookrow struct {
	row  int
	Name string
}

func main() {

	Row, _ := json.Marshal(11)
	fmt.Println(string(Row))

	Name, _ := json.Marshal("library")
	fmt.Println(string(Name))

}
