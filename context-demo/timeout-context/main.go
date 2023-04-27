package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx1, cancelFunc1 := context.WithTimeout(context.Background(), 5*time.Second)
	ctx2, cancelFunc2 := context.WithTimeout(ctx1, 3*time.Second)

	fmt.Println("main is start.")

	//defer func() {
	//	fmt.Println("execute cancelFunc1")
	//	cancelFunc1()
	//}()
	defer func() {
		fmt.Println("execute cancelFunc2")
		cancelFunc2()
	}()
	//
	go func() {
		fmt.Println("execute cancelFunc1 after 1s")
		time.Sleep(1 * time.Second)

		cancelFunc1()
	}()


	select {
	//case <-ctx1.Done():
	//	fmt.Println("5s timeout ctx1 is done.")
	case <-ctx2.Done():
		fmt.Println("3s timeout ctx2 is done. err:" + ctx2.Err().Error())
		deadline, ok := ctx2.Deadline()
		fmt.Printf("deadline:%s, ok:%b\n", deadline.String(), ok)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
}
