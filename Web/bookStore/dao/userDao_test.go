package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("开始测试User")
	t.Run("testLogin:", testLogin)
	t.Run("testUserName:", testUserName)
	t.Run("testSave:", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckLogin("admin1", "123456")
	fmt.Println("获取用户信息是:", user)
}

func testUserName(t *testing.T) {
	user, _ := CheckUserName("admin1")
	fmt.Println("获取用户信息是:", user)
}

func testSave(t *testing.T) {
	err := Saveuser("admin1", "123456", "a@a.afff")
	fmt.Println("获取用户信息是:", err)
}
