package main

import (
	"fmt"
	"github.com/go-stomp/stomp"
	"net"
	"strconv"
	"time"
)

type ecuMsg struct {
	Type string `json:"type"`
	Data interface{} `json:"data"`
}

func ConnectActiveMQ(host, port string) *stomp.Conn {

	stompConn, _ := stomp.Dial("tcp", net.JoinHostPort(host, port))
	return stompConn
}

func Producer(c chan string, destination string, conn *stomp.Conn) {
	for {
		err := conn.Send(destination, "text/plain", []byte(<-c))
		if err != nil {
			fmt.Println("ActiveMq message send error: " + err.Error())
		}
	}
}

func Consumer(destination string, conn *stomp.Conn) {

	sub, _ := conn.Subscribe(destination, stomp.AckAuto)
	for {
		select {
		case v := <-sub.C:
			fmt.Println(string(v.Body))
		case <-time.After(time.Second * 30):
			return
		}
	}
}

func main() {
	host := "192.168.110.101"
	port := "61613"
	topic := "/topic/toEcuTopic"

	activeMq := ConnectActiveMQ(host, port)
	defer activeMq.Disconnect()

	c := make(chan string)

	go Consumer(topic, activeMq)
	time.Sleep(time.Second)
	go Producer(c, topic, activeMq)


	for i := 0; i < 100; i++ {
		c <- "this is a msg " + strconv.Itoa(i)
	}
	time.Sleep(time.Second)
}