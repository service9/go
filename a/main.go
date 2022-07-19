package main

import (
	"fmt"

	"package.com/book"
)

func main() {
	var info, errors = book.ShowBookInfo("你是猪", "你才是猪")
	fmt.Println(info, errors)

}
