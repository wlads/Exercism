package luhn

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Valid determine whether or not number is valid per the Luhn formula
func Valid(number string) (isValid bool) {
	cleanNumber := strings.Replace(number, " ", "", -1)
	fmt.Printf("\n* number = [%v], clean = [%v]", number, cleanNumber)
	if len(cleanNumber) <= 1 {
		return false
	}
	sum, err := checksum(cleanNumber)
	fmt.Printf(", sum = [%v]", sum)
	return err == nil && (sum%10 == 0)
}

func checksum(number string) (int, error) {
	digits := 1
	sum := 0
	for i := len(number) - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(number[i]))
		switch {
		case err != nil:
			return 0, errors.New("can't checksum this number: " + string(number[i]))
		case digits%2 != 0:
			sum += n
		default:
			double := 2 * n
			if double > 9 {
				double -= 9
			}
			sum += double
		}
		digits++
	}
	return sum, nil
}
