package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//var ser server
	ser := make(map[string]interface{})
	dir, _ := os.Getwd()
	data, err := os.ReadFile(dir + "\\intinterface-demo\\" + "test.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = json.Unmarshal(data, &ser); err != nil {
		fmt.Printf("unmarshal json fail, err: %s\n", err)
		return
	}

	port := 21

	if port == ser["port"] {
		fmt.Println("equal")
	} else {
		fmt.Println("no equal")
	}
}

type server struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}
