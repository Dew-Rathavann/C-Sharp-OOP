package main

import "fmt"

func main() {
	countryCapitalMap := map[string]string{"France": "Paris", "Thailand": "Bangkok", "Japan": "Tokyo", "America": "Washington DC"}
	fmt.Println("Original map")
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")
	fmt.Println("Updated map")
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}
