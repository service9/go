package main

//导入 三种方式
import (
	"fmt"
	. "fmt"
	"time"
) //格式化 bin/gofmt ctrl+s

// func newFun() (string, int) {
// 	return "你好啊", 123
// }

// type Phone interface {
// 	call()
// }

// type MyPhone struct{}
// type YourPhone struct{}

// func call() {
// 	Println("woshizhu")
// }
// func (p YourPhone) call() {
// 	Println("nishizhu")
// }

func say(s string, a bool) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		Println(s)
		// Println(a)
	}
}

//main是入口函数
func main() {
	//通道
	//定义
	// var intChan999 intChanan int
	//初始化
	//带缓冲区，第三个参数为缓冲区大小
	intChan := make(chan int, 1)
	// 向通道传值
	intChan <- 3
	close(intChan)
	num, ok := <-intChan
	fmt.Printf("get num: %d, ok: %t\n", num, ok)
	// intChan <- 1
	// 从通道取值
	// var num, num1 = <-intChan, <-intChan
	// Println(num, num1)
	Println(num)

	// say("hello", *new(bool))
	// go say("hello", *new(bool))
	// say("world", *new(bool))

	// // call()
	// // var phone Phone
	// var yourPhone YourPhone
	// // phone = new(MyPhone)
	// // phone.call()
	// // phone = new(YourPhone)
	// // phone.call()
	// yourPhone.call()
	// var phone Phone = yourPhone
	// if phone != nil {
	// 	phone.call()
	// }
	// var you = new(YourPhone)

	// var map1 = make(map[string]string) //map是关键字
	// map1["FrenintChan"] = "巴黎"
	// map1["FrenintChan1"] = "巴黎"
	// Printf("%v", map1)
	// for key := range map1 {
	// 	Printf("\n%v", map1[key])
	// }
	// delete(map1, "FrenintChan")
	// Printf("%v", map1)

	// Printf("%v\n", oldArr)
	// Printf("%v\n", newArr)
	// //range迭代数组切片通道集合
	// for key, value := range oldArr {
	// 	newArr[key] = value
	// }
	// // slice := []int{12}
	// // slice1 := []int{12}
	// // copy(slice, slice1)
	// newArr[0] = 99
	// Printf("%v\n", cap(oldArr))
	// Printf("%v\n", newArr == oldArr)

	// var arr = [...]int{2, 4, 6}
	// var slice = []int{1, 3, 5}
	// var slice1 = arr[:]
	// slice = arr[:]
	// Printf("arr=%v,%T\n", arr, arr)
	// Printf("slice=%v,%T\n", slice, slice)
	// Printf("slice1=%v,%T\n", slice1, slice1)
	// Printf("slice1=%v,%T\n", cap(slice1), cap(slice1))
	// slice1 = slice[:]
	// Printf("slice=%v,%T\n", slice, slice)
	// slice = append(slice, 1, 2)
	// Printf("slice1=%v,%T\n", slice1, slice1)

	//Sizeof存储空间 对string无效
	//查找位置strings.index
	//判断包含strings.Contains
	//合并字符串strings.Join
	//分割字符串strings.Split
	//返回字符串fmt.Sprintf 相当于+
	//字符串长度len()
	//var str = ``//多行字符串

	//输出
	// Println("hello go") //printline 输出并且换行
	// fmt.Print("asf")
	// var a = "123"
	// Printf(a)
	// f.Printf("%v", a)   //格式化输出 推荐
	// f.Printf("%v\n", a) //格式化输出 \n换行

	//占位符 格式化输出
	//%v相对类型 英文转换为ASCLL码 中文Unicode编码十进制
	//%d十进制
	//%T变量类型

	//定义变量并赋值
	// var 变量名 数据类型 =
	// var a int = 3 //变量定义后必须要使用
	//类型自动推导(2)
	// var 变量名 =
	// var a = 3
	// a := 10//只能局部使用

	//定义函数并定义返回类型
	// func 函数名() (返回类型..){函数体}

	// var username, age = newFun()
	// var username, _ = newFun()//匿名变量 不会分配内存给_ 不存在重复声明

	//定义常量 不能更改
	// const u = 3.14
	//iota const a = iota//0//const + iota=0
	// const (
	// 	n1 = iota//0
	// 	_
	// 	n3//2
	// 	n4//3
	// )
	// const (
	// 	n1 = iota//0
	// 	n2 = 100//100
	// 	n3 = iota//2
	// 	n4//3
	// )
	// const (
	// 	n1, n2 = iota + 1, iota + 2
	// 	n3, n4
	// )
}

//go run 运行
//go build 编译
