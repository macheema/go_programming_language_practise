package section17

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutex sync.Mutex
var count int

//Server2 ..
func Server2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count/", counter)
	log.Fatal(http.ListenAndServe("localhost:4040", nil))
}

func handler2(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("All Hanlder Called")
	mutex.Lock()
	count++
	mutex.Unlock()
	fmt.Fprintf(responseWriter, "URL.Path = %q\n", request.URL.Path)
}

func counter(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Counter Called")
	mutex.Lock()
	fmt.Fprintf(responseWriter, "Counter = %v\n", count)
	mutex.Unlock()
}
