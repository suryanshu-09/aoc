package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func strArrToIntArr(strArr []string) []int {
	var result []int
	for _, char := range strArr {
		intChar, _ := strconv.Atoi(char)
		result = append(result, intChar)
	}
	return result
}

func Day2_0(input string) {
	levels := strings.Split(strings.TrimSpace(input), "\n")
	safety := 0
	for _, report := range levels {
		reportArr := strings.Split(report, " ")
		intReport := strArrToIntArr(reportArr)
		safe := true
		for i := 1; i < len(intReport); i++ {
			prev := intReport[i-1]
			now := intReport[i]

			diff := math.Abs(float64(prev - now))
			if diff != 1 && diff != 2 && diff != 3 {
				safe = false
				break
			}
		}
		if safe {
			safety++
		}
	}

	fmt.Println("Output Day 2 Part 1", safety)
}

func Day2_1(input string) {
}
