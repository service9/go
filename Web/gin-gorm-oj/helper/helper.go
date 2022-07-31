package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
)

type UserClaims struct {
	Identity           string `json:"identity"`
	Name               string `json:"name"`
	IsAdmin            int    `json:"is_admin"`
	jwt.StandardClaims        //继承
}

// GetMd5
// 生成md5
func GetMd5(s string) string {
	//%x 16进制
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("gin-gorm-oj-key")

// GenerateToken
// 生成token
func GenerateToken(identity, name string, isAdmin int) (string, error) {
	UserClaims := &UserClaims{
		Identity:       identity,
		Name:           name,
		IsAdmin:        isAdmin,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaims := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaims, func(t *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("Analyse Token Error:%v", err)
	}
	return userClaims, nil
}

// SendEmailCode
// 邮箱发送验证码 单条
func SendEmailCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "你是猪 <ljj001210@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "猪子，请查收验证码！" //主题
	e.HTML = []byte("您的验证码是:<b>" + code + "</b>")
	// err := e.Send("smtp.163.com:465", smtp.PlainAuth("", "ljj001210@163.com", "AGWTEXDWASMZDRBI", "smtp.163.com"))
	return e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "ljj001210@163.com", "AGWTEXDWASMZDRBI", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}

// GetUUid
// 生成唯一码
func GetUUid() string {
	return uuid.NewV4().String()
}

// GetRand
// 随机生成验证码
func GetRand() string {
	//种子
	rand.Seed(time.Now().UnixNano())
	//rand.Intn(10) //生成0-9的数
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}
