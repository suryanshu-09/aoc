// Package days
package days

import (
	"fmt"
	"strconv"
	"strings"
)

func Day1_1(input string) {
	dial := 50
	password := 0
	for rotation := range strings.SplitSeq(input, "\n") {
		if len(rotation) == 0 {
			break
		}
		if rotation[0] == 'L' {
			rot, err := strconv.Atoi(rotation[1:])
			if err != nil {
				panic(fmt.Errorf("rotation conv err: %v", err))
			}
			dial = (((dial - rot) % 100) + 100) % 100
			if dial == 0 {
				password++
			}
		} else {
			rot, err := strconv.Atoi(rotation[1:])
			if err != nil {
				panic(fmt.Errorf("rotation conv err: %v", err))
			}
			dial = (dial + rot) % 100
			if dial == 0 {
				password++
			}
		}
	}
	fmt.Println(password)
}

func Day1_2(input string) {
	dial := 50
	password := 0
	for rotation := range strings.SplitSeq(input, "\n") {
		if len(rotation) == 0 {
			break
		}
		rot, err := strconv.Atoi(rotation[1:])
		if err != nil {
			panic(fmt.Errorf("rotation conv err: %v", err))
		}
		password += int(rot / 100)
		rot %= 100
		if rotation[0] == 'L' {
			if dial != 0 && dial-rot < 0 {
				password++
			}

			dial = (dial - rot + 100) % 100
		} else {
			if dial != 0 && rot+dial > 100 {
				password++
			}
			dial = (dial + rot) % 100
		}
		if dial == 0 {
			password++
		}
	}
	fmt.Println(password)
}
