package persist

import (
	"context"
	"encoding/json"
	"go-crawler/engine"
	"go-crawler/model"
	"testing"

	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "https://album.zhenai.com/u/110107668",
		Type: "zhenai",
		Id:   "110107668",
		Payload: model.Profile{
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
		},
	}

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"

	// View by http://127.0.0.1:9200/dating_profile/zhenai/_search
	// Save expected item
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%+v", resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
