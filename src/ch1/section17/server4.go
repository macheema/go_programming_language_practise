package section17

import (
	"ch1/section14"
	"log"
	"net/http"
)

//Server4 ..
func Server4() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count/", counter)
	http.HandleFunc("/lissajous/", func(responseWriter http.ResponseWriter, request *http.Request) {
		section14.Lissajous(responseWriter)
	})
	log.Fatal(http.ListenAndServe("localhost:4040", nil))
}
