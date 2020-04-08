package myhttp

import (
	"math/rand"
	"time"
)

// 随机生成str 返回
func GetRandomStr(num int) string {
	strLetter := "qwertyuiopasdfghjklzxcvbnm"
	strNum := "01234567890"
	allStr := strNum + strLetter
	bytes := make([]byte, num)
	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		n := rand.Intn(36)
		bytes[i] = allStr[n]
	}
	return string(bytes)
}

// 生成packageNumber str
func MakePackageNumber() string {
	return GetRandomStr(9)
}

