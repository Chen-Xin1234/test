//package main
//
//import (
//	"fmt"
//)
//
//func printbites(s string){
//	for i:=0;i<len(s);i++{
//		fmt.Printf("%x\n",s[i])
//	}
//}
//func printchar(s string){
//	runes:=[]rune(s)
//	for i:=0;i<len(runes);i++{
//		fmt.Printf("%c",runes[i])
//	}
//}
//func main() {
//	name := "Hello World"
//	printbites(name)
//	printchar(name)
//	name1:="Señor"
//	printbites(name1)
//	printchar(name1)
//}
//package main
//import "fmt"
//func main(){
//	biteslice:=[]rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
//	str:=string(biteslice)
//	fmt.Print(str)
//}
//package main
//import("fmt"
//)
//
//func length(s []rune) string{
//	    s[0]='a'
//		return string(s)
//}
//func main(){
//	name:="hello"
//	fmt.Println(length([]rune(name)))
//}
//package main
//import "fmt"
//func main(){
//	b:=255
//	var a*int=&b
//	fmt.Println("output a: ",a)
///}
//package main
//import "fmt"
//func main(){
//	b:=233
//	var a*int
//	if a==nil {
//		fmt.Println("b is",a)
//		a=&b
//		fmt.Println("b is",a)
//	}
//
//}
//package main
//import "fmt"
//func main(){
//	b:=255
//	a:=&b
//	fmt.Println("the vablue of b is",*a)
//	fmt.Println("the address of b is",a)
//	*a++
//	fmt.Print(b)
//}
//package main
//import "fmt"
//func change(m*int){
//	*m=55
//}
//func main(){
//	a:=54
//	fmt.Print(a)
//	fmt.Print("\n")
//	b:=&a
//	change(b)
//	fmt.Print(a)
//}
//package main
//import "fmt"
//func change(arry*[3]int){
//	arry[0]=1  //a[x] 是 (*a)[x] 的简写形式，因此上面代码中的 (*arr)[0] 可以替换为 arr[0]。下面我们用简写形式重写以上代码
//}
//func main(){
//	a:=[3]int{234,12,1}
//	fmt.Print(a)
//	fmt.Print("\n")
//	change(&a)
//	fmt.Print(a)
//}
//package main
//
//import (
//	"fmt"
//)
//
//func modify(sls []int) {
//	sls[0] = 90
//}
//
//func main() {
//	a := [3]int{89, 90, 91}
//	modify(a[:])
//	fmt.Println(a)
//}
//package main
//import "fmt"
//type employee struct{
//	firstname,lastname string
//	age,salary int
//}
//
//func main(){
//	employee1:=employee{
//		firstname: "ada",
//		lastname: "mith",
//		age:        25,
//		salary:    2000,
//	}
//	employee2:=employee{"Thomas", "Paul", 29, 800}
//	fmt.Println("Employee 1", employee1)
//	fmt.Println("Employee 2", employee2)
//}
//package main
//import "fmt"
//func main(){
//	emp3:=struct{
//		firstname,lastname string
//		age,salary int
//	}{
//		firstname:"ada",
//		lastname:"panada",
//		age:13,
//		salary:1000,
//	}
//	fmt.Println("Employ 3 ",emp3)
//}
//package main
//import "fmt"
//type Employee struct {
//	firstname,lastname string
//	age,salary int
//}
//func main(){
//	var emp4 Employee
//	emp4.firstname="ada"
//	emp4.lastname="panada"
//
//	fmt.Println("Employee:",emp4)
//}
//package main
//import "fmt"
//type Employee struct {
//	 string
//	 int
//}
//func main(){
//	emp5:=&Employee{"hio",6000}
//	fmt.Println("epm5:",(*emp5))
//}
//package main
//import "fmt"
//type Address struct{
//	city,state string
//}
//type person struct {
//	age int
//	name string
//	Address
//}
//func main(){
//	var p person
//	p.name="smith"
//	p.age=56
//	p.Address=Address{
//		city:  "ChongQing",
//		state: "China",
//	}
//	fmt.Println("city:",p.city)
//}
//package main
//import "fmt"
//type Employee struct {
//	name string
//	salary int
//	currency string
//}
//func displaysalary(a Employee) {
//	fmt.Printf("Salary of %s is %s%d", a.name, a.currency, a.salary)
//}
//func main(){
//	a:=Employee{
//		name:     "smith",
//		salary:   800,
//		currency: "$",
//	}
//	displaysalary(a)
//}
package main

import "fmt"

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}