package main

import "fmt"

type Person struct { //人类
	name string
	age  int
}

type MeiMeiBeiTai struct { //备胎类
	Person      //继承了人类
	personality string
}

func (person Person) getNameAndAge() (string, int) {
	return person.name, person.age
}

func (studentBeiTai MeiMeiBeiTai) getPersonality() string {
	return studentBeiTai.personality
}

func main() {
	zhang := new(MeiMeiBeiTai)
	zhang.name = "张"
	zhang.age = 128
	name, age := zhang.getNameAndAge() //属性方法都能从人类继承
	zhang.personality = "吃饭"
	// personality := zhang.personality
	personality := zhang.getPersonality()
	fmt.Println(name, age)
	fmt.Println(personality)
}
