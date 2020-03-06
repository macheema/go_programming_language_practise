package exercise

import (
	"ch1/section14"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var done chan string = make(chan string)

// Exercise1_12 ....
func Exercise1_12() {
	server4()
}

func server4() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:4040", nil))
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	query, er := url.ParseQuery(request.URL.RawQuery)
	if er != nil {
		fmt.Fprint(responseWriter, er)
		return
	}
	cyclesStr := query["cycles"]
	cycles, er := strconv.Atoi(cyclesStr[0])
	if er != nil {
		fmt.Fprint(responseWriter, er)
		return
	}

	section14.Lissajous2(responseWriter, float64(cycles))
}
