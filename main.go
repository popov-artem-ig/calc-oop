package main

import (
	"bufio"
	"calc-oop/structures"
	"fmt"
	"os"
	"strings"
)

func main() {

	for true {
		start()
		str, err := readInputStr()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("'"+str+"'")

		notation := structures.ReversePolishNotation{}
		rpn := notation.New(str)


		fmt.Println("reverse polish notation:")
		rpnString, _ := rpn.ToString()
		fmt.Println(rpnString)


		fmt.Println("calculate...")
		result, err := rpn.Calc()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(result)
	}
}

func start() {
	fmt.Println("Calculation of an arbitrary notation with output in reverse polish notation")
	fmt.Println("example:\n 10*(2+5)-14/(1+2*(1+2)) \n (1 + 2) * 4 + 3 \n 2.2 * 10 - 15 \n -5 + 5 * 6")
	fmt.Println("input expression and press Enter")
}

/**
Read input string
*/
func readInputStr() (string, error) {
	in := bufio.NewReader(os.Stdin)
	scanStr, err := in.ReadString('\n')
	fmt.Println("clear space...")
	scanStr = strings.TrimSuffix(scanStr, "\n")
	scanStr = strings.Replace(scanStr, " ", "", -1)
	fmt.Println("replace ',' to '.'...")
	scanStr = strings.Replace(scanStr, ",", ".", -1)
	if err != nil {
		return "", err
	}

	return scanStr, nil
}
