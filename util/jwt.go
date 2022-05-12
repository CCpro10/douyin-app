package util

import (
	"douyin-app/conf"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	//token秘钥
	mySecret []byte

	//token过期时间
	tokenExpireDuration time.Duration
)

func InitJWTVal() {
	tokenExpireDuration = time.Duration(int64(conf.Config.Jwt.TokenExpireDuration) * int64(time.Hour))
	mySecret = []byte(conf.Config.Jwt.Secret)
}

// MyClaims自定义声明结构体并内嵌jwt.StandardClaims
// 这里额外记录两个字段
type MyClaims struct {
	Id     int64 //用户唯一自增主键ID
	UserId int64 //表示用户业务ID
	jwt.StandardClaims
}

//生成token,传入ID,userId,生成JWTString和err
func GenerateToken(Id int64, UserId int64) (string, error) {
	// 创建一个自己的声明/请求
	c := MyClaims{
		Id,
		UserId, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireDuration).Unix(), // 过期时间
			Issuer:    "douyin-app",                               // 签发人
			Subject:   "user token",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的秘钥签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// 解析tokenString，返回一个包含信息的用户声明
func ParseToken(tokenString string) (*MyClaims, error) {
	// 通过(tokenString,请求结构,返回秘钥的一个回调函数)这三个参数,返回一个token结构体
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}

	// 校验token,token有效则返回myClaims请求
	if myClaims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return myClaims, nil
	}
	//token无效，返回错误
	return nil, errors.New("invalid token")
}
