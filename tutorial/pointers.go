package tutorial

import (
	"fmt"
	"log"
)

var p *float64

var a float64 = 45

var Names map[int]string

var Countries = []string{"Austria", "Belgium", "Bulgaria"}

func PointersTest() {
	//p = &a
	//fmt.Println(p)

	addCountries(&Countries)
	log.Println(Countries)

	//Names := make(map[int]string)
	//maps(Names)
	//Names[7] = "Maryam"
	//Names[8] = "Sara"
	//
	//fmt.Println(Names)

	//var myPointerVar *int
	//fmt.Println(myPointerVar)
}

func addCountries(countries *[]string) {
	*countries = append(*countries, []string{"Croatia", "Republic of Cyprus", "Czech Republic", "Denmark", "Estonia", "Finland", "France", "Germany", "Greece", "Hungary", "Ireland", "Italy", "Latvia", "Lithuania", "Luxembourg", "Malta", "Netherlands", "Poland", "Portugal", "Romania", "Slovakia", "Slovenia", "Spain", "Sweden"}...)
	log.Println(countries)
}

func maps(values map[int]string) {
	values[5] = "Ali"
	values[125] = "Reza"
	values[45666] = "Mahdi"
	fmt.Println(values)
}
