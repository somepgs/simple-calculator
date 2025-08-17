package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter: <num1> <operator> <num2> (e.g. 1 + 2): ")
	reader.Scan()
	input := strings.Fields(reader.Text())

	if len(input) != 3 {
		fmt.Println("Error: invalid input")
		os.Exit(1)
	}

	num1, err1 := strconv.ParseFloat(input[0], 64)
	if err1 != nil {
		fmt.Println("Error: num1 is not a number")
		os.Exit(1)
	}

	num2, err2 := strconv.ParseFloat(input[2], 64)
	if err2 != nil {
		fmt.Println("Error: num2 is not a number")
		os.Exit(1)
	}

	result, err := calculate(num1, num2, input[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Result:", result)
}

func calculate(num1, num2 float64, op string) (float64, error) {
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("division by zero")
		}
		return num1 / num2, nil
	default:
		return 0, errors.New("unknown operator")
	}
}
