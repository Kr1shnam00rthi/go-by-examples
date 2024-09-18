package main

import "fmt"
	
func main(){
	s[2]="c"
	
	fmt.Println("set: ",s)
	fmt.Println("get: ",s[2])
	fmt.Println("len: ",len(s))
	s=append(s,"d")
	s=append(s,"e","f")
	fmt.Println("len: ",len(s))
	fmt.Println("apd: ",s)
	
	c:=make([]string,len(s))
	copy(c,s)
	fmt.Println("cpy: ",c)
	
	l:=s[2:3]
	fmt.Println("sl1: ",l)
	
	t:= []string{"g","h","i"}
	fmt.Println("dcl:",t)
	
	t2:=[]string{"g","h","i"}
	if slices.Equal(t,t2){
		fmt.Prinln("t==t2")
	}
}

// Difference Between slices and array
// 1.Slices are not fixed size 
// 2.Slices are passed by default as a pass by reference 
// 3.It has very rich built-in function
