package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v31/github"
	"golang.org/x/oauth2"
	"gopkg.in/square/go-jose.v2/json"
)

type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	Arg string `json:"arg"`
	Subtitle string `json:"subtitle"`
	Title string `json:"title"`
}

func main() {
	ctx := context.Background()
	//TODO: move github token into a cli flag
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ""})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opts := &github.SearchOptions{Sort: "created", Order: "desc"}
	//TODO: read input query from cli flag
	res, _, err := client.Search.Code(ctx, "docker repo:wayanjimmy/notebook", opts)
	if err != nil {
		log.Fatal(err)
	}

	items := []Item{}

	for _, cr := range res.CodeResults {
		item := Item{Arg: cr.GetHTMLURL(), Subtitle: cr.GetPath(), Title: cr.GetName()}
		items = append(items, item)
	}

	sr := Items{
		Items: items,
	}

	m, err := json.Marshal(sr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(m))
}