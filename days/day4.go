package days

import (
	"fmt"
	"strings"
)

func Day4_1(input string) {
	output := 0
	// re := regexp.MustCompile(`XMAS`)
	// found := re.FindAllString(input, -1)
	// output += len(found)
	//
	// re = regexp.MustCompile(`SMAX`)
	// found = re.FindAllString(input, -1)
	// output += len(found)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for i, line := range lines {
		for j, r := range line {
			switch r {
			case 'X':
				if i+3 < len(lines) {
					if lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
						output++
					}
				}
				if j+3 < len(line) {
					if line[j+1] == 'M' && line[j+2] == 'A' && line[j+3] == 'S' {
						output++
					}
				}
				if i+3 < len(lines) && j+3 < len(line) {
					if lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
						output++
					}
				}
				if i+3 < len(lines) && j-3 >= 0 {
					if lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
						output++
					}
				}
			case 'S':
				if i+3 < len(lines) {
					if lines[i+1][j] == 'A' && lines[i+2][j] == 'M' && lines[i+3][j] == 'X' {
						output++
					}
				}
				if j+3 < len(line) {
					if line[j+1] == 'A' && line[j+2] == 'M' && line[j+3] == 'X' {
						output++
					}
				}
				if i+3 < len(lines) && j+3 < len(line) {
					if lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'M' && lines[i+3][j+3] == 'X' {
						output++
					}
				}
				if i+3 < len(lines) && j-3 >= 0 {
					if lines[i+1][j-1] == 'A' && lines[i+2][j-2] == 'M' && lines[i+3][j-3] == 'X' {
						output++
					}
				}
			}
		}
	}

	fmt.Println("Output Day 4 Part 1", output)
}

func Day4_2(input string) {
	output := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for i, line := range lines {
		for j, r := range line {
			switch r {
			case 'A':
				//M.S.M
				//.A.A.
				//M.S.M
				if i-1 >= 0 && j-1 >= 0 && i+1 < len(lines) && j+1 < len(line) {
					masLD := lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S'
					masRD := lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S'
					samLD := lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M'
					samRD := lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M'
					if (masLD && (masRD || samRD)) || (samLD && (masRD || samRD)) {
						output++
					}
				}
			}
		}
	}

	fmt.Println("Output Day 4 Part 2", output)
}
