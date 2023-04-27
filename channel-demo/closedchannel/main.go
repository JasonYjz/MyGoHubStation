package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Data map[string]interface{} `json:"data"`
}


func (p Person) String() string {
	return fmt.Sprintf("Person {Name=%s, Age=%d, Data=%s}", p.Name, p.Age, p.Data["aaa"])
}

func main() {
	ch := make(chan *Person, 100)

	fmt.Println("main is start to run")

	go func() {
		fmt.Println("start to read after 8s.")
		time.Sleep(8 * time.Second)
		for {
			data := <-ch
			if data == nil {
				fmt.Println("the data in channel is read over, exit.")
				return
			}
			time.Sleep(1 * time.Second)
			fmt.Println("the received data is " + data.String())
		}
	}()


	go func() {
		fmt.Println("start to write.")

		age := 10
		mp := make(map[string]interface{})
		mp["aaa"] = "bbb"
		for {
			p := &Person{
				Name: "Test",
				Age:  age,
				Data: mp,
			}
			fmt.Printf("Age[%d] person is writen into ch.\n", age)
			ch <- p
			time.Sleep(1 * time.Second)
			age++

			if age == 15 {
				fmt.Println("close channel.")
				close(ch)
				break
			}
		}

		fmt.Println("end to write.")
	}()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
}

