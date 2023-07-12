package news

import (
	"fmt"
	"github.com/ryze2048/go-itapi/model"
	"testing"
)

var newsAPI = NewsAPI{}

func TestNews(t *testing.T) {
	data := model.GeneralNewsRequest{
		Common: model.Common{Key: ``},
	}
	s, err := newsAPI.GeneralNews(&data)
	if err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(s)
}

func TestNewsAPI_Hours24(t *testing.T) {
	data := model.Common{Key: ``}
	newsAPI.Hours24(&data)
}

func TestNewsAPI_HotNews360(t *testing.T) {
	data := model.Common{Key: ``}
	resp, err := newsAPI.Kwai(&data)
	if err != nil {
		t.Error("err -->", err)
	}
	fmt.Println(resp)
}
