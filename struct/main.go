package main

import "fmt"

type people struct {
	firstName string
	lastName  string
	//contact   contactCard
	contactCard
}
type contactCard struct {
	phoneNumber string
	address     string
}

func (p *people) updateName(newName string) {
	(*p).firstName = newName
}

func main() {
	//var name people
	n := contactCard{
		phoneNumber: "ascc",
		address:     "jhjh",
	}
	fmt.Print(n)

	name := people{firstName: "Megha",
		lastName: "Agarwal",
		contactCard: contactCard{
			phoneNumber: "aac",
			address:     "lane",
		}}
	fmt.Println(name)
	fmt.Printf("%+v", name)
	// namePointer := &name
	// namePointer.updateName("Ray")
	name.updateName("Ray")
	fmt.Print(name)

}
