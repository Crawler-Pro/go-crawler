package parser

import (
	"go-crawler-distributed/config"
	"go-crawler/engine"
	"go-crawler/model"
	"regexp"
	"strconv"
)

var genderReg = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var argReg = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`)
var heightReg = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)cm</div>`)
var weightReg = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)kg</div>`)
var incomeReg = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)
var marriageReg = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var educationReg = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var occupationReg = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var hokouReg = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>籍贯:([^<]+)</div>`)
var xinzuoReg = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var houseReg = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div>`)
var carReg = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div>`)

var commonReg = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)

var idUrlReg = regexp.MustCompile(`https://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	arg, err := strconv.Atoi(extractSubString(contents, argReg))
	if err == nil {
		profile.Age = arg
	}

	height, err := strconv.Atoi(extractSubString(contents, heightReg))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractSubString(contents, weightReg))
	if err == nil {
		profile.Weight = weight
	}

	profile.Gender = extractAllString(contents, commonReg, 0)
	profile.Income = extractSubString(contents, incomeReg)
	profile.Marriage = extractAllString(contents, commonReg, 0)
	profile.Education = extractAllString(contents, commonReg, 8)
	profile.Occupation = extractAllString(contents, commonReg, 7)
	profile.Hokou = extractSubString(contents, hokouReg)
	profile.Xinzuo = extractAllString(contents, commonReg, 2)
	profile.House = extractSubString(contents, houseReg)
	profile.Car = extractSubString(contents, carReg)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractSubString([]byte(url), idUrlReg),
				Payload: profile,
			},
		},
	}

	return result
}

func extractSubString(contents []byte, re *regexp.Regexp) string {
	matchs := re.FindSubmatch(contents)

	if len(matchs) >= 2 {
		return string(matchs[1])
	}
	return ""
}

func extractAllString(contents []byte, re *regexp.Regexp, index int) string {
	matchs := re.FindAllStringSubmatch(string(contents), -1)

	if len(matchs) >= index+1 {
		return matchs[index][1]
	}
	return ""
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseProfile(contents, url, p.userName)

}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	//return "ProfileParser", p.userName
	return config.ParseProfile, p.userName //
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
