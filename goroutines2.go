package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	users := []string{"A", "B", "C"}

	wg := sync.WaitGroup{}
	wg.Add(len(users))

	for _, user := range users {
		go func(u string) { // still bad, should use a pool
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
			fmt.Println(u)
		}(user)
	}

	wg.Wait()
}
