package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	DateTime = "2006-01-02 15:04:05"
)

type Login struct {
	Account  string `validate:"required"`
	Password string `validate:"required"`
	jwt.StandardClaims
}

func BcryptHash(str string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetNow() time.Time {
	return time.Now()
}

func FormatTime(time time.Time, format string) string {
	return time.Format(format)
}

// GenerateToken 生成token
func GenerateToken(account, password string) (string, error) {
	now := GetNow()
	shouldTime := now.Add(24 * time.Hour)
	claims := Login{
		Account:  account,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			//token 有效期
			ExpiresAt: shouldTime.Unix(),
			Issuer:    "go-blog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SigningString()
	return tokenString, err
}

// ParseToken 解析token
func ParseToken(token string) (*Login, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Login{}, func(token *jwt.Token) (interface{}, error) {
		return []string{viper.GetString("jwt.secret")}, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Login); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
