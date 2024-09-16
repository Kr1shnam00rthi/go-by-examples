package main

import "fmt"

func sum(nums ... int){
	fmt.Println(nums)
	total:=0
	for i:=0;i<len(nums);i++{
		total+=nums[i]
	}
	fmt.Println(total);
}

func main(){
	sum(1,2)
	sum(1,2,3)
	
	nums:=[]int{1,2,3,4}
	
	sum(nums...)
}


