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

	// input flags
	isInputInteractive := flag.Bool("i", false, "Paste input manually (interactive mode)")
	filePath := flag.String("f", "", "Path to input file")
	// output flags
	isOutputInteractive := flag.Bool("oi", false, "Each chunk will be displayed on user interaction")
	chunkSize := flag.Int("n", 0, "Size of each chunk (0 = all at once)")

	flag.Parse()

	// checking the stdin piped status
	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error checking stdin: %v\n", err)
		os.Exit(1)
	}
	isPiped := stat.Mode()&os.ModeCharDevice == 0

	// checking the input method count
	method, err := inputMethodChecker(*filePath, *isInputInteractive, isPiped)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// receiving the input
	content, err := receiveInput(method, *filePath)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	showOutput(content, *isOutputInteractive, *chunkSize)
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

func receiveInput(method int, filePath string) ([]string, error) {
	switch method {
	case 1:
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("Error: %w\n", err)
		}
		lines := strings.Split(string(content), "\n")
		return cleanLines(lines), nil

	case 2:
		fmt.Println("Paste your wordlist and press Ctrl+D when done:")
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("Error reading pasted input: %w\n", err)
		}
		return cleanLines(lines), nil

	case 3:
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("Error reading piped input: %w\n", err)
		}
		return cleanLines(lines), nil

	default:
		return nil, errors.New("invalid input method")
	}
}

func cleanLines(lines []string) []string {
	var cleaned []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			cleaned = append(cleaned, line)
		}
	}
	return cleaned
}

func showOutput(content []string, isInteractive bool, chunkSize int) {
	if chunkSize <= 0 {
		// Output all at once
		fmt.Println(buildParamString(content))
		return
	}

	// Chunked output
	chunks := chunkParams(content, chunkSize)

	if isInteractive {
		reader := bufio.NewReader(os.Stdin)
		for i, chunk := range chunks {
			fmt.Printf("Chunk %d:\n%s\n", i+1, chunk)
			if i != len(chunks)-1 {
				fmt.Print("[press Enter to continue]")
				reader.ReadString('\n')
			}
		}
	} else {
		// Output all at once with separator
		fmt.Println(strings.Join(chunks, "\n---\n"))
	}
}

func buildParamString(lines []string) string {
	var builder strings.Builder
	for i, param := range lines {
		if i > 0 {
			builder.WriteString("&")
		}
		builder.WriteString(fmt.Sprintf("%s=XNLV%d", param, i+1))
	}
	return builder.String()
}

func chunkParams(lines []string, size int) []string {
	var chunks []string
	for i := 0; i < len(lines); i += size {
		end := i + size
		if end > len(lines) {
			end = len(lines)
		}
		chunk := buildParamString(lines[i:end])
		chunks = append(chunks, chunk)
	}
	return chunks
}
