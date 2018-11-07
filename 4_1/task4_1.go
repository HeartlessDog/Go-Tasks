package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(
				func() {
					defer fmt.Println("World")
					fmt.Println("Hello")
				})
		}()
	}
	time.Sleep(100 * time.Millisecond)
}