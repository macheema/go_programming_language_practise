package section45

//IssuesSearchResult ...
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
