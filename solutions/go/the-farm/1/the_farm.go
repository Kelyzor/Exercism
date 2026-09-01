package thefarm

import (
    "fmt"
    "errors"
)

// TODO: define the 'DivideFood' function
func DivideFood(fCalc FodderCalculator, cows int) (float64, error) {
    amount, err := fCalc.FodderAmount(cows)
    if err != nil {
        return 0, err
    }
    factor, err := fCalc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    return float64(amount) / float64(cows) * factor, err
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fCalc FodderCalculator, cows int) (float64 , error) {
    if cows > 0 {
        value, err := DivideFood(fCalc, cows)
        return value, err
    }
    return 0, errors.New("invalid number of cows")
}

// TODO: define the 'ValidateNumberOfCows' function

type InvalidCowsError struct {
  	number int
  	details string
}

func (e *InvalidCowsError) Error() string {
  	return fmt.Sprintf("%d cows are invalid: %s", e.number, e.details)
}

func ValidateNumberOfCows(cows int) error {
    if cows < 0 {
        return &InvalidCowsError {
            number: cows,
            details: "there are no negative cows",
        }
    } else if cows == 0 {
        return &InvalidCowsError {
            number: cows,
            details: "no cows don't need food",
        }
    }
    return nil
}