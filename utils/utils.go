package utils

import (
	"calc-oop/structures"
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

func DebugLog(currentSymbol string, stack structures.Stack, queue structures.Queue) {
	fmt.Println("current symbol")
	fmt.Println(currentSymbol)
	fmt.Println("stack")
	fmt.Println(stack.All())
	fmt.Println("queue")
	fmt.Println(queue.All())
}





