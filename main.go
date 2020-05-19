package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/google/go-github/v31/github"
	"golang.org/x/oauth2"
	"gopkg.in/square/go-jose.v2/json"
)

type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	Arg      string `json:"arg"`
	Subtitle string `json:"subtitle"`
	Title    string `json:"title"`
}

type Flags struct {
	Query       string
	AccessToken string
}

func main() {
	f := Flags{}
	flag.StringVar(&f.Query, "query", "something", "something you want to search...")
	flag.StringVar(&f.AccessToken, "token", "", "github access token")
	flag.Parse()

	exitCode := 0
	if err := run(f); err != nil {
		fmt.Printf("%v", err)
		exitCode = -1
	}
	os.Exit(exitCode)
}

func run(f Flags) error {
	ctx := context.Background()
	if len(f.AccessToken) == 0 {
		return errors.New("access token can't be blank")
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: f.AccessToken})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opts := &github.SearchOptions{Sort: "created", Order: "desc"}

	res, _, err := client.Search.Code(ctx, fmt.Sprintf("%s repo:wayanjimmy/notebook", f.Query), opts)
	if err != nil {
		return err
	}

	items := []Item{}

	for _, cr := range res.CodeResults {
		item := Item{Arg: cr.GetHTMLURL(), Subtitle: cr.GetPath(), Title: cr.GetName()}
		items = append(items, item)
	}

	res2, _, err := client.Search.Code(ctx, fmt.Sprintf("%s repo:wayanjimmy/zettelkasten", f.Query), opts)
	if err != nil {
		return err
	}

	for _, cr := range res2.CodeResults {
		item := Item{Arg: cr.GetHTMLURL(), Subtitle: cr.GetPath(), Title: cr.GetName()}
		items = append(items, item)
	}

	sr := Items{
		Items: items,
	}

	m, err := json.Marshal(sr)
	if err != nil {
		return err
	}

	fmt.Println(string(m))

	return nil
}
