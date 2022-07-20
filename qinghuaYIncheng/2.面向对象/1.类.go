package main

import "fmt"

//定义类型，包含数据的属性
type MeiMeiBeiTai struct {
	name     string
	age      int
	tall     float32
	money    float64
	handsome int
}

//数据初始化
func (beitai *MeiMeiBeiTai) setData(name string,
	age int, tall float32, money float64, handsome int) {
	beitai.name = name
	beitai.age = age
	beitai.tall = tall
	beitai.money = money
	beitai.handsome = handsome
}

func (beitai *MeiMeiBeiTai) getData() (string, int, float32, float64, int) { //必须要标识返回的类型
	return beitai.name, beitai.age, beitai.tall, beitai.money, beitai.handsome
}

func (oldBeitai *MeiMeiBeiTai) compare(newBeitai *MeiMeiBeiTai) bool { //返回值类型
	return oldBeitai.money < newBeitai.money
}

func main() {
	fmt.Println("game start")
	//创建beitai
	威威 := new(MeiMeiBeiTai)
	李佳奇 := new(MeiMeiBeiTai)

	威威.setData("zhang", 39, 1.49, 200000.34, 45)
	李佳奇.setData("zhang", 39, 1.49, 200000.44, 45)

	fmt.Println(威威.getData())
	fmt.Println(李佳奇.getData())

	// fmt.Println(威威 < 李佳奇) //<需要重载
	fmt.Println(威威.compare(李佳奇))
}
