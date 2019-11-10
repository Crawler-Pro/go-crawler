package parser

import (
	"go-crawler/engine"
	"log"
	"regexp"
)

var (
	profileReg = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
)

// ParseCity 解析城市
func ParseCity(contents []byte) engine.ParseResult {
	matches := profileReg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		log.Printf("Analysis of the user: %v", name)
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}

	return result
}
