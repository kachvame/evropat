package main

type City struct {
	ID      int
	Name    string
	Country string
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
	Cities []City
}
