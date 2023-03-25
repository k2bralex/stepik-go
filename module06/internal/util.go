package internal

import (
	"errors"
)

func ReverseNumber(val interface{}) (interface{}, error) {
	switch v := val.(type) {
	case int:
		res := 0
		isPositive := true
		if v < 0 {
			isPositive = false
			v *= -1
		}

		for v > 0 {
			remainder := v % 10
			res = remainder + (res * 10)
			v /= 10
		}

		if isPositive {
			return res, nil
		}
		return res * (-1), nil

	case string:
		res := ""
		for _, sub := range v {
			res = string(sub) + res
		}
		return res, nil
	default:
		return nil, errors.New("incorrect type")
	}
}
