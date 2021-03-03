package main

import "fmt"

func multi(b []int, c chan int) {
	multti := 0
	for _, v := range b {
		add += v
	}
	c <- add // send add to c
}

func main() {
	s := []int{7, 23,-11, 12,-4, 2, 23, -9, 4, 0}

	c := make(chan int)
	go multi(s[:len(s)/2], c)
	go multi(s[len(s)/2:], c)
	x, y, := <-c, <-c, <-c // receive from c

	fmt.Println(x, y, x*y)
}
