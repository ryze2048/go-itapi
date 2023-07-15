package news

import (
	"fmt"
	"github.com/ryze2048/go-itapi/model"
	"testing"
)

func TestNewsAPI_Unify(t *testing.T) {
	var client = new(NewsAPI)
	data := &model.UnifyNews{
		Common:     model.Common{Key: ``},
		RequestUrl: model.ZhihuUrl,
	}
	info, err := client.Unify(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(info))
	var res model.ZhihuResponse
	err = client.Unmarshal(info, &res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}

func TestNewsAPI_All(t *testing.T) {
	var client = new(NewsAPI)

	allRequest := model.AllRequest{
		Common: model.Common{Key: ``},
		Type:   "baidu",
	}

	res, err := client.All(&allRequest)
	if err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(res)
}

func TestNewsAPI_GeneralNews(t *testing.T) {
	var client = new(NewsAPI)
	data := &model.GeneralNewsRequest{
		Common: model.Common{Key: ``},
		Type:   "baidu",
		Num:    "1",
		Page:   "20",
	}

	resp, err := client.GeneralNews(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
