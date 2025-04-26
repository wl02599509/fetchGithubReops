package fetch

import (
	"fmt"
	"encoding/json"
	"github.com/go-zoox/fetch"
)

type Repository struct {
	Stars	int `json:"stargazers_count"`
}

func GetTotalCount(account string, column string) int {
	repos := getUserRepos(account)

	var totalCount int

	for _, repo := range repos {
		totalCount += int(getIntDataBy(repo, column))
	}

	return totalCount
}

func getUserRepos(account string) []Repository {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", account)
	response, err := fetch.Get(url)

	if err != nil {
		panic(err)
	}

	var repos []Repository

	body := response.Body

	json.Unmarshal(body, &repos)

	return repos
}

func getIntDataBy(repo Repository, column string) int {
	switch column {
		case "stargazers_count":
			return repo.Stars
		default:
			return 0
	}
}