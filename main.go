package main

import (
	"bufio"
	"encoding/json"
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
		fmt.Println(str)

		//validator.Validate()
		rpnArr, err := Convert(str)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("reverse polish notation:")
		fmt.Println(rpnArr)
		fmt.Println("calculate...")

		result, err := Calculate(rpnArr)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(result)
	}
}

func init() {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error load config:", err)
	}
	fmt.Println(configuration.IsDebug) // output: [UserA, UserB]
}

func start() {
	fmt.Println("Calculation of an arbitrary notation with output in reverse polish notation")
	fmt.Println("\\t example:\\t 10*(2+5)-14/(1+2*(1+2)) \\n (1 + 2) * 4 + 3 \\n 2.2 * 10 - 15 \\n -5 + 5 * 6")
	fmt.Println("input expression and press Enter")
}

/**
Read input string
*/
func readInputStr() (string, error) {
	in := bufio.NewReader(os.Stdin)
	scanStr, err := in.ReadString('\n')
	fmt.Println("clear space...")
	scanStr = strings.Replace(scanStr, " ", "", -1)
	fmt.Println("replace ',' to '.'...")
	scanStr = strings.Replace(scanStr, ",", ".", -1)
	if err != nil {
		return "", err
	}

	return scanStr, nil
}
