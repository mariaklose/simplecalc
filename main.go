package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//1) Defining reader
	var reader *bufio.Reader
	reader = bufio.NewReader(os.Stdin)

	//2) Asking user for operation
	fmt.Println("Enter operation 1) sum 2) rest 3) multiplication 4) division")
	//3) Reading operation from user
	var operation string
	operation, _ = reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	//4) Showing operation to the user
	fmt.Println("operation is:", operation)
	//5) Reading Operad1
	var operand1 float64
	var operand1Str string
	fmt.Println("Operand 1:")
	operand1Str, _ = reader.ReadString('\n')
	operand1Str = strings.TrimSpace(operand1Str)
	operand1, _ = strconv.ParseFloat(operand1Str, 64)
	//6) Reading Operand2
	var operand2 float64
	var operand2Str string
	fmt.Println("Operand 2:")
	operand2Str, _ = reader.ReadString('\n')
	operand2Str = strings.TrimSpace(operand2Str)
	operand2, _ = strconv.ParseFloat(operand2Str, 64)
	//7) Performing Operation
	if operation == "sum" {
		fmt.Println("sum result is:", sum(operand1, operand2))
	} else if operation == "rest" {
		fmt.Println("rest result is:", rest(operand1, operand2))

	} else if operation == "multiplication" {
		fmt.Println("multiplication result is:", multiplication(operand1, operand2))
	} else if operation == "division" {
		fmt.Println("division result is:", division(operand1, operand2))
	}
}

//avaiable operations for numbers +, -, *,/
func sum(operand1 float64, operand2 float64) float64 {
	return operand1 + operand2
}

func rest(operand1 float64, operand2 float64) float64 {
	return operand1 - operand2
}

func multiplication(operand1 float64, operand2 float64) float64 {
	return operand1 * operand2

}

func division(operand1 float64, operand2 float64) float64 {
	return operand1 / operand2
}
