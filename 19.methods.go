package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5} // Correct initialization of struct
	fmt.Println("area: ", r.area())  // Using the value receiver
	fmt.Println("perim: ", r.perim())

	rp := &r                      // Pointer to the struct
	fmt.Println("area: ", rp.area())  // Using the pointer receiver
	fmt.Println("perim: ", rp.perim()) // Pointer automatically dereferenced
}
 
