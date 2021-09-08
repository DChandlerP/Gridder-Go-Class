package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type Person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	p := Person{
		firstName: "Gopher",
		lastName:  "Con",
		contactInfo:     contactInfo{
			email:   "email@email.com",
			zipCode: 12345,
	},
}
	p.updateName("GopherCon")
	p.print()
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

func (p *Person) updateName(name string) {
	p.firstName = name
}
