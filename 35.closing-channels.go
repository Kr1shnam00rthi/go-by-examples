package main

import (
	"fmt"
	"time"
)

func main(){

	messages:= make(chan string)
	go func(){
		for i:=1;i<=5;i++{
			messages<-fmt.Sprintf("message %d",i)
			time.Sleep(time.Millisecond*100)
		}
		close(messages) // close the channel when done
	}()
	
	for msg:=range messages{
		fmt.Println("Received: ",msg)
	}
	_ , ok:=<-messages
	fmt.Println(ok)
	fmt.Println("No more messages")
}
