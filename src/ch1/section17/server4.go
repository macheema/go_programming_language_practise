package section17

import (
	"ch1/section14"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//Server4 ..
func Server4() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count/", counter)
	http.HandleFunc("/print/", writeRequestData)
	http.HandleFunc("/lissajous/", func(responseWriter http.ResponseWriter, request *http.Request) {
		printLissajous(responseWriter, request)
	})
	log.Fatal(http.ListenAndServe("localhost:4040", nil))
}

func printLissajous(resp http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprint(resp, fmt.Errorf("%v", err))
	}
	if cycles, ok := req.Form["cycles"]; ok && len(cycles) > 0 {
		if cycle, err := strconv.ParseFloat(cycles[0], 32); err != nil {
			fmt.Fprint(resp, fmt.Errorf("%v", err))
		} else {
			section14.Lissajous2(resp, cycle)
		}
	} else {
		section14.Lissajous1(resp)
	}
}
