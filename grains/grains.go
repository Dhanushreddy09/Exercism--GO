package grains

import (
	"errors"
)

var sum uint64

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("should not be less than zero")
	}
	ans := 1 << (number - 1)
	return uint64(ans), nil
}

func Total() uint64 {
	return uint64((1 << (64)) - 1)
}
