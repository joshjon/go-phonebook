package phonebook

import (
	"fmt"
	"regexp"
	"strings"
)

// Contact represents a contact found in a phone book.
type Contact struct {
	Number    string
	FirstName string
	LastName  string
	Address   string
}

// Validate checks that each field value is valid.
func (c Contact) Validate() error {
	if matched, err := regexp.Match("^\\d{10}$", []byte(c.Number)); err != nil || !matched {
		return fmt.Errorf("phone number must contain 10 digits")
	} else if c.FirstName == "" {
		return fmt.Errorf("first name required")
	} else if c.LastName == "" {
		return fmt.Errorf("last name required")
	} else if c.Address != "" && len(strings.Split(c.Address, ",")) != 5 {
		return fmt.Errorf("address must be in the format '[street address], [city], [state/province], [zip code], [country]'")
	}

	return nil
}
