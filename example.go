package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/joshjon/go-phonebook/phonebook"
)

func main() {
	book := phonebook.New()
	loadDummyContacts(book)

	number := "0410000000"
	contact, _ := book.Get(number)
	fmt.Printf("- Got contact for '%s': %+v\n", number, contact)

	prefix := "0410000"
	contacts := book.FindByPrefix(prefix)
	fmt.Printf("- Found %d contacts for prefix '%s'\n", len(contacts), prefix)

	firstName := "firstname2000"
	contacts = book.FindByName(firstName, "")
	fmt.Printf("- Found %d contacts for first name '%s'\n", len(contacts), firstName)

	city := "city3000"
	contacts = book.FindByCity(city)
	fmt.Printf("- Found %d contacts for city '%s'\n", len(contacts), city)

	search := "lastname3000"
	contacts = book.Find(search)
	fmt.Printf("- Found %d contacts for generic search '%s'\n", len(contacts), search)

	search = "0410020"
	contacts = book.Find(search)
	fmt.Printf("- Found %d contacts for generic search '%s'\n", len(contacts), search)
}

func loadDummyContacts(pb *phonebook.PhoneBook) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 200000; i++ {
		contact := phonebook.Contact{}
		contact.Number = "04" + strconv.Itoa(10000000+i)

		num := i % 10000
		contact.FirstName = "firstname" + strconv.Itoa(num)
		contact.LastName = "lastname" + strconv.Itoa(num)

		if rand.Intn(2) == 1 {
			city := "city" + strconv.Itoa(num)
			contact.Address = fmt.Sprintf("1 foo st, %s, foo state, 1111, foo country", city)
		}

		if err := pb.Add(contact); err != nil {
			panic(err)
		}
	}
}
