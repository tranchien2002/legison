package main

import "regexp"

type article struct{
	header string	`json:header`
	content string	`json:content`
}