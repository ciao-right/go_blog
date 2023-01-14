package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func BcryptHash(str string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)

}

func GetNow() time.Time {
	return time.Now()
}

func FormatTime(time time.Time, format string) string {
	return time.Format(format)
}
