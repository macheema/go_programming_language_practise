package section15

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Fetch1 ..
func Fetch1() {
	urls := os.Args[1:]
	for _, url := range urls {
		resp, er := http.Get(url)
		if er != nil {
			fmt.Fprintf(os.Stderr, "fetch1\t%v\n", er)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch1\t%v\n", er)
			os.Exit(1)
		}
		resp.Body.Close()
		fmt.Printf("%s", body)
	}
}
