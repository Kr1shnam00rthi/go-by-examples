package main

import "fmt"

// in unbuffered channel make(chan string) there must be a receiver need to get the message passed
// in buffered channel make(chan string,size) the sender holds size of data, till the receiver receives it
func main(){
	messages:=make(chan string,2)
	messages<-"buffered"
	messages<-"channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
