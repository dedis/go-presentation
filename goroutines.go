package main

import (
	"fmt"
	"time"
)

func main() {
	users := []string{"A", "B", "C"}
	for _, user := range users {
		go func() {
			time.Sleep(time.Millisecond * 100)
			fmt.Println(user)
		}()
	}
	time.Sleep(time.Second)
}
