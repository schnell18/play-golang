package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/schnell18/play-golang/ch4/github"
)

const templ = `
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Github Issues</title>
</head>
<body>
<h1>{{len .Items}}/{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
  <th>Age(days)</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.HTMLURL}}'>{{.User.Login}}</td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</td>
  <td>{{.CreateAt | daysAgo}}</td>
</tr>
{{end}}
</table>
</body>
</html>
`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	// result, err := github.SearchIssues(os.Args[1:])
	result, err := github.SearchIssues([]string{"repo:golang/go 3133 10535"})
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
