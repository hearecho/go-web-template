package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	url := "http://127.0.0.1:8081/api/v1/tags?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjEyMyIsInBhc3N3b3JkIjoiNDU2IiwiZXhwIjoxNTk4MDA4MTM5LCJpc3MiOiJnby13ZWIifQ.C4yfHKuOhK3EyyW2h330tlnY2YNnXP75lW9ksFCjarg"
	wg.Add(8)
	go func() {
		resp,_ := http.Get(url)
		fmt.Println(resp.StatusCode)
		wg.Done()
	}()
	go func() {
		resp,_ := http.Get(url)
		fmt.Println(resp.StatusCode)
		wg.Done()
	}()
	go func() {
		resp,_ := http.Get(url)
		fmt.Println(resp.StatusCode)
		wg.Done()
	}()
	go func() {
		resp,_ := http.Get(url)
		fmt.Println(resp.StatusCode)
		wg.Done()
	}()
	go func() {
		resp,_ := http.Get(url)
		fmt.Println(resp.StatusCode)
		wg.Done()
	}()
	go func() {
		resp,_ := http.Get(url)
		fmt.Println(resp.StatusCode)
		wg.Done()
	}()
	go func() {
		resp,_ := http.Get(url)
		fmt.Println(resp.StatusCode)
		wg.Done()
	}()
	go func() {
		resp,_ := http.Get(url)
		fmt.Println(resp.StatusCode)
		wg.Done()
	}()
	wg.Wait()

}
