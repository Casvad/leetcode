package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "log"
	"net/http"
)

type Data struct {
	City       string `json:"city"`
	Name       string `json:"name"`
	UserRating struct {
		AverageRating int32 `json:"average_rating"`
		Votes         int32 `json:"votes"`
	} `json:"user_rating"`
	Id int `json:"id"`
}
type PageResponse struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []Data `json:"data"`
}

func finestFoodOutlet(city string, votes int32) string {
	data := getData(city, 1, make([]Data, 0))

	bestRes := ""
	bestAvg := int32(0)

	for _, i := range data {
		if i.UserRating.Votes >= votes {
			if bestRes == "" || bestAvg < i.UserRating.AverageRating {
				bestRes = i.Name
				bestAvg = i.UserRating.AverageRating
			}
		}
	}

	return bestRes
}

func getData(city string, currentPage int, currentData []Data) []Data {
	res, err := http.Get(fmt.Sprintf("https://jsonmock.hackerrank.com/api/food_outlets?city=%s&page=%d", city, currentPage))
	if err != nil {
		panic(err)
	}
	var pageResponse PageResponse
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &pageResponse)

	if err != nil {
		panic(err)
	}

	if currentPage > pageResponse.TotalPages {
		return append(currentData, pageResponse.Data...)
	}

	return getData(city, currentPage+1, append(currentData, pageResponse.Data...))
}

func main() {

	finestFoodOutlet("Seattle", 1000)
}
