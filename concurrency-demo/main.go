package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//但是开启100条小协程是主协程干的
	//迅速开启100条小协程
	for i:=0;i<100000;i++{
		go doSomething("面包人"+strconv.Itoa(i))
	}

	for{
		fmt.Println("我是主协程")
		time.Sleep(time.Second)
	}
}

func doSomething(grname string)  {
	for{
		fmt.Println("来了一车",grname)
		time.Sleep(time.Second)
	}
}