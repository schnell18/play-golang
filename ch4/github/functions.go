package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	defer func() {
		resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
