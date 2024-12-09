package field_test

import (
	"Battleship/internal/field"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fieldTemplate = []int{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

var fieldCorrect = []int{
	1, 0, 0, 0, 0, 1, 1, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 1, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 1, 1,
	0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 0, 0, 0, 0, 1, 0, 0,
	0, 0, 1, 0, 0, 1, 0, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

var fieldIncorrect1 = []int{
	1, 0, 0, 0, 0, 1, 1, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 1, 0, 0, 0,
	1, 1, 0, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 1, 1,
	0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1, 0, 0,
	0, 0, 1, 0, 0, 1, 0, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

var fieldIncorrect2 = []int{
	1, 0, 0, 0, 0, 1, 1, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 0, 1, 0, 0,
	1, 0, 0, 1, 0, 0, 1, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 1, 1,
	0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1, 0, 0,
	0, 0, 1, 0, 0, 1, 0, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

var fieldIncorrectNo4SizeShip = []int{
	1, 0, 0, 0, 0, 1, 1, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 1, 0, 0, 1,
	0, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 1, 1,
	0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 0, 0, 0, 0, 1, 0, 0,
	0, 0, 1, 0, 0, 1, 0, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

var fieldIncorrectTooManyShips = []int{
	1, 0, 0, 0, 0, 1, 1, 0, 0, 1,
	1, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 1, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 1, 1,
	0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 0, 0, 0, 0, 1, 0, 0,
	0, 0, 1, 0, 0, 1, 0, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

var fieldIncorrectNotEnoughShips = []int{
	1, 0, 0, 0, 0, 1, 1, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 1, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 1, 1,
	0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 1, 0, 0, 1, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

var fieldIncorrectStruct = []int{
	1, 0, 0, 0, 0, 1, 1, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 1, 0, 0, 1, 0, 0, 1,
	0, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 1, 1,
	0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 0, 0, 0, 0, 1, 0, 0,
	0, 0, 1, 0, 0, 1, 0, 1, 0, 0,
}

func TestCheckField(t *testing.T) {
	for _, tc := range []struct {
		name          string
		field         []int
		expectedError error
	}{
		{"fully correct", fieldCorrect, nil},
		{"incorrect field, ships are side to side", fieldIncorrect1, field.ErrFieldIncorrectShipsPositions},
		{"incorrect field, ships are diagonally connected", fieldIncorrect2, field.ErrFieldIncorrectShipsPositions},
		{"no 4 sells size ship", fieldIncorrectNo4SizeShip, field.ErrIncorrectShipsNumber},
		{"too many ships", fieldIncorrectTooManyShips, field.ErrFieldTooManyShips},
		{"not enough ships", fieldIncorrectNotEnoughShips, field.ErrFieldNotEnoughShips},
		{"Incorrect struct", fieldIncorrectStruct, field.ErrFieldStructError},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.ErrorIs(t, field.CheckField(tc.field), tc.expectedError)
		})
	}
}
