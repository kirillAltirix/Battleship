package field_test

import (
	"Battleship/internal/field"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// var fieldTemplate = []int{
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// }

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
	} {
		t.Run(tc.name, func(t *testing.T) {
			f, err := field.New(tc.field)
			require.NoError(t, err)
			assert.ErrorIs(t, field.CheckField(f), tc.expectedError)
		})
	}
}

func TestNew(t *testing.T) {
	_, err := field.New(fieldCorrect)
	require.NoError(t, err)
	_, err = field.New(fieldIncorrectStruct)
	require.ErrorIs(t, err, field.ErrFieldStructError)
}
