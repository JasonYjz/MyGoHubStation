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
	ctx, cancelFunc := context.WithCancel(context.Background())
	//ctx, cancelFunc := context.WithTimeout(context.Background(), 5 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		cancelFunc()
		fmt.Println("cancelFunc is executed")
	}()

	//go func() {
	//	fmt.Println("will execute cancel func after 3s")
	//	time.Sleep(10 * time.Second)
	//
	//}()

	fmt.Println("wait for context is cancel.")
	//select {
	//case <-ctx.Done():
	//	fmt.Println("ctx is done. err:" + ctx.Err().Error())
	//}
	go func() {
		i := 0
		for {
			i++
			v, ok := <-ctx.Done()
			fmt.Printf("[%d] -> value is %v, ok is %b\n", i, v, ok)
			time.Sleep(1 * time.Second)
		}
	}()


	fmt.Println("check ctx.Done, maybe is blocked.")

    <-ctx.Done()
	fmt.Println("ctx is done. err:" + ctx.Err().Error())

	<-ctx.Done()
	fmt.Println("ctx is done. err:" + ctx.Err().Error())

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
}
