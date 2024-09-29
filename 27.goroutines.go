package main

import (
	"fmt"
	"time"
)

func f(from string){
	for i:=0;i<3;i++{
		fmt.Println(from," : ",i)
	}
}
func main(){
	f("direct")
	go f("goroutine")
	
	go func(msg string){
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}

// Threads -Managed by OS Fixed stack 1mb
// Goroutines -Managed by Go runtime flexible stack 2kb

// Go runtime schedules and allocates the Goroutines to each core if one cores queue is empty and gobal queue is also empty the goruntime take a goroutine form others cores local queue which not happens in thread

// Goroutines do not communicate by sharing memory instead, share memory by communicatingbetween

