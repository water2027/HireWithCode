package utils

import (
	"crypto/rand"
	"log"
	"strconv"
)

func GetRandomCode() string {
	const min = 100000
	const max = 999999

	// 使用 crypto/rand 生成4个随机字节
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		log.Fatal("生成验证码失败:", err)
	}
	// 将随机结果转换为正整数，并计算出对应范围内的数字
	n := int(b[0])<<24 | int(b[1])<<16 | int(b[2])<<8 | int(b[3])
	if n < 0 {
		n = -n
	}
	code := n%(max-min+1) + min
	return strconv.Itoa(code)
}
