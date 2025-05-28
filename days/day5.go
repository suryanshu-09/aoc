package days

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day5_1(input string) {
	ruleMap := make(map[int][]int)

	rules := strings.SplitSeq(strings.TrimSpace(strings.Split(input, "\n\n")[0]), "\n")

	for rule := range rules {
		left, _ := strconv.Atoi(strings.Split(rule, "|")[0])
		right, _ := strconv.Atoi(strings.Split(rule, "|")[1])

		ruleMap[left] = append(ruleMap[left], right)
	}

	updatesString := strings.Split(strings.TrimSpace(strings.Split(input, "\n\n")[1]), "\n")
	updates := make([][]int, len(updatesString))
	for i, update := range updatesString {
		commas := strings.SplitSeq(update, ",")
		for page := range commas {
			pageNum, _ := strconv.Atoi(page)
			updates[i] = append(updates[i], pageNum)
		}
	}

	output := 0
	for _, update := range updates {
		isValid := true
	outer:
		for i := len(update) - 1; i >= 1; i-- {
			for _, up := range update[:i-1] {
				if slices.Contains(ruleMap[update[i]], up) {
					isValid = false
					break outer
				}
			}
		}
		if isValid {
			output += update[len(update)/2]
		}
	}

	fmt.Println("Output Day 5 Part 1", output)
}

func Day5_2(input string) {
	fmt.Println(input)
}
