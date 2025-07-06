package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// parsing the input flags values
	isInteractive := flag.Bool("i", false, "Paste input manually (interactive mode)")
	filePath := flag.String("f", "", "Path to input file")
	flag.Parse()

	// checking the stdin piped status
	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error checking stdin: %v\n", err)
		os.Exit(1)
	}
	isPiped := stat.Mode()&os.ModeCharDevice == 0

	// checking the input method count
	method, err := inputMethodChecker(*filePath, *isInteractive, isPiped)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// receiving the input
	content, err := receiveInput(method, *filePath)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	result := strings.Join(*content, "&")
	fmt.Println(result)
}

func inputMethodChecker(filePath string, isInteractive bool, isPiped bool) (int, error) {
	var count int
	var method int

	// number 1
	if filePath != "" {
		count++
		method = 1
	}

	// number 2
	if isInteractive {
		count++
		method = 2
	}

	// number 3
	if isPiped {
		count++
		method = 3
	}

	if count != 1 {
		return 0, errors.New("Error: Use exactly one input method: -i, -f <file>, or piped stdin.")
	}

	return method, nil
}

func receiveInput(method int, filePath string) (*[]string, error) {
	switch method {
	case 1:
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("Error: %w\n", err)
		}

		result := strings.Split(string(content), "\n")

		return &result, nil

	case 2:
		var lines []string
		fmt.Println("Paste your wordlist and press Ctrl+D when done:")

		reader := bufio.NewReader(os.Stdin)
		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			line := scanner.Text()
			if line != "" {
				lines = append(lines, line)
			}
		}

		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("Error reading pasted input: %w\n", err)
		}

		return &lines, nil

	case 3:
		var lines []string

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			lines = append(lines, line)
		}

		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("Error reading piped input: %w\n", err)
		}

		return &lines, nil
	}

	return nil, nil
}
