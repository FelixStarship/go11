package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义结构体（继承jwt.StandardClaims 结构体）
type MyClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// 设置常量 定义过期时间时使用
const (
	TokenExpireDuration = time.Hour * 2
)

// 定义签名
var MySecret = []byte("Red B")

// 生成token
func GenToken(username string) (string, error) {
	//新建结构体
	c := MyClaims{
		Username: username, //用户名
		Role:     "rwt",    //角色权限等。。。。
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "my-project",                               //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(MySecret)
}

// 解析token （返回的是MyClaims的结构体）
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
func main() {
	token, err := GenToken("momo")
	if err != nil {
		fmt.Println(err)
		return
	}
	//验证生成token
	fmt.Println(token)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>开始解析token")
	parseToken, err := ParseToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := json.Marshal(parseToken)
	fmt.Println(string(data))

}
