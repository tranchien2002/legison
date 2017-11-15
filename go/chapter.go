package main

type chapter struct {
	firstArticle int     `json:firstArticle`
	lastArticle  int     `json:lastArticle`
	header       string  `json:chapterHeader`
	topicArray   []topic `json:topicArray`
}
