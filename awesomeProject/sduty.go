package main

import (
	"fmt"
)


func main() {
	num_S := []int{1,2, 3, 4}
	var num_Count  int
	for i,v := range(num_S){
		for y,m := range(num_S){
			for z,n := range(num_S){
				if (i!=z) && (i!=y) && (y!=z){
					fmt.Println(v*100+m*10+n)
					num_Count ++
				}
			}
		}
	}
	fmt.Println("总数为：", num_Count)
}