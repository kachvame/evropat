package main

type City struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Country string `json:"country,omitempty"`
}

var cities = []City{
	{1, "Русе", "България"},
	{2, "Трявна", "България"},
	{3, "Пловдив", "България"},
	{4, "Плевен", "България"},
	{5, "София", "България"},
	{6, "Сливен", "България"},
}

type citiesResponse struct {
	Cities []City `json:"cities"`
}

type Office struct {
	ID     int    `json:"id,omitempty"`
	CityID int    `json:"city_id,omitempty"`
	Name   string `json:"name,omitempty"`
}

var offices = []Office{
	{1, 1, "Майкати"},
	{2, 2, "DJ Шанаджия"},
	{3, 3, "Zeron"},
	{4, 4, "Курви"},
	{5, 5, "Ало"},
	{6, 6, "Пико"},
}

type officesResponse struct {
	Offices []Office `json:"offices"`
}
