package main

import "fmt"

const(
	a=25
	b=25
	c=a+b
	x=25.1
	y=25.2
	z=x*y
)
var(
	m=100.0
	n=8.0
	h=m/n
)
func main(){
	fmt.Println(c)
	fmt.Println(z)
	fmt.Println(h)

}