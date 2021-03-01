package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
)
func main(){
	file, err:= os. Create("kloudone work.text")
	if err!= nil {
		log.Fatal("we got error",err)

	}
	defer file.Close()
	fmt.Fprint(file,"bijedra kloudone golan trainee")

	}
}