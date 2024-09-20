package main

import "fmt"


type shape interface{
	area() float64
	perim() float64
} 

type rect struct{
	width float64
	height float64
}

type circle struct{
	radius float64
}

func (r rect) area() float64{
	return r.width*r.height 
}

func (r rect) perim() float64{
	return 2*r.width + 2*r.height
}

func (c circle) area() float64{
	return 3.14*c.radius*c.radius
}

func (c circle) perim() float64{
	return 2*3.14*c.radius
}

func measure(g shape){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){
	r:=rect{1,3}
	c:=circle{10}
	measure(r)
	measure(c)	
}
