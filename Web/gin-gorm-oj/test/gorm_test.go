package test

import (
	"fmt"
	"gin_gorm_oj/models"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//测试
func TestGormTest(t *testing.T) {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "go:010911@tcp(106.54.177.233:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.Problem, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("problem==>%v>?????????????????????????", v)
	}
}
