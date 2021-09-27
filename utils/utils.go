package utils

import (
	"fmt"
	"strconv"
)

func IsDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func IsFloat(s string) bool {
	const bitSize = 64 // Don't think about it to much. It's just 64 bits.
	_, err := strconv.ParseFloat(s, bitSize)
	if err == nil {
		return true
	}
	return false
}

func IsContainsFunc(x string) bool {
	arr := []string{"*", "/", "+", "-"}
	for _, n := range arr {
		if x == n {
			return true
		}
	}
	return false
}

func DebugLog(data ICustomType) {
	v, err := data.ToString()
	if err != nil {
		fmt.Println("logging: " + v)
	} else {
		fmt.Println("logging error:")
		fmt.Println(err)
	}
}





