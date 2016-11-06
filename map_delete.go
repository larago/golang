package main

import "fmt"

func main() {
	countryCapitalMap := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
	}

	fmt.Println("Map")

	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	delete(countryCapitalMap, "France")
	fmt.Println("Entry of France was deleted")

	fmt.Println("Has been removed")

	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}
