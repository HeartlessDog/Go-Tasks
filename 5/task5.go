package main

import (
    "fmt"
    "net/http"
    "time"
	"flag"
)
const timeout=-1

func main() {
	url := flag.String("url", "https://google.com", "a string")
    numb := flag.Int("number", 10, "an int")
	timeout:= flag.Int("timeout", 500, "a int")
	var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	ch := make(chan float64)
	var timeStorage []float64

	client := http.Client{
		Timeout: time.Duration(*timeout) * time.Millisecond,
	}

	start := time.Now()

	for i:=0; i < *numb; i++ {
		go MakeRequest(*url, ch, client)
	}
	for i:=0; i < *numb; i++ {
		timeStorage=append(timeStorage, <-ch)
	}

	duration := time.Since(start).Seconds()

	if timeStorage[0] != -1 {
		minTime:=timeStorage[0]
		maxTime:=timeStorage[len(timeStorage)-1]
		end:=len(timeStorage)-1
		sum:=0.0

		for i,v := range timeStorage{
			if v!=-1{
				sum+=v
			fmt.Printf("%v request - %.4fs\n",i+1, v)
			} else {
				maxTime=timeStorage[i-1]
				end=i-1
				break
			}	
		}	
		averageTime:=sum/float64(end+1)
		notWait:=len(timeStorage)-end-1
	
		fmt.Printf("\nElapsed time - %.4fs\n", duration)
		fmt.Printf("Average request time - %.4fs\n", averageTime)
		fmt.Printf("Max response time - %.4fs\n", maxTime)
		fmt.Printf("Min response time - %.4fs\n", minTime)
		fmt.Printf("No response for %v requests\n\n", notWait)
  	} else {
		fmt.Printf("\nNo responce. Сheck url or increase timeout.\n\n")
  }
}

func MakeRequest(url string, ch chan<-float64, client http.Client) {
		start := time.Now()
		response, err := client.Get(url)
		if err != nil {
			ch <- timeout
			return
		}
		defer response.Body.Close()
		duration := time.Since(start).Seconds()
		ch <- duration
}
  
