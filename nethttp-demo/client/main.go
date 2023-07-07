package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(256)
	tr := &http.Transport{
		MaxConnsPerHost: 5,
	}
	client := http.Client{
		Transport: tr,
	}
	for i := 0; i < 256; i++ {
		go func(i int) {
			defer wg.Done()
			resp, err := client.Get("http://localhost:5858")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			fmt.Printf("g-%d: %s\n", i, string(body))
		}(i)
	}
	wg.Wait()
}
