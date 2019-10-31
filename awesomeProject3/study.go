package main

import "fmt"

type person struct {
	name   string // 姓名
	age    int    // 年龄
	gender string // 性别
}

type dove interface {
	gugugu() // 鸽
}

type repeater interface {
	repeat(string) // 复读
}
type zhenxiang interface {
	zhenxiang()
}
type ningmengjing interface {
	ningmengjing()
}
func (p *person) repeat(word string) {
	fmt.Println(word)
}

func (p *person) gugugu() {
	fmt.Println(p.name, "又鸽了")
}
func (p*person) zhenxiang(){
	fmt.Println("真香")
}
func (p*person) ningmengjing(){
	fmt.Println("柠檬精")
}
func main() {

	p := &person{
		name:   "lcm",
		age:    18,
		gender: "male",
	}
	p.gugugu()

	var r repeater

	r = p
	r.repeat("helloworld")

	var z zhenxiang
	z = p
	z.zhenxiang()

	var n ningmengjing
	n=p
	n.ningmengjing()
}







