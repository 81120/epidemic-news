package model

import "fmt"

// HttpRes hold the data for http response
type HttpRes struct {
	Ret  int    `json:"ret"`
	Data string `json:"data"`
}

// CityData hold the data for city item
type CityData struct {
	Country string `json:"country"`
	Area    string `json:"area"`
	City    string `json:"city"`
	Confirm int    `json:"confirm"`
	Suspect int    `json:"suspect"`
	Dead    int    `json:"dead"`
	Heal    int    `json:"heal"`
}

// CityList []CityData
type CityList []CityData

// DisplayCityData as name
func (city *CityData) DisplayCityData() {
	fmt.Println(city.City, "确诊:", city.Confirm, "死亡:", city.Dead, "治愈:", city.Heal)
}

// DisplayCityList as name
func (cityList *CityList) DisplayCityList() {
	for _, v := range *cityList {
		v.DisplayCityData()
	}
}
