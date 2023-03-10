package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go_blog/common/global"
	"go_blog/model"
	"go_blog/model/other"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
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

func TimeToString(time model.FTime) string {
	return time.Format(DateTime)
}

// GenerateToken 生成token
func GenerateToken(account, password string) (string, error) {
	now := GetNow()
	secret := []byte(global.GlobalViper.GetString("jwt.secret"))

	shouldTime := now.Add(24 * time.Hour)
	claims := Login{
		Account:  account,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			//token 有效期
			IssuedAt:  GetNow().Unix(),
			ExpiresAt: shouldTime.Unix(),
			Issuer:    "go-blog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	return tokenString, err
}

// ParseToken 解析token
func ParseToken(token string) (*Login, error) {
	secret := []byte(global.GlobalViper.GetString("jwt.secret"))
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Login{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Login); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func GetPage(page, limit int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * limit
	}
	return result
}

func ResStruct(code int, message string, data interface{}) other.ResStruct {
	return other.ResStruct{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func ErrToString(err error) string {
	errString := fmt.Sprintf("%s", err)
	return errString
}

func ErrorRes(err error, c *gin.Context) {
	c.JSON(http.StatusOK, ResStruct(201, ErrToString(err), nil))
}
func StringToInt(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}

func MakeExcel[T any](list []T) string {
	// 生成excel
	f := excelize.NewFile()
	// 设置表头
	fieldList := GetStructLabel(list[0])
	f.SetCellValue("sheet1", "A1", "序号")
	for i, v := range fieldList {
		f.SetCellValue("sheet1", fmt.Sprintf("%v%d", string(rune(65+i+1)), 1), v)
	}
	// 设置内容
	for i, v := range list {
		word := rune(66)
		f.SetCellValue("sheet1", fmt.Sprintf("A%d", i+2), i+1)
		for j := 0; j < reflect.TypeOf(v).NumField(); j++ {
			if reflect.TypeOf(reflect.ValueOf(v).Field(j).Interface()).String() != "model.BaseModel" {
				f.SetCellValue("sheet1", fmt.Sprintf("%v%d", string(word), i+2), reflect.ValueOf(v).Field(j).Interface())
				word++
			}
		}
	}
	// 根据指定路径保存文件
	err := f.SaveAs("../file/excel/test.xlsx")
	if err != nil {
		return ""
	}
	return "/file/excel/test.xlsx"

}

func GetStructLabel[T any](target T) []string {
	refType := reflect.TypeOf(target)
	var list []string
	for i := 0; i < refType.NumField(); i++ {
		list = append(list, refType.Field(i).Tag.Get("gorm"))
	}
	var result []string
	for _, v := range list {
		if BetweenString(v, "comment:", ";") != "" {
			result = append(result, BetweenString(v, "comment:", ";"))
		}
	}
	fmt.Println(result)
	return result
}

func BetweenString(str, starting, ending string) string {
	s := strings.Index(str, starting)
	if s < 0 {
		return ""
	}
	s += len(starting)
	e := strings.Index(str[s:], ending)
	if e < 0 {
		return ""
	}
	return str[s : s+e]
}

func InitPage(page, limit *int) {
	if *page <= 0 {
		*page = 1
	}
	if *limit <= 0 {
		*limit = 10
	}
}
