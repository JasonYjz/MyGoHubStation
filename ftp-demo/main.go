package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"os"
	"time"
)

func main() {

	addr := "192.168.110.22:21"
	//addr := "192.168.110.227:21"

	fmt.Printf("ready to connect to addr:%v\n", addr)
	conn, err := ftp.Dial(addr, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		fmt.Printf("connect failed, err:%v\n", err)
		return
	}

	//err = conn.Login("ftpuser", "ftpuser")
	fmt.Printf("ready to login with user[%s] and password[%s]\n", "aidlux", "aidlux")

	err = conn.Login("aidlux", "aidlux")
	//err = conn.Login("aidlux", "aidlux123456")
	if err != nil {
		fmt.Printf("login failed, err:%v\n", err)
		return
	}

	currentDir, err := conn.CurrentDir()
	if err != nil {
		fmt.Printf("current dir failed, err:%v\n", err)
		return
	}

	fmt.Printf("current dir:%v\n", currentDir)

	err = conn.ChangeDir("upload")
	if err != nil {
		fmt.Printf("ready to make dir.\n")
		conn.MakeDir("upload")
		return
	}

	// stor file
	//
	file := "C:\\Users\\yul04\\Documents\\Work\\Code\\My\\MyGoHubStation\\ftp-demo\\release.txt"

	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer f.Close()

	err = conn.Stor("test-release.txt", f)
	if err != nil {
		fmt.Printf("stor file failed, err:%v\n", err)
		return
	}

	fmt.Println("stor file success.")
}
