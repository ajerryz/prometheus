package main

import (
	"net/http"
	"strconv"
	"time"
)

func main() {

	for {
		for i := 0; i < 50; i++ {
			go func() {
				_, _ = http.Get("http://localhost:8788/user/list")
			}()
		}
		for i := 0; i < 10; i++ {
			go func() {
				_, _ = http.Get("http://localhost:8788/user/create")
			}()
		}
		for i := 0; i < 30; i++ {
			go func() {
				_, _ = http.Get("http://localhost:8788/user/user/" + strconv.Itoa(i))
			}()
		}
		time.Sleep(time.Second * 3)

	}
}
