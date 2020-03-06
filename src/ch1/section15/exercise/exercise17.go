package exercise

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//Exercise1_7 ...
func Exercise1_7() {
	fetch2()
}

func fetch2() {
	urls := os.Args[1:]
	for _, url := range urls {
		resp, er := http.Get(url)
		if er != nil {
			fmt.Fprintf(os.Stderr, "fetch1\t%v\n", er)
			os.Exit(1)
		}
		_, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch1\t%v\n", er)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
