package main

import "regexp"

type tobject struct {
	sign          string    `json:sign`
	id            string    `json:id`
	enforcer      string    `json:enforcer`
	baseon        string    `json:baseon`
	effectiveDate string    `json:effectiveDate`
	passDate      string    `json:passDate`
	chapterArray  []chapter `json:chapterArray`
	articleArray  []article `json:articleArray`
}

func creatobject(legis string) *tobject {
	result := tobject{}

	re := regexp.MustCompile(SIGN_REGX)
	match := re.FindAllString(legis, -1)
	result.sign = match[0]

	
	
	return &result
}
