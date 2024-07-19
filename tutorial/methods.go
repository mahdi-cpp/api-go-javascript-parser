package tutorial

import "fmt"

type Person struct {
	name string
	Age  int
}

// Method with a receiver of type Person
func (p Person) greet(name string) {
	p.name = name
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.Age)
}

func (p *Person) celebrateBirthday(name string) {
	p.name = name
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.Age)
}

type Book struct {
	Name    string
	timeAdd string
	count   int32
}

func bookAdd(pBook *Book) {
	pBook.Name = "Alice 2"
	pBook.count = 45
	fmt.Println(pBook)
}

func Methods() {
	// Create an instance of the Person struct
	person := Person{name: "Alice5", Age: 30}

	person2 := &Person{name: "Reza", Age: 24}

	pBook := &Book{Name: "Alice in", count: 5, timeAdd: "2014"}

	fmt.Println(pBook.Name)
	pBook.count = 7
	fmt.Println(pBook.count)

	bookAdd(pBook)
	fmt.Println(pBook.Name)

	// Call the greet method on the person instance
	person.greet("Mahdi")
	person2.celebrateBirthday("Mahdi")

	fmt.Println(person2.name)

}
