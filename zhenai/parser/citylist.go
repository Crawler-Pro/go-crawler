package parser

import (
	"go-crawler-distributed/config"
	"go-crawler/engine"
	"log"
	"regexp"
)

const cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// ParseCityList is parse city
func ParseCityList(contents []byte, _ string) engine.ParseResult {
	reg := regexp.MustCompile(cityListReg)
	matches := reg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		log.Printf("Analysis of the city: %v", string(m[2]))
		//result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//ParserFunc: ParseCity,
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	return result
}
