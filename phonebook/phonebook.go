package phonebook

import (
	"fmt"
	"regexp"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/joshjon/go-phonebook/internal/index"
	"github.com/joshjon/go-phonebook/internal/trie"
)

const (
	indexFirstName int = iota
	indexLastName
	indexFullName
	indexCity
)

type PhoneBook struct {
	contacts *trie.NumberTrie[Contact]
	indexes  *index.Indexes[Contact]
}

func New() *PhoneBook {
	return &PhoneBook{
		contacts: trie.NewNumberTrie[Contact](),
		indexes: index.NewIndexes[Contact](
			index.NewMapIndex(indexFirstName, func(contact Contact) (string, bool) { return contact.FirstName, true }),
			index.NewMapIndex(indexLastName, func(contact Contact) (string, bool) { return contact.LastName, true }),
			index.NewMapIndex(indexFullName, func(contact Contact) (string, bool) { return contact.FirstName + contact.LastName, true }),
			index.NewMapIndex(indexCity, func(contact Contact) (string, bool) { return cityFromAddress(contact.Address) }),
		),
	}
}

func (p *PhoneBook) Add(contact Contact) error {
	if err := contact.Validate(); err != nil {
		return err
	}

	if err := p.contacts.Insert(contact.Number, contact); err != nil {
		return err
	}

	p.indexes.Add(contact)
	return nil
}

func (p *PhoneBook) Update(number string, update Contact) error {
	if number != update.Number {
		if _, ok := p.Get(update.Number); ok {
			return fmt.Errorf("contact already exists for new number %s", update.Number)
		}
	}
	p.Delete(number)
	return p.Add(update)
}

func (p *PhoneBook) Get(number string) (Contact, bool) {
	return p.contacts.Get(number)
}

func (p *PhoneBook) FindByPrefix(numberPrefix string) []Contact {
	if contacts, ok := p.contacts.FindByPrefix(numberPrefix); ok {
		return contacts
	}
	return []Contact{}
}

func (p *PhoneBook) FindByName(firstName string, lastName string) []Contact {
	var contacts []Contact
	var ok bool
	if firstName != "" && lastName != "" {
		contacts, ok = p.indexes.Get(indexFullName, firstName+lastName)
	} else if firstName != "" {
		contacts, ok = p.indexes.Get(indexFirstName, firstName)
	} else if lastName != "" {
		contacts, ok = p.indexes.Get(indexLastName, lastName)
	}
	if ok {
		return contacts
	}
	return []Contact{}
}

func (p *PhoneBook) FindByCity(city string) []Contact {
	if contacts, ok := p.indexes.Get(indexCity, city); ok {
		return contacts
	}
	return []Contact{}
}

func (p *PhoneBook) Find(search string) []Contact {
	union := mapset.NewSet[Contact]()
	if matched, err := regexp.Match("^\\d{1,10}$", []byte(search)); err == nil && matched {
		union = union.Union(mapset.NewSet(p.FindByPrefix(search)...))
	}
	union = union.Union(mapset.NewSet(p.FindByName(search, "")...))
	union = union.Union(mapset.NewSet(p.FindByName("", search)...))
	union = union.Union(mapset.NewSet(p.FindByCity(search)...))
	return union.ToSlice()
}

func (p *PhoneBook) Delete(number string) {
	if contact, ok := p.Get(number); ok {
		p.contacts.Delete(number)
		p.indexes.Delete(contact)
	}
}

func cityFromAddress(address string) (string, bool) {
	if address != "" {
		return strings.Split(address, ", ")[1], true
	}
	return "", false
}
