package main

import "fmt"

func main(){
	
	nums:= []int{1,2,3,4}
	sum:=0
	
	for _,num := range(nums){
		sum+=num
	}
	fmt.Println("Sum:",sum)
	
	for i,num :=range(nums){
		if num==3{
			fmt.Println("index",i)
		}	
	}
	m:= map[int]int{1:1,2:2,3:3}
	for k,v := range(m){
		fmt.Println(k,"->",v)
		
	}
	for i:= range m{
		fmt.Println("keys: ",i)
	}
}
