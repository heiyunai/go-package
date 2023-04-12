package utils

import (
	"math/rand"
	"strconv"
)

func GenVerifyCode() string {
	//1.生成随机6位数
	return strconv.Itoa(rand.Intn(899999) + 100000)
}
