package test

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

//安装JWT go get github.com/dgrijalva/jwt-go

type UserClaims struct {
	Identity           string `json:"identity"`
	Name               string `json:"name"`
	jwt.StandardClaims        //继承
}

var myKey = []byte("gin-gorm-oj-key")

//生成token
func TestGenerateToken(t *testing.T) {
	UserClaims := &UserClaims{
		Identity:       "user_1",
		Name:           "猪啊啊啊啊",
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tokenString)
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InVzZXJfMSIsIm5hbWUiOiLnjKrllYrllYrllYrllYoifQ.vE6L1vgDz32uj3UBZ43g9MFc6QuUkaz-KkE5LA4ellU
}

//解析token
func TestAnalyseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InVzZXJfMSIsIm5hbWUiOiLnjKrllYrllYrllYrllYoifQ.vE6L1vgDz32uj3UBZ43g9MFc6QuUkaz-KkE5LA4ellU"
	userClaims := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaims, func(t *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if claims.Valid {
		fmt.Println(userClaims)
	}
}

