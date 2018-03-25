package main

import (
	"context"
	"fmt"
	"os"

	"log"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
) //library for networking

type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "9ee5d1475328cc8cd397b6c41910707e3d5b7bc5"},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := github.NewClient(tokenClient)

	repo, _, err := client.Repositories.Get(context, "openebs", "openebs")

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	pack := &Package{
		FullName:    *repo.FullName,
		Description: *repo.Description,
		ForksCount:  *repo.ForksCount,
		StarsCount:  *repo.StargazersCount,
	}

	fmt.Fprintf(w, "%s\n", pack.FullName)
	fmt.Fprintf(w, "Stargazers Count : %d\n", pack.StarsCount)
}

func main() {
	http.HandleFunc("/", index_handler)      //http handler //index_handler is a func
	err := http.ListenAndServe(":8000", nil) //localhost port 8000
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}

}
