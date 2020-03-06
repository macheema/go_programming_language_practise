package section16

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//FetchConcurent ..
func FetchConcurent() {
	urls := os.Args[1:]
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		result := <-ch
		fmt.Println(result)
	}
	fmt.Printf("%.2fs total time eplased\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, er := http.Get(url)
	if er != nil {
		ch <- fmt.Sprintf("Error\t%s\t%.2fs eplased for\t%s", fmt.Sprint(er), time.Since(start).Seconds(), url)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("Error while Reading\t%s\t%.2fs eplased for\t%s", fmt.Sprint(err), time.Since(start).Seconds(), url)
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d eplased for %v", time.Since(start).Seconds(), nbytes, url)
}
