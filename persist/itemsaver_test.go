package persist

import (
	"context"
	"encoding/json"
	"go-crawler/model"
	"testing"

	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Name:       "我的后半生",
		Gender:     "女",
		Age:        34,
		Height:     174,
		Weight:     56,
		Income:     "4000-8000",
		Marriage:   "离异",
		Education:  "大学本科",
		Occupation: "人事/行政",
		Hokou:      "湖南长沙",
		Xinzuo:     "牧羊座",
		House:      "已购房",
		Car:        "未购车",
	}

	// View by http://127.0.0.1:9200/dating_profile/zhenai/_search
	id, err := save(expected)
	if err != nil {
		panic(err)
	}

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%+v", resp.Source)

	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
