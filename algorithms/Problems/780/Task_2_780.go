package algorithms

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunTask_2() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Number Of Calculations : ")
	input, _ := reader.ReadString('\n')
	numeberOfCalcs := readNumber(input)
	res, _ := recursiveReading(numeberOfCalcs)
	fmt.Println("Result :", res)
}

func recursiveReading(counter int) (int, int) {
	res := 0
	if counter == 0 {
		return res, counter
	}
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = ReplaceSplitters(input)
	splitedInput := strings.Split(input, " ")
	out, _ := recursiveAddNumbers(splitedInput)
	res += out
	counter--
	if counter > 0 {
		var out int
		out, counter = recursiveReading(counter)
		res += out
	}
	return res, counter
}

func recursiveAddNumbers(input []string) (int, []string) {
	res := 0
	if len(input) > 0 {
		out := 0
		res += readNumber(input[0])
		out, _ = recursiveAddNumbers(input[1:])
		res += out
	}
	return res, nil
}

func ReplaceSplitters(input string) string {
	input = strings.ReplaceAll(input, ",", " ")
	input = strings.ReplaceAll(input, "+", " ")
	return input
}
 
func readNumber(input string) int {
	res := 0
	_, err := fmt.Sscan(input, &res)
	if err != nil {
		fmt.Println("Wrong Syntax")
		return 0
	}
	return res
}
