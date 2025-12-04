package days

import (
	"fmt"
	"strings"
)

func Day4_1(input string) {
	lines := strings.Split(input, "\n")
	newLines := []string{}
	for _, line := range lines {
		if line != "\n" && line != "\t" && line != "" {
			newLines = append(newLines, line)
		}
	}
	out, _ := rollDaDoe(newLines)
	fmt.Println(out)
}

func rollDaDoe(newLines []string) (int, []string) {
	output := 0
	for idx, line := range newLines {
		for jdx, val := range line {
			if val == '@' {
				if checkRolls(idx, jdx, newLines) {
					row := []byte(newLines[idx]) // make row mutable
					row[jdx] = '.'               // write byte
					newLines[idx] = string(row)  // write back to slice
					output++
				}
			}
		}
	}
	return output, newLines
}

func checkRolls(idx, jdx int, lines []string) bool {
	count := 0
	height := len(lines)
	width := len(lines[idx])
	if idx-1 >= 0 && jdx-1 >= 0 {
		if lines[idx-1][jdx-1] == '@' {
			count++
		}
	}
	if idx-1 >= 0 {
		if lines[idx-1][jdx] == '@' {
			count++
		}
	}
	if idx-1 >= 0 && jdx+1 < width {
		if lines[idx-1][jdx+1] == '@' {
			count++
		}
	}
	if jdx-1 >= 0 {
		if lines[idx][jdx-1] == '@' {
			count++
		}
	}
	if jdx+1 < width {
		if lines[idx][jdx+1] == '@' {
			count++
		}
	}
	if idx+1 < height && jdx-1 >= 0 {
		if lines[idx+1][jdx-1] == '@' {
			count++
		}
	}
	if idx+1 < height {
		if lines[idx+1][jdx] == '@' {
			count++
		}
	}
	if idx+1 < height && jdx+1 < width {
		if lines[idx+1][jdx+1] == '@' {
			count++
		}
	}
	return count < 4
}

func Day4_2(input string) {
	lines := strings.Split(input, "\n")
	newLines := []string{}
	for _, line := range lines {
		if line != "\n" && line != "\t" && line != "" {
			newLines = append(newLines, line)
		}
	}
	out := 0
	sum := 2
	inp := newLines
	for sum != 0 {
		sum, inp = rollDaDoe(inp)
		out += sum
	}
	fmt.Println(out)
}
