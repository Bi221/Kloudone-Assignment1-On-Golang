package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile() {

	file, err := os.Create("library book. count")
	file.WriteString("lirary book")
	file.Close()

	if err != nil {
		log.Fatal(err)
	}

}

func ReadFile() {

	file, err := ioutil.ReadFile("library book.counnt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
}
