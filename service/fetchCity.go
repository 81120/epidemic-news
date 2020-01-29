package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/81120/epidemic-news/model"
)

// FetchAllData get the city data
func FetchAllData() model.CityList {
	baseURL := "https://view.inews.qq.com/g2/getOnsInfo?name=wuwei_ww_area_counts"
	res, _ := http.Get(baseURL)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var httpRes model.HttpRes
	json.Unmarshal(body, &httpRes)
	var cities model.CityList
	json.Unmarshal([]byte(httpRes.Data), &cities)
	return cities
}

// FetchDataByArea get the data by area
func FetchDataByArea(area string) model.CityList {
	citiesData := FetchAllData()
	var res model.CityList
	for _, v := range citiesData {
		if v.Area == area {
			res = append(res, v)
		}
	}
	return res
}

// FetchDataByCity get the data by city
func FetchDataByCity(city string) model.CityData {
	citiesData := FetchAllData()
	var res model.CityData
	for _, v := range citiesData {
		if v.City == city {
			res = v
			break
		}
	}
	return res
}
