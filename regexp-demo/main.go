package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 输入字符串
	input := "Hello, the price is $123.45,\n and the quantity is 100."

	// 创建一个正则表达式模式，用于匹配数字
	pattern := "\\d+"
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("正则表达式编译失败:", err)
		return
	}

	// 使用正则表达式查找所有匹配的数字
	matches := regex.FindAllString(input, -1)

	// 打印找到的数字
	for _, match := range matches {
		fmt.Println("找到的数字:", match)
	}
}
