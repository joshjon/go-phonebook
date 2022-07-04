package phonebook

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContact_Validate(t *testing.T) {
	tests := []struct {
		name    string
		contact Contact
		wantErr string
	}{
		{
			name: "valid contact",
			contact: Contact{
				Number:    "0123456789",
				FirstName: "foo",
				LastName:  "bar",
			},
		},
		{
			name: "phone number contains invalid chars",
			contact: Contact{
				Number:    "0123K56P89",
				FirstName: "foo",
				LastName:  "bar",
			},
			wantErr: "phone number must contain 10 digits",
		},
		{
			name: "phone number not length 10",
			contact: Contact{
				Number:    "012345678",
				FirstName: "foo",
				LastName:  "bar",
			},
			wantErr: "phone number must contain 10 digits",
		},
		{
			name: "first name empty",
			contact: Contact{
				Number:    "0123456789",
				FirstName: "",
				LastName:  "bar",
			},
			wantErr: "first name required",
		},
		{
			name: "last name empty",
			contact: Contact{
				Number:    "0123456789",
				FirstName: "foo",
				LastName:  "",
			},
			wantErr: "last name required",
		},
		{
			name: "invalid address format",
			contact: Contact{
				Number:    "0123456789",
				FirstName: "foo",
				LastName:  "bar",
				Address:   "11 Fake St, Fake City, Fake State, 1111",
			},
			wantErr: "address must be in the format '[street address], [city], [state/province], [zip code], [country]'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.contact.Validate()
			if tt.wantErr != "" {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
