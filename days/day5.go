package days

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day5_1(input string) {
	idRange := [][2]int{}
	ingredients := []int{}
	out := 0
	swivel := false
	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			swivel = true
			continue
		}
		if !swivel {
			r := strings.Split(line, "-")
			a, err := strconv.Atoi(r[0])
			if err != nil {
				fmt.Println(err)
			}
			b, err := strconv.Atoi(r[1])
			if err != nil {
				fmt.Println(err)
			}

			idRange = append(idRange, [2]int{a, b})
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			}
			ingredients = append(ingredients, i)
		}
	}
	for _, val := range ingredients {
		for _, r := range idRange {
			if val >= r[0] && val <= r[1] {
				out++
				break
			}
		}
	}
	fmt.Println(out)
}

func Day5_2(input string) {
	type iv struct{ a, b int64 }
	var ranges []iv

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		a, _ := strconv.ParseInt(parts[0], 10, 64)
		b, _ := strconv.ParseInt(parts[1], 10, 64)
		if a > b {
			a, b = b, a
		}
		ranges = append(ranges, iv{a, b})
	}

	if len(ranges) == 0 {
		fmt.Println(0)
		return
	}

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].a == ranges[j].a {
			return ranges[i].b < ranges[j].b
		}
		return ranges[i].a < ranges[j].a
	})

	var sum int64
	curA, curB := ranges[0].a, ranges[0].b
	for i := 1; i < len(ranges); i++ {
		a, b := ranges[i].a, ranges[i].b
		if a <= curB+1 {
			if b > curB {
				curB = b
			}
		} else {
			sum += curB - curA + 1
			curA, curB = a, b
		}
	}
	sum += curB - curA + 1
	fmt.Println(sum)
}
