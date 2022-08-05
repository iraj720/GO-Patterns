package algorithms

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var Words = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
}

func RunTask_1() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	sentence, _ := reader.ReadString('\n')
	_, res, isOk := StartR(sentence)
	if !isOk {
		fmt.Println("Grammer Is Incorrect")
	} else {
		fmt.Println(res)
	}
}

func StartR(sentence string) ([]string, int, bool) {
	words := strings.Split(sentence, " ")
	if !isVerb(words[0]) {
		return words, 0, false
	}
	return recursiveSentenceToNumber(words, words[0])
}


// this function is recursive and all of the logic is here
// the algorithm has some exceptions so the length of function increased to handle them
// such as abcd dede a b should return incorrect
// abcd dede a dede a b should return incorrect
// and any other problem with more than 2 consecutive verbs
func recursiveSentenceToNumber(words []string, verb string) ([]string, int, bool) {
	sum := make([]int, 0)
	isOk := true
	if isVerb(words[0]) {
		for len(words) > 2 {
			res := 0
			if isVerb(words[0]) && isVerb(words[1]) {
				if len(words) > 3 {
					if isVerb(words[2]) || isVerb(words[3]) || len(words) == 4 {
						isOk = false
						return words, 0, isOk
					}
				}
				words, res, isOk = recursiveSentenceToNumber(words[1:], words[0])
				sum = append(sum, res)
				break
			} else {
				words, res, isOk = recursiveSentenceToNumber(words[1:], words[0])
				sum = append(sum, res)
			}
			if !isOk {
				return words, 0, isOk
			}
		}
	} else if len(words) > 1 {
		numbers := make([]int, 0)
		for i, val := range words {
			if isVerb(words[i]) {
				return words[i:], doVerb(verb, numbers...), true
			} else {
				numbers = append(numbers, translateWordToNumber(val))
			}
		}
		return words[:0], doVerb(verb, numbers...), isOk
	}
	return words, doVerb(verb, sum...), isOk
}

func translateWordToNumber(word string) int {
	sameAlphabets := make([]int, 1)
	alphabetcounter := 0
	var lastAlphabet rune
	for i, val := range word {
		if i == 0 {
			sameAlphabets[0] += Words[val]
		} else {
			if val == lastAlphabet {
				sameAlphabets[alphabetcounter] += Words[val]
			} else {
				alphabetcounter++
				sameAlphabets = append(sameAlphabets, 0)
				sameAlphabets[alphabetcounter] += Words[val]
			}
		}
		lastAlphabet = val
	}
	sum := 0
	for i := range sameAlphabets {
		sum += int(math.Pow(float64(sameAlphabets[i]%5), 2))
	}
	return sum
}

func doVerb(verb string, data ...int) int {
	switch verb {
	case "abcd":
		res := 0
		for _, val := range data {
			res += val
		}
		return res
	case "bcde":
		res := data[0]
		data = data[1:]
		for _, val := range data {
			res -= val
		}
		return res
	case "dede":
		res := 1
		for _, val := range data {
			res *= val
		}
		return res

	}
	return 0
}

func isVerb(verb string) bool {
	if verb == "abcd" || verb == "bcde" || verb == "dede" {
		return true
	}
	return false
}
