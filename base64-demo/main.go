package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	strone := "Z2QAHqy0BQHtKQUGBQbQoTU="
	strtwo := "aO4G8sA="

	decodeStringOne, _ := base64.StdEncoding.DecodeString(strone)
	decodeStringTwo, _ := base64.StdEncoding.DecodeString(strtwo)

	fmt.Println(string(decodeStringOne))
	fmt.Println(string(decodeStringTwo))
}
