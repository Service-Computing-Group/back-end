package token

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const (
	Secretkey = "123456"
)

func GetToken(r *http.Request) (token *jwt.Token, err error) { //由request获取token
	t := request.AuthorizationHeaderExtractor
	// t是已经实现extract接口的对象，对request进行处理得到tokenString并生成为解密的token
	// request.ParseFromRequest的第三个参数是一个keyFunc，具体的直接看源代码
	// 该keyFunc参数需要接受一个“未解密的token”，并返回Secretkey的字节和错误信息
	// keyFunc被调用并传入未解密的token参数，返回解密好的token和可能出现的错误
	// 若解密是正确的，那么返回的token.valid = true
	return request.ParseFromRequest(r, t,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(Secretkey), nil
		})
}
