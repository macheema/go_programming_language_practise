package section17

import (
	"fmt"
	"log"
	"net/http"
)

//Server3 ..
func Server3() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count/", counter)
	http.HandleFunc("/print/", writeRequestData)
	log.Fatal(http.ListenAndServe("localhost:4040", nil))
}

func writeRequestData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, fmt.Errorf("%v", err))
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %v\n", k, v)
	}
}
