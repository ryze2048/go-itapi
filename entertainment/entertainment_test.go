package entertainment

import (
	"fmt"
	"github.com/ryze2048/go-itapi/model"
	"testing"
)

var entertainment = new(Entertainment)

func TestGetRandomQuote(t *testing.T) {
	request := &model.RandomQuoteRequest{Key: ``}
	var response model.RandomQuoteResponse
	if err := entertainment.Any(request, &response, model.RandomQuoteUrl); err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(response)
}

func TestGetJoke(t *testing.T) {
	request := &model.JokeRequest{
		Key:  "",
		Num:  "50",
		Page: "1",
	}
	var response model.JokeResponse
	if err := entertainment.Any(request, &response, model.JokeUrl); err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(response)
}

func TestEditEmoji(t *testing.T) {
}
