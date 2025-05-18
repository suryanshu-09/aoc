package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/suryanshu-09/aoc/days"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	days.Day1_0(getInput(1))
	days.Day1_1(getInput(1))
	days.Day2_0(getInput(2))
}

func getInput(day int) string {
	inputDir := "inputs"
	inputFile := filepath.Join(inputDir, fmt.Sprintf("day%d.txt", day))

	if _, err := os.Stat(inputFile); err == nil {
		content, _ := os.ReadFile(inputFile)
		return string(content)
	}
	sessionCookie := os.Getenv("SESSION")
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, _ := http.NewRequest("GET", url, nil)

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "Got Error"
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	_ = os.MkdirAll(inputDir, os.ModePerm)
	_ = os.WriteFile(inputFile, body, os.ModePerm)

	return string(body)
}
