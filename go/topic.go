package main

type topic struct {
	firstArticle int	`json:firstArticle`
	lastArticle  int	`json:lastArticle`
	header       string	`json:topicHeader`
}
