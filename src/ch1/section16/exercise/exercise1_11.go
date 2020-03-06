package exercise

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var done chan string = make(chan string)

// Exercise1_11 ....
func Exercise1_11() {
	input, _ := os.OpenFile("input", os.O_RDONLY, os.ModePerm)
	defer input.Close()
	output, _ := os.OpenFile("output", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	defer input.Close()
	fetchAll(input, output)
}

func fetchAll(input io.Reader, output io.Writer) {
	count := 0
	urls := make([]string, 0, 135)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		urls = append(urls, line)
	}
	ch := make(chan string)
	go readInput()
	var wg sync.WaitGroup
	for _, url := range urls {
		// if cancelled() {
		// 	break
		// }
		wg.Add(1)
		count++
		println("Started for ", url, " ", count)
		go fetch(url, ch, &wg)
	}
	go func() {
		for !cancelled() {
			select {
			case <-done:
				return
			case result := <-ch:
				fmt.Fprintln(output, result)
				println("Ended ", count)
				if count == 1 {
					close(ch)
					return
				}
				count--
			case <-done:
				break
			default:
			}
		}
	}()
	wg.Wait()
	for result := range ch {
		fmt.Fprintln(output, result)
		println("Ended After ", count)
		if count == 1 {
			close(ch)
		}
		count--
	}

}

func fetch(url string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	if cancelled() {
		ch <- fmt.Sprintf("Cancelled\t%s", url)
		return
	}
	if has := strings.HasPrefix(url, "https://www."); !has {
		//way3 Most optimize
		var strBuilder strings.Builder
		strBuilder.WriteString("https://www.")
		strBuilder.WriteString(url)
		url = strBuilder.String()
		//way1
		// fmt.Sprintf("%s%s", "http://", url)
		//way2
		// url = strings.Join([]string{"http://", url}, "")
	}
	start := time.Now()
	resp, er := http.Get(url)
	if er != nil {
		ch <- fmt.Sprintln(er)
		return
	}
	if cancelled() {
		ch <- fmt.Sprintf("Cancelled\t%s", url)
		return
	}
	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintln(er)
		return
	}
	ch <- fmt.Sprintf("%7d\t%.4fs\t%s", nBytes, time.Since(start).Seconds(), url)
}

func readInput() {
	os.Stdin.Read(make([]byte, 1)) // read a single byte
	close(done)
}

func cancelled() bool {
	select {
	case <-done:
		// println("Cancelled -> Done -> True")
		return true
	default:
		// println("Cancelled -> Default -> False")
		return false
	}
}
