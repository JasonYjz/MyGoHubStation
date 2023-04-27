package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
	fmt.Fprintf(conn, "who?\n")
	data, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func process(conn net.Conn) {
	// 处理完关闭连接
	defer func() {
		fmt.Println("close the conn")
		conn.Close()
	}()

	fmt.Println("process data on this conn")
	// 针对当前连接做发送和接受操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}

		recv := string(buf[:n])
		fmt.Printf("收到的数据：%v\n", recv)

		// 将接受到的数据返回给客户端
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			fmt.Printf("write from conn failed, err:%v\n", err)
			break
		}
	}
}

func processConn(conn net.Conn) {
	// 处理完关闭连接
	defer func() {
		fmt.Println("close the conn")
		conn.Close()
	}()

	fmt.Println("process data on this conn")
	// 针对当前连接做发送和接受操作
	for {
		reader := bufio.NewReader(conn)
		n, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}

		fmt.Printf("the received data length is：%d\n", len(n))
		//fmt.Printf("the received data length is：%d\n", len(n))

		// 将接受到的数据返回给客户端
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			fmt.Printf("write from conn failed, err:%v\n", err)
			break
		}
	}
}

func main() {
	fmt.Println("the tcp server is listening at :2300")

	l, err := net.Listen("tcp", ":2300")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//sigc := make(chan os.Signal, 1)
	//signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	for {
		//select {
		//case <-sigc:
		//	break
		//	case <-time.Tick(5 * time.Microsecond):
		//		continue
		//
		//default:
		//}

		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("===================================")
		fmt.Printf("a new conn is established. add:%s\n", conn.RemoteAddr().String())
		//go handleConnection(conn)
		go processConn(conn)
	}
}
