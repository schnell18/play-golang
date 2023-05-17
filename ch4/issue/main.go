package main

import (
	"fmt"
	"log"
	"os"

	"github.com/schnell18/play-golang/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%06d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
