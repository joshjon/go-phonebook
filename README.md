# 📒 Go Phone Book

Contains Go package `phonebook` that implements a data structure to represent a phone book with the following
properties:

- [x] It should store contact records each containing a first name (required), a last name (required), an address (
  optional), and a phone number (required).
- [x] Phone numbers are expected to be unique per contact and 10 digits long (assume that phone numbers do not contain
  country codes or non-numeric characters for simplicity).
- [x] Multiple contacts may have the same addresses, first names, and last names.
- [x] A method that takes a phone number and looks up and returns the contact associated with the phone number
- [x] A method the consumes a first name, last name, phone number, and address (optional) and creates a new contact in
  the phone book. The method should not allow conflicts (two contacts with the same phone number).
- [x] A method for updating a contact (any fields associated with a contact should be able to be updated)
- [x] A method for looking up a list of contacts by first and/or last name. The method should take two variables, one
  for first name and one for last name, and at least one of them must not be empty.
- [x] A method for looking up a list of contacts by city (assume all addresses are in the form "[street address], [city]
  , [state/province], [zip code], [country]"
- [x] A method for deleting a contact by phone number. No data pertaining to the contact should remain anywhere in the
  data structure once the contact is deleted.
- [x] A method for looking up a list of contacts by an arbitrary string (could be a city, name, country, zip code, etc.)
  . It is expected that the string is a complete value (i.e. not half of a country name or the first three digits of a
  phone number). The method should return all contacts with metadata containing that string. There is flexibility here
  in how this function is implemented. Feel free to decide which fields should be searchable.
- [x] A method for looking up a list of contacts by the phone number prefix. The method should consume a string
  containing n number of digits, where 0 < n <= 10 and return an of all contacts whose phone numbers start with the
  given prefix.

## Usage

Refer to `example.go` for example usage of the phone book.

## Implementation Details

The phone book utilises map indexes for each searchable field to dramatically reduce search times to O(1). A trie is
also used to store phone numbers which allows retrieval and prefix search to be performed at O(n), where 0 < n <= 10.
