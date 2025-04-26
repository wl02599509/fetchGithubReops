package main

import (
	"fmt"
	"flag"
	"fetchGithubRepos/fetch"
	"fetchGithubRepos/values"
)

func main() {
	account := flag.String("account", "", "GitHub 帳號")
	column := flag.String("column", "stargazers_count", "要計算的欄位")
	symbol := flag.String("symbol", "*", "要顯示的符號")
	breakline := flag.Int("breakline", 0, "要折行的單位數")
	
	flag.Parse()

	count := fetch.GetTotalCount(*account, *column)
	result := values.TransCount(count, *symbol, *breakline)

	fmt.Print(result)
}