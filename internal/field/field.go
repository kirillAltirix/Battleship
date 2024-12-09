package field

import "errors"

var ErrFieldStructError = errors.New("incorrect number of elements in field struct")
var ErrFieldIncorrectShipsPositions = errors.New("incorrect ships positions")
var ErrIncorrectShipsNumber = errors.New("incorrect ships number")
var ErrFieldTooManyShips = errors.New("too many ships")
var ErrFieldNotEnoughShips = errors.New("not enough ships")
var ErrOutOfBounds = errors.New("incorrect index")

const numSheep1SellSize = 4
const numSheep2SellSize = 3
const numSheep3SellSize = 2
const numSheep4SellSize = 1
const fleetSize = 4*numSheep4SellSize + 3*numSheep3SellSize + 2*numSheep2SellSize + 1*numSheep1SellSize

type FieldImpl interface {
	GetAt(i, j int) (int, error)
}

func CheckField(field FieldImpl) error {
	shipSizes := make(map[int]int)
	totalFleetSize := 0
	checkedSells := make([]int, 100)

	defaultUpdater := 0
	updaters := []func(ii int, jj int) (int, int){
		func(ii int, jj int) (int, int) { return ii, jj },
		func(ii int, jj int) (int, int) { return ii, jj + 1 },
		func(ii int, jj int) (int, int) { return ii, jj - 1 },
		func(ii int, jj int) (int, int) { return ii + 1, jj },
		func(ii int, jj int) (int, int) { return ii - 1, jj },
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			shipSizeTotal := 0
			for updaterID, updater := range updaters {
				ii, jj := updater(i, j)
				shipSize, res := checkSells(field, checkedSells, ii, jj, updater)
				if !res {
					return ErrFieldIncorrectShipsPositions
				}
				if updaterID == defaultUpdater && shipSize == 0 {
					break
				}
				shipSizeTotal += shipSize
			}
			if shipSizeTotal != 0 {
				shipSizes[shipSizeTotal]++
				totalFleetSize += shipSizeTotal
			}
		}
	}

	if totalFleetSize > fleetSize {
		return ErrFieldTooManyShips
	} else if totalFleetSize < fleetSize {
		return ErrFieldNotEnoughShips
	}

	if shipSizes[1] != numSheep1SellSize ||
		shipSizes[2] != numSheep2SellSize ||
		shipSizes[3] != numSheep3SellSize ||
		shipSizes[4] != numSheep4SellSize {
		return ErrIncorrectShipsNumber
	}

	return nil
}

func checkSells(field FieldImpl, checkedSells []int, i, j int, ijUpdater func(int, int) (int, int)) (int, bool) {
	if checkSell(field, checkedSells, i, j) == 0 {
		return 0, true
	}
	for _, ii := range []int{i - 1, i + 1} {
		for _, jj := range []int{j - 1, j + 1} {
			if checkSell(field, checkedSells, ii, jj) == 1 {
				return 0, false
			}
		}
	}

	i, j = ijUpdater(i, j)

	count, res := checkSells(field, checkedSells, i, j, ijUpdater)
	return count + 1, res
}

func checkSell(field FieldImpl, checkedSells []int, i, j int) int {
	res, err := field.GetAt(i, j)
	if err != nil {
		return 0
	}

	if checkedSells[i*10+j] == 1 {
		return 0
	}

	defer func() {
		checkedSells[i*10+j] = 1
	}()

	if res == 0 {
		return 0
	}

	return 1
}
