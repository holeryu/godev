package main

import "fmt"
import "github.com/holeryu/godev/util"

func main() {
	countryCapitalMap := util.CountryCapitalMap();

	capital,ok := countryCapitalMap [ "China" ]

	if(ok) {
		fmt.Println(capital)
	}
}
