package parser

import (
	"go-crawler-distributed/config"
	"go-crawler/engine"
	"log"
	"regexp"
)

var (
	profileReg = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityURLReg = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileReg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		log.Printf("Analysis of the user: %v", name)
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(name),
		})
	}

	matches = cityURLReg.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//ParserFunc:ParseCity,
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	return result
}
