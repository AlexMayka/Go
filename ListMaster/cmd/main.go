package main

import "ListMaster/internal"

func main() {
	persons := internal.ListPerson{}
	persons.AddPerson("Alex", "Mayka", 33)
	persons.AddPerson("Alex", "Ivanov", 33)
	persons.AddPerson("Ivan", "Ivanov", 42)
	persons.AddPerson("John", "Doe", 33)
	persons.AddPerson("Sasha", "Ivanova", 23)
	persons.AddPerson("Alex", "Mayka", 23)
	persons.AddPerson("Ivan", "Ivanov", 42)
	persons.AddPerson("John", "Doe", 33)
	persons.AddPerson("Sasha", "Ivanova", 23)

	persons.SortBy(func(p *internal.Person) any {
		return p.FirstName + " " + p.LastName
	},
		false,
	)

	persons.PrintList()
}
