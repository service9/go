package book1

import "errors"

func ShowBookInfo1(bookName, authorName string) (string, error) {
	//开头大写
	if bookName == "" {
		return "", errors.New("不能为空")
	}
	if authorName == "" {
		return "123", errors.New("不能为空")
	}
	return bookName + "作者" + authorName, nil
}
