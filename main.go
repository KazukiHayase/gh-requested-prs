package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/cli/go-gh/v2"
	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"
)

func main() {
	client, err := api.DefaultGraphQLClient()
	if err != nil {
		log.Fatal(err)
	}

	var query struct {
		Search struct {
			Edges []struct {
				Node struct {
					PullRequest struct {
						URL string
					} `graphql:"... on PullRequest"`
				}
			}
		} `graphql:"search(first: 100, query: $q, type: ISSUE)"`
	}

	filters := []string{
		"is:pr",
		"is:open",
		"review-requested:@me",
		"sort:created-asc",
	}

	var orgsArg string
	flag.StringVar(&orgsArg, "orgs", "", "comma separated list of orgs to include")
	flag.Parse()

	var orgs []string
	if orgsArg != "" {
		orgs = strings.Split(orgsArg, ",")
	} else {
		orgList, _, err := gh.Exec("org", "list")
		if err != nil {
			log.Fatal(err)
		}

		orgs = strings.Split(orgList.String(), "\n")
	}

	for _, org := range orgs {
		if org == "" {
			continue
		}

		filters = append(filters, fmt.Sprintf("org:%s", org))
	}

	variables := map[string]any{
		"q": graphql.String(strings.Join(filters, " ")),
	}

	if err := client.Query("ReviewRequestedPullRequests", &query, variables); err != nil {
		log.Fatal(err)
	}

	var urls []string
	for _, edge := range query.Search.Edges {
		urls = append(urls, edge.Node.PullRequest.URL)
	}

	if len(urls) == 0 {
		fmt.Println("no review requested pull requests")
		return
	}

	fmt.Println("open pull requests...")
	for _, url := range urls {
		fmt.Println(url)
		err := exec.Command("open", url).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
