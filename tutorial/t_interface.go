package tutorial

import "fmt"

// Car interface ---------------------------------------
type Car interface {
	Name() string
	Year() int
}

// BMW ---------------------------------------
type BMW struct {
	name string
	year int
}

func (m BMW) Name() string {
	return m.name
}
func (m BMW) Year() int {
	return m.year
}

//---------------------------------------

func ShowInterface() {

	// Accessing elements of
	// the tank interface

	var car Car
	car = BMW{name: "BMW", year: 2007}

	fmt.Println("Name of car :", car.Name())
	fmt.Println("Year of car:", car.Year())
}
