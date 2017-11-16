package main

type chapter struct {
	FirstArticle int64   `json:"first_article"`
	LastArticle  int64   `json:"last_article"`
	Header       string  `json:"chapter_header"`
	TopicArray   []topic `json:"topic_array"`
}
