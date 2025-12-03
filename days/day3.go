package days

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Day3_1(input string) {
	sum := 0
	for battery := range strings.SplitSeq(input, "\n") {
		// fmt.Println("batter:", battery)
		maxJolt := 0
		for idx := range battery {
			// fmt.Println("idx:", idx)
			for jdx := idx + 1; jdx < len(battery); jdx++ {
				// fmt.Println("jdx:", jdx)
				// fmt.Println("batter[idx]:", battery[idx])
				// fmt.Println("batter[jdx]:", battery[jdx])
				joltage, _ := strconv.Atoi(fmt.Sprintf("%c%c", battery[idx], battery[jdx]))
				// fmt.Println("joltage:", joltage)
				maxJolt = max(joltage, maxJolt)
				// fmt.Println("maxJolt:", maxJolt)
			}
		}
		sum += maxJolt
	}
	fmt.Println(sum)
}

func Day3_2(inputStr string) {
	sum := int64(0)
	maxLen := 12

	for batteryLine := range strings.SplitSeq(inputStr, "\n") {
		line := strings.TrimSpace(batteryLine)
		if line == "" {
			continue
		}

		digits := make([]int, 0, len(line))
		for _, r := range line {
			if !unicode.IsDigit(r) {
				continue
			}
			n, _ := strconv.Atoi(string(r))
			digits = append(digits, n)
		}
		if len(digits) == 0 {
			continue
		}

		maxNum := int64(0)
		for i := maxLen - 1; i >= 0; i-- {
			prefixLen := len(digits) - i
			if prefixLen <= 0 {
				break
			}
			window := digits[:prefixLen]
			m := slices.Max(window)
			maxNum = maxNum*10 + int64(m)
			idx := slices.Index(digits, m)
			if idx+1 >= len(digits) {
				digits = digits[:0]
			} else {
				digits = digits[idx+1:]
			}
			if len(digits) == 0 {
				break
			}
		}
		sum += maxNum
	}
	fmt.Println("sum:", sum)
}
