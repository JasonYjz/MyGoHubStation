package main

import "fmt"

func main() {
	var args = 1

	p := &args
	fmt.Printf("原始指针的内存地址时：%p\n", &p) //存放指针类型变量
	fmt.Printf("原始指针指向变量的内存地址时：%p\n", p) //存放int变量

	modifyPointer(p)
	fmt.Printf("修改后的值：%v\n", *p)
}

func modifyPointer(p *int) {
	fmt.Printf("函数里接收到指针的内存地址时：%p\n", &p)
	fmt.Printf("函数里接收到指针指向变量的内存地址时：%p\n", p)
	*p = 10
}
