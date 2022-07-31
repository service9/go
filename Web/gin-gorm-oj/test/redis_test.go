package test

import (
	"context"
	"gin_gorm_oj/models"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

//go get github.com/go-redis/redis/v8
//安装
//启动redis

//模拟网络请求会话
var ctx = context.Background()

//redis连接
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0, //default DB
})

func TestRedisSet(t *testing.T) {
	rdb.Set(ctx, "name", "1", time.Second*3)
}

func TestRedisGet(t *testing.T) {
	value, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		t.Fatal(err)
	}
	println(value)
}

func TestRedisGetByModels(t *testing.T) {
	value, err := models.RDB.Get(ctx, "name").Result()
	if err != nil {
		t.Fatal(err)
	}
	println(value)
}
