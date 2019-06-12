package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calc(input chan string, output chan string) {
	keepRunning := true

	for keepRunning {
		inputStr := <-input
		inputL := strings.Split(inputStr, " ")

		op := inputL[1]
		//		println(inputL[2])
		var num1, num2 float64

		num1, _ = strconv.ParseFloat(inputL[0], 64)
		num2, _ = strconv.ParseFloat(inputL[2], 64)

		var result float64

		switch op {
		case "*":
			result = num1 * num2
		case "/":
			result = num1 / num2
		case "+":
			result = num1 + num2
			//			fmt.Printf("adding %f + %f = %f\n", num1, num2, result)
		case "-":
			result = num1 - num2
		}
		output <- fmt.Sprintf("%f", result)
	}
	return
}

func main() {
	keepRunning := true
	fmt.Printf("Hello - This is Silly Canopy 1.0\n")
	inputC := make(chan string)
	outputC := make(chan string)
	go calc(inputC, outputC)

	reader := bufio.NewReader(os.Stdin)
	for keepRunning {
		fmt.Printf(">")
		text, _ := reader.ReadString('\n')
		inputC <- text
		fmt.Printf("Calculated: %s\n", <-outputC)
	}

}
