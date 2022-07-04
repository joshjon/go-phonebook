package phonebook

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPhoneBook_Add(t *testing.T) {
	phoneBook := New()
	contact := Contact{
		Number:    "0123456789",
		FirstName: "Foo",
		LastName:  "Bar",
	}
	require.NoError(t, phoneBook.Add(contact))
}

func TestPhoneBook_Add_duplicateError(t *testing.T) {
	phoneBook := New()
	contact := Contact{
		Number:    "0123456789",
		FirstName: "Foo",
		LastName:  "Bar",
	}
	require.NoError(t, phoneBook.Add(contact))
	require.Error(t, phoneBook.Add(contact))
}

func TestPhoneBook_Get(t *testing.T) {
	phoneBook := New()
	want := Contact{
		Number:    "0123456789",
		FirstName: "Foo",
		LastName:  "Bar",
	}
	require.NoError(t, phoneBook.Add(want))

	tests := []struct {
		name   string
		number string
		found  bool
	}{
		{
			name:   "success",
			number: want.Number,
			found:  true,
		},
		{
			name:   "not found",
			number: "0000000000",
			found:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := phoneBook.Get(tt.number)
			require.Equal(t, tt.found, ok)
			if ok {
				require.Equal(t, want, got)
			} else {
				require.Empty(t, got)
			}
		})
	}
}

func TestPhoneBook_FindByPrefix(t *testing.T) {
	phoneBook := New()
	prefix := "11"
	want1 := Contact{Number: prefix + "23456789", FirstName: "One", LastName: "One"}
	want2 := Contact{Number: prefix + "76543210", FirstName: "Two", LastName: "Two"}
	dummy := Contact{Number: "1232167890", FirstName: "Three", LastName: "Three"}
	require.NoError(t, phoneBook.Add(want1))
	require.NoError(t, phoneBook.Add(want2))
	require.NoError(t, phoneBook.Add(dummy))

	tests := []struct {
		name   string
		prefix string
		found  bool
	}{
		{
			name:   "success",
			prefix: prefix,
			found:  true,
		},
		{
			name:   "not found",
			prefix: "48",
			found:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := phoneBook.FindByPrefix(tt.prefix)
			if tt.found {
				require.ElementsMatch(t, got, []Contact{want1, want2})
			} else {
				require.Empty(t, got)
			}
		})
	}
}

func TestPhoneBook_FindByName(t *testing.T) {
	phoneBook := New()

	firstName, lastName := "One", "Two"
	want1 := Contact{Number: "0123456789", FirstName: firstName, LastName: lastName}
	want2 := Contact{Number: "9876543210", FirstName: firstName, LastName: lastName}
	dummy := Contact{Number: "5432167890", FirstName: "Three", LastName: "Three"}

	require.NoError(t, phoneBook.Add(want1))
	require.NoError(t, phoneBook.Add(want2))
	require.NoError(t, phoneBook.Add(dummy))

	tests := []struct {
		name      string
		firstName string
		lastName  string
		found     bool
	}{
		{
			name:      "first name success",
			firstName: want1.FirstName,
			found:     true,
		},
		{
			name:      "first name not found",
			firstName: "Random",
			found:     false,
		},
		{
			name:     "last name success",
			lastName: want1.LastName,
			found:    true,
		},
		{
			name:     "last name not found",
			lastName: "Random",
			found:    false,
		},
		{
			name:      "full name success",
			firstName: want1.FirstName,
			lastName:  want1.LastName,
			found:     true,
		},
		{
			name:      "full name not found",
			firstName: "Random",
			lastName:  "Random",
			found:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := phoneBook.FindByName(tt.firstName, tt.lastName)
			if tt.found {
				require.ElementsMatch(t, got, []Contact{want1, want2})
			} else {
				require.Empty(t, got)
			}
		})
	}
}

