package util

import (
	"math/rand"
	"strconv"
	"time"
)

func StrToUint(str string) (uint, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}

var digits = [...]byte{
	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
}

func CodeGen() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	pwd := make([]byte, 6)
	for j := 0; j < 6; j++ {
		index := r.Int() % len(digits)
		pwd[j] = digits[index]
	}

	return string(pwd)
}
