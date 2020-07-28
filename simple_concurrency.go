package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup //waitgroup
var mut sync.Mutex    //synchronization and locks

func sendRequest(url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	mut.Lock()
	defer mut.Unlock()
	fmt.Printf("[%d] %s\n", res.StatusCode, url)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run main.go <url1> <url2> ... ")
	}

	for _, url := range os.Args[1:] {
		go sendRequest("https://" + url) //each function call gets own routine
		wg.Add(1)
	}
	wg.Wait()
}

//Benchmarks
//#1 4.242s wasting time waiting for response to come back one by one, use goroutines
//#2 1.547s
//#4 1.01s
//synchronization locks because terminal is shared resource
