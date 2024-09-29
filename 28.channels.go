package main

import "fmt"

var message = make(chan string)

func tomain() {
	message <- "Hello from goroutine"
}

func frommain() {
	msg := <-message
	fmt.Println(msg)
}

func main() {
	go tomain()
	msg := <-message
	go frommain()
	fmt.Println(msg)
	message <- "Hello from main!"
}

