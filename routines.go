package main

import (
	"fmt"
	"time"
)

func say(golang string) {
	for i := 0; i < 3; i++ {
		time.Sleep(33 * time.Millisecond)
		fmt.Println(golang)
	}
}

func main() {
	go say("bijendra")
	say("hello, world")
}
