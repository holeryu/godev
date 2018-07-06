package main

import "fmt"
import "godev/util"

func main() {
	countryCapitalMap := util.CountryCapitalMap();

	capital,ok := countryCapitalMap [ "China" ]

	if(ok) {
		fmt.Println(capital)
	}
}
