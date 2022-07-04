package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var input [3]string
	fmt.Scan(&input[0], &input[1], &input[2])
	result, err := kalkulator(input[:])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func kalkulator(input []string) (int, error) {
	operand1 := input[0]
	operand2 := input[1]
	operand3 := input[2]

	op1, _ := strconv.Atoi(operand1)
	op3, _ := strconv.Atoi(operand3)

	if operand2 == "+" {
		return op1 + op3, nil
	} else if operand2 == "-" {
		return op1 - op3, nil
	} else if operand2 == "*" {
		return op1 * op3, nil
	} else if operand2 == "/" {
		return op1 / op3, nil
	} else {
		return 0, errors.New("operand tidak diketahui")
	}
}
