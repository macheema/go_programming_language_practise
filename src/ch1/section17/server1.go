package section17

import (
	"fmt"
	"log"
	"net/http"
)

//Server1 ..
func Server1() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:4040", nil))
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "URL.Path = %q\n", request.URL.Path)
}
