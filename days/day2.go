package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func repeats(num int) bool {
	n := numLen(num)
	return num%int(math.Pow(float64(10), float64(n/2))) == num/int(math.Pow(float64(10), float64(n/2)))
}

func repeats2(num int) bool {
	nStr := strconv.Itoa(num)
	n := len(nStr)
	for i := n / 2; i > 0; i-- {
		if strings.Repeat(nStr[:i], n/i) == nStr {
			return true
		}
	}
	return false
}

func numLen(num int) int {
	l := 0
	for num > 0 {
		num = num / 10
		l++
	}
	return l
}

func Day2_1(input string) {
	sum := 0
	for r := range strings.SplitSeq(input, ",") {
		r = strings.ReplaceAll(r, "\n", "")
		r = strings.ReplaceAll(r, "\t", "")
		ran := strings.Split(r, "-")
		start, _ := strconv.Atoi(ran[0])
		end, _ := strconv.Atoi(ran[1])
		for i := start; i <= end; i++ {
			if repeats(i) {
				sum += i
			}
		}
	}
	fmt.Println(sum)
}

func Day2_2(input string) {
	sum := 0
	for r := range strings.SplitSeq(input, ",") {
		r = strings.ReplaceAll(r, "\n", "")
		r = strings.ReplaceAll(r, "\t", "")
		ran := strings.Split(r, "-")
		start, _ := strconv.Atoi(ran[0])
		end, _ := strconv.Atoi(ran[1])
		for i := start; i <= end; i++ {
			if repeats2(i) {
				sum += i
			}
		}
	}
	fmt.Println(sum)
}