func TestPhoneBook_FindByCity(t *testing.T) {
	phoneBook := New()
	city := "Foo City"
	address := fmt.Sprintf("1 Foo St, %s, Foo State, 1111, Foo Country", city)
	want1 := Contact{Number: "0123456789", FirstName: "One", LastName: "One", Address: address}
	want2 := Contact{Number: "9876543210", FirstName: "Two", LastName: "Two", Address: address}
	dummy := Contact{Number: "5432167890", FirstName: "Three", LastName: "Three", Address: "1 Dummy St, Dummy City, Dummy State, 2222, Dummy Country"}
	require.NoError(t, phoneBook.Add(want1))
	require.NoError(t, phoneBook.Add(want2))
	require.NoError(t, phoneBook.Add(dummy))

	tests := []struct {
		name  string
		city  string
		found bool
	}{
		{
			name:  "success",
			city:  city,
			found: true,
		},
		{
			name:  "not found",
			city:  "random",
			found: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := phoneBook.FindByCity(tt.city)
			if tt.found {
				require.ElementsMatch(t, got, []Contact{want1, want2})
			} else {
				require.Empty(t, got)
			}
		})
	}
}

func TestPhoneBook_Find(t *testing.T) {
	phoneBook := New()
	prefix, firstName, lastName, city := "0011", "Foo", "Bar", "Foo City"
	common := "Common"
	want1 := Contact{Number: prefix + "223344", FirstName: firstName, LastName: lastName, Address: newAddress(city)}
	want2 := Contact{Number: prefix + "225566", FirstName: firstName, LastName: lastName, Address: newAddress(city)}
	want3 := Contact{Number: "1111111111", FirstName: common, LastName: "Three", Address: newAddress("Dummy")}
	want4 := Contact{Number: "9999999999", FirstName: "Four", LastName: "Four", Address: newAddress(common)}
	require.NoError(t, phoneBook.Add(want1))
	require.NoError(t, phoneBook.Add(want2))
	require.NoError(t, phoneBook.Add(want3))
	require.NoError(t, phoneBook.Add(want4))

	tests := []struct {
		name   string
		search string
		want   []Contact
	}{
		{
			name:   "find using number",
			search: prefix + "223344",
			want:   []Contact{want1},
		},
		{
			name:   "find using number prefix",
			search: prefix,
			want:   []Contact{want1, want2},
		},
		{
			name:   "find using first name",
			search: firstName,
			want:   []Contact{want1, want2},
		},
		{
			name:   "find using last name",
			search: lastName,
			want:   []Contact{want1, want2},
		},
		{
			name:   "find using city",
			search: city,
			want:   []Contact{want1, want2},
		},
		{
			name:   "find using search term that matches multiple fields",
			search: common,
			want:   []Contact{want3, want4},
		},
		{
			name:   "not found",
			search: "random",
			want:   []Contact{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := phoneBook.Find(tt.search)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestPhoneBook_Delete(t *testing.T) {
	phoneBook := New()
	prefix, city := "01", "Foo City"
	want := Contact{
		Number:    prefix + "23456789",
		FirstName: "Foo",
		LastName:  "Bar",
		Address:   newAddress(city),
	}
	require.NoError(t, phoneBook.Add(want))
	phoneBook.Delete(want.Number)

	tests := []struct {
		name string
		fn   func() any
		want any
	}{
		{
			name: "get fail",
			fn: func() any {
				contact, _ := phoneBook.Get(want.Number)
				return contact
			},
		},
		{
			name: "find by prefix fail",
			fn:   func() any { return phoneBook.FindByPrefix(prefix) },
		},
		{
			name: "find by name (first) fail",
			fn:   func() any { return phoneBook.FindByName(want.FirstName, "") },
		},
		{
			name: "find by name (last) fail",
			fn:   func() any { return phoneBook.FindByName("", want.LastName) },
		},
		{
			name: "find by name (first & last) fail",
			fn:   func() any { return phoneBook.FindByName(want.LastName, want.LastName) },
		},
		{
			name: "find by name city fail",
			fn:   func() any { return phoneBook.FindByCity(city) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn()
			require.Empty(t, got)
		})
	}
}

func TestPhoneBook_Update(t *testing.T) {
	phoneBook := New()
	oldPrefix, updatedPrefix, oldCity, updatedCity := "00", "01", "Dummy City", "Updated City"
	old := Contact{
		Number:    oldPrefix + "00000000",
		FirstName: "Dummy",
		LastName:  "Dummy",
		Address:   newAddress(oldCity),
	}
	updated := Contact{
		Number:    updatedPrefix + "23456789",
		FirstName: "Foo",
		LastName:  "Bar",
		Address:   newAddress(updatedCity),
	}
	require.NoError(t, phoneBook.Add(old))
	require.NoError(t, phoneBook.Update(old.Number, updated))

	tests := []struct {
		name  string
		fn    func() any
		found bool
		want  any
	}{
		{
			name: "get updated contact success",
			fn: func() any {
				contact, _ := phoneBook.Get(updated.Number)
				return contact
			},
			found: true,
			want:  updated,
		},
		{
			name: "get old contact fail",
			fn: func() any {
				contact, _ := phoneBook.Get(old.Number)
				return contact
			},
			found: false,
		},
		{
			name:  "find updated contact by prefix success",
			fn:    func() any { return phoneBook.FindByPrefix(updatedPrefix) },
			found: true,
			want:  []Contact{updated},
		},
		{
			name:  "find old contact by prefix fail",
			fn:    func() any { return phoneBook.FindByPrefix(oldPrefix) },
			found: false,
		},
		{
			name:  "find updated contact by name (first) success",
			fn:    func() any { return phoneBook.FindByName(updated.FirstName, "") },
			found: true,
			want:  []Contact{updated},
		},
		{
			name:  "find old contact by name (first) fail",
			fn:    func() any { return phoneBook.FindByName(old.FirstName, "") },
			found: false,
		},
		{
			name:  "find updated contact by name (last) success",
			fn:    func() any { return phoneBook.FindByName("", updated.LastName) },
			found: true,
			want:  []Contact{updated},
		},
		{
			name:  "find old contact by name (last) fail",
			fn:    func() any { return phoneBook.FindByName("", old.LastName) },
			found: false,
		},
		{
			name:  "find updated contact by name (first & last) success",
			fn:    func() any { return phoneBook.FindByName(updated.FirstName, updated.LastName) },
			found: true,
			want:  []Contact{updated},
		},
		{
			name:  "find old contact by name (first & last) fail",
			fn:    func() any { return phoneBook.FindByName(old.LastName, old.LastName) },
			found: false,
		},
		{
			name:  "find updated contact by city success",
			fn:    func() any { return phoneBook.FindByCity(updatedCity) },
			found: true,
			want:  []Contact{updated},
		},
		{
			name:  "find old contact by city fail",
			fn:    func() any { return phoneBook.FindByCity(oldCity) },
			found: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn()
			if tt.found {
				require.Equal(t, tt.want, got)
			} else {
				require.Empty(t, got)
			}
		})
	}
}

func TestPhoneBook_Update_contactExistsError(t *testing.T) {
	phoneBook := New()

	oldPrefix, updatedPrefix := "00", "01"

	old := Contact{
		Number:    oldPrefix + "00000000",
		FirstName: "Dummy",
		LastName:  "Dummy",
	}

	existing := Contact{
		Number:    updatedPrefix + "23456789",
		FirstName: "Foo",
		LastName:  "Bar",
	}

	updated := existing

	require.NoError(t, phoneBook.Add(old))
	require.NoError(t, phoneBook.Add(existing))
	err := phoneBook.Update(old.Number, updated)
	require.EqualError(t, err, fmt.Sprintf("contact already exists for new number %s", existing.Number))
}

func newAddress(city string) string {
	return fmt.Sprintf("1 Foo St, %s, Foo State, 1111, Foo Country", city)
}
