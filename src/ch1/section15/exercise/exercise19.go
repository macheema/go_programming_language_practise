package exercise

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//Exercise1_9 ...
func Exercise1_9() {
	fetch3()
}

func fetch4() {
	urls := os.Args[1:]
	for _, url := range urls {
		if has := strings.HasPrefix(url, "http://"); !has {
			//way3 Most optimize
			var strBuilder strings.Builder
			strBuilder.WriteString("http://")
			strBuilder.WriteString(url)
			url = strBuilder.String()
			//way1
			// fmt.Sprintf("%s%s", "http://", url)
			//way2
			// url = strings.Join([]string{"http://", url}, "")
		}
		resp, er := http.Get(url)
		if er != nil {
			fmt.Fprintf(os.Stderr, "fetch1\t%v\n", er)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "fetch1\tStatus\t%v\n", resp.Status)
		_, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch1\t%v\n", er)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
