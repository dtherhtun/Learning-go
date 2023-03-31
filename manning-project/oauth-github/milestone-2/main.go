package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v50/github"

	"milestone-2/testutils"
)

func main() {
	ctx := context.Background()
	httpClient := testutils.HttpClientWithGithubStub("")
	client := github.NewClient(httpClient)
	u, _, err := client.Users.Get(ctx, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GitHub login: %s\n", *u.Login)
}
