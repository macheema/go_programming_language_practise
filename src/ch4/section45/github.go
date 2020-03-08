package section45

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//IssuesURL ...
const IssuesURL = "https://api.github.com/search/issues"

//DoQuery ...
func DoQuery(searchItems []string) {
	result, err := SearchIssues(searchItems)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%10d%20s %1s\n",
			item.Number, item.User.Login, item.Title)
	}
}

//SearchIssues queries the GitHub issue tracker.
func SearchIssues(searchItem []string) (*IssuesSearchResult, error) {
	query := fmt.Sprintf("%s?q=%s", IssuesURL, url.QueryEscape(strings.Join(searchItem, " ")))
	resp, err := http.Get(query)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	var result IssuesSearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
