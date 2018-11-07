package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup 
    wg.Add(10)  
	var once sync.Once
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			once.Do(
				func() {
					defer fmt.Println("World")
					fmt.Println("Hello")
				})
		}()
	}
	wg.Wait() 
}