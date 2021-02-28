package main

import "fmt"

func main(){
	jobs:= make(chan int,13) 
	result:= make(chan int,13)
}
// 3 workers
for x:= 1;x<=3;x++{
	goworker (x,job,result)
}

give them jobs
for i := 0; i < count; i++ {
	job <-i
}
close (jobs)
for i := 0; i < count; i++ {
	fmt.println("result recived from worker;<-results")
}
}
func workers(id int,jobs <- chan int,results chan <- int){
	for _, v := range v {
		int.println("wrker",id,"is working on job")
		duration:= (time duration(source := rand.NewSource(time.Now().Unix())
		r := rand.New(source)
		fmt.Print(r.Intn(100))))
		time sleep duration
		results <- id
		
	}
}
	
