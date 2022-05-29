package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	message string
	cows    int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	fodder, err := weightFodder.FodderAmount()
	if fodder < 0 {
		if err == ErrScaleMalfunction || err == nil {
			return 0, errors.New("negative fodder")
		} else {
			return 0, errors.New("non-scale error")
		}
	} else if cows == 0 {
		return 0, errors.New("division by zero")
	} else if cows < 0 {
		return 0, &SillyNephewError{"", cows}
	}
	if err != nil {
		if err == ErrScaleMalfunction && fodder > 0 {
			fodder *= 2
		} else {
			return 0, err
		}
	}

	return fodder / float64(cows), nil
}
