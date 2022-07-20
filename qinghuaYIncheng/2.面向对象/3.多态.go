package main

import "fmt"

type qinshou interface { //接口不变，行为拓展
	run()
	live()
}

type yiguanqinshou interface {
	run()
}

type Person struct {
}

func (p Person) run() {
	fmt.Println("Person跑---")
}

func (p Person) live() {
	fmt.Println("Person活着---")
}

type Tiger struct {
}

func (t Tiger) run() {
	fmt.Println("Tiger跑---")
}

func (t Tiger) live() {
	fmt.Println("Tiger活着---")
}

func main() {
	var qs qinshou
	var ygqs yiguanqinshou
	zhang := new(Person)
	qs = zhang
	qs.run()
	qs.live()
	ygqs = qs
	ygqs.run()
	// ygqs.live()没有这个方法 多态

	老虎 := new(Tiger)
	qs = 老虎
	qs.run()
	qs.live()
}
