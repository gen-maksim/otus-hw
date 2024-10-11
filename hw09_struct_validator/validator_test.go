package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:6"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Order struct {
		Items int `validate:"max:a5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: App{Version: "123456"},
			expectedErr: ValidationErrors{ValidationError{
				Field: "Version",
				Err:   ErrorLen,
			}},
		},
		{
			in: Response{Code: 300, Body: ""},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Code", Err: ErrorIn},
			},
		},
		{
			in:          Response{Code: 200, Body: ""},
			expectedErr: nil,
		},
		{
			in: User{
				ID:     "123423",
				Name:   "John",
				Age:    88,
				Email:  "@.@mail",
				Role:   "worker",
				Phones: []string{"1234567891011", "123456789101121"},
			},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Age", Err: ErrorMax},
				ValidationError{Field: "Email", Err: ErrorRegex},
				ValidationError{Field: "Role", Err: ErrorIn},
				ValidationError{Field: "Phones", Err: ErrorLen},
				ValidationError{Field: "Phones", Err: ErrorLen},
			},
		},
		{
			in: Order{
				Items: 12,
			},
			expectedErr: strconv.ErrSyntax,
		},
		{
			in:          1234,
			expectedErr: nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)
			require.ErrorIs(t, err, tt.expectedErr)
		})
	}
}
