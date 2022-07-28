package models

import (
	"time"
)

//定义公共的模块

//自定义的函数 时间戳转换为日期
func UnixToTime(timestamp int) string { //首字母大写公共方法
	t := time.Unix(int64(timestamp), 0) //ms,μs
	return t.Format("2006-01-02 15:05:04")
}

//获取年月日
func GetDay() string {
	template := "20060103"
	return time.Now().Format(template)
}
