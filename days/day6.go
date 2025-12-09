package days

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Day6_1(input string) {
	operations := []string{}
	numbers := [][]int{}
	out := 0
	inputArr := strings.Split(input, "\n")
	for idx, line := range inputArr {
		if idx == 4 {
			for l := range strings.SplitSeq(line, " ") {
				if l != "" {
					operations = append(operations, l)
				}
			}
			break
		}
		l := strings.Split(line, " ")
		numArr := []int{}
		for _, x := range l {
			if x == "" {
				continue
			}
			num, err := strconv.Atoi(x)
			if err != nil {
				fmt.Println(err, x)
			}
			numArr = append(numArr, num)
		}
		numbers = append(numbers, numArr)
	}
	for idx, symbol := range operations {
		symbol = strings.TrimSpace(symbol)
		switch symbol {
		case "+":
			s := 0
			for _, arr := range numbers {
				s += arr[idx]
			}
			out += s
		case "*":
			p := 1
			for _, arr := range numbers {
				p *= arr[idx]
			}
			out += p
		}
	}
	fmt.Println(out)
}

const (
	Multiply = "*"
	Addition = "+"
)

func Day6_2(input string) {
	lines := strings.Split(input, "\n")

	numberLineRE := regexp.MustCompile(`^\s*\d`)

	part2Strings := make([]string, 0)
	calcs := make([]string, 0)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if numberLineRE.MatchString(line) {
			part2Strings = append(part2Strings, line)
		} else {
			for _, r := range line {
				if r == '+' || r == '*' {
					calcs = append(calcs, string(r))
				}
			}
		}
	}

	part2Nums := make([]int, 0)

	for i := len(part2Strings[0]) - 1; i >= 0; i-- {
		runningTotal := 0
		for j := 0; j < len(part2Strings); j++ {
			if part2Strings[j][i] != ' ' {
				blah, _ := strconv.Atoi(string(part2Strings[j][i]))
				runningTotal = runningTotal*10 + blah
			}
		}
		part2Nums = append(part2Nums, runningTotal)
	}

	coi := 1
	currentOperand := calcs[len(calcs)-coi]
	answer := 0
	runningTotal := math.MaxInt

	for _, n := range part2Nums {
		if n == 0 {
			coi++
			if coi <= len(calcs) {
				currentOperand = calcs[len(calcs)-coi]
			}
			answer += runningTotal
			runningTotal = math.MaxInt
			continue
		}

		if runningTotal == math.MaxInt {
			runningTotal = n
		} else {
			switch currentOperand {
			case Multiply:
				runningTotal *= n
			case Addition:
				runningTotal += n
			}
		}
	}

	answer += runningTotal
	fmt.Println(answer)
}
