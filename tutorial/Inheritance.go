package tutorial

import "fmt"

type Animal struct {
	Name string
}

func (animal Animal) showName() string {
	return animal.Name
}
func (animal Animal) setName(name string) string {
	animal.Name = name
	return animal.Name
}

type Cat struct {
	// anonymous field,
	// this is composition where

	// the struct is embedded
	Animal
}

type Dog struct {
	Age int
	Animal
}

// object type detect with interface
func shapeType(i interface{}) {

	switch v := i.(type) {
	case Animal:
		fmt.Println("object is Animal")
	case Cat:
		fmt.Println("object is Cat")
	case Dog:
		fmt.Println("object is Dog")
	default:
		fmt.Println("I don't know about type", v)
	}
}

func Start() {
	a := Animal{Name: "animal"}
	cat1 := Cat{
		Animal{Name: "Cat 1"},
	}
	cat2 := Cat{
		Animal{Name: "Cat 2"},
	}

	obj := Dog{
		Age:    14,
		Animal: Animal{Name: "dog"},
	}

	cat2.setName("Good Cat")
	cat2.Name = "Ali 5"

	shapeType(a)
	shapeType(cat1)
	shapeType(cat2)
	shapeType(obj)

	fmt.Println(cat1.showName())
	fmt.Println(cat2.showName())

	fmt.Println(obj.Age)
}
