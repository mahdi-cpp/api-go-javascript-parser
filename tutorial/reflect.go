package tutorial

import (
	"fmt"
	"reflect"
)

type Person1 struct {
	Name string
	Age  int
}

func ReflectTest() {

	p := Person1{Name: "Alice", Age: 30}

	// Get the type of p
	t := reflect.TypeOf(p)
	fmt.Println("Type:", t)

	// Get the value of p
	v := reflect.ValueOf(p)
	fmt.Println("Value:", v)

	// Iterate over fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("Field: %s, Value: %v\n", field.Name, value.Interface())
	}
	//
	// Update a field value
	if v.FieldByName("Name").IsValid() {
		ageField := v.FieldByName("Name")
		ageField.SetString("Ali")
	}

	//fmt.Println("Updated Age:", p.Age)
}
