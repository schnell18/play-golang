package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"github.com/schnell18/play-golang/ch4/github"
)

const templ = `{{len .Items}}/{{.TotalCount}} issues:
{{range .Items}}--------------------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreateAt | daysAgo}} days
{{end}}`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	// result, err := github.SearchIssues(os.Args[1:])
	result, err := github.SearchIssues([]string{"repo:kubernetes/kubernetes is:open api"})
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
