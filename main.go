package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type options struct {
	count      bool
	duplicates bool
	unique     bool
	ignoreCase bool
	skipFields int
	skipChars  int
}

func processLine(line string, opts *options) string {
	if opts.skipFields > 0 {
		fields := strings.Fields(line)
		if len(fields) > opts.skipFields {
			line = strings.Join(fields[opts.skipFields:], " ")
		} else {
			line = ""
		}
	}
	if opts.skipChars > 0 && len(line) > opts.skipChars {
		line = line[opts.skipChars:]
	}
	if opts.ignoreCase {
		line = strings.ToLower(line)
	}
	return line
}

func readLines(reader io.Reader, opts *options) ([]string, map[string]int) {
	scanner := bufio.NewScanner(reader)
	lines := []string{}
	lineCounts := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		processed := processLine(line, opts)
		lines = append(lines, processed)
		if processed != "" {
			lineCounts[processed]++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
	return lines, lineCounts
}

func outputLines(lines []string, lineCounts map[string]int, opts *options) {
	seen := make(map[string]bool)

	for i, line := range lines {
		count := lineCounts[line]
		if line == "" {
			fmt.Println()
		} else {
			if opts.count {
				if !seen[line] {
					fmt.Printf("%d %s\n", count, line)
					seen[line] = true
				}
			} else if opts.duplicates && count > 1 {
				if !seen[line] {
					fmt.Println(line)
					seen[line] = true
				}
			} else if opts.unique && count == 1 {
				if !seen[line] {
					fmt.Println(line)
					seen[line] = true
				}
			} else if !opts.count && !opts.duplicates && !opts.unique {
				if !seen[line] {
					fmt.Println(line)
					seen[line] = true
				}
			}
		}

		if i == len(lines)-1 && line != "" {
			fmt.Println(line)
		}
	}
}

func main() {
	countFlag := flag.Bool("c", false, "Подсчитать количество повторений строк")
	duplicateFlag := flag.Bool("d", false, "Вывести только повторяющиеся строки")
	uniqueFlag := flag.Bool("u", false, "Вывести только уникальные строки")
	ignoreCaseFlag := flag.Bool("i", false, "Игнорировать регистр")
	skipFieldsFlag := flag.Int("f", 0, "Пропустить первые num_fields полей")
	skipCharsFlag := flag.Int("s", 0, "Пропустить первые num_chars символов")

	flag.Parse()

	if (*countFlag && *duplicateFlag) || (*countFlag && *uniqueFlag) || (*duplicateFlag && *uniqueFlag) {
		fmt.Fprintln(os.Stderr, "Нельзя использовать -c, -d, -u одновременно")
		os.Exit(1)
	}

	opts := &options{
		count:      *countFlag,
		duplicates: *duplicateFlag,
		unique:     *uniqueFlag,
		ignoreCase: *ignoreCaseFlag,
		skipFields: *skipFieldsFlag,
		skipChars:  *skipCharsFlag,
	}

	var input io.Reader = os.Stdin
	if flag.NArg() > 0 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	lines, lineCounts := readLines(input, opts)

	if flag.NArg() > 1 {
		file, err := os.Create(flag.Arg(1))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error creating output file:", err)
			os.Exit(1)
		}
		defer file.Close()
		os.Stdout = file
	}

	outputLines(lines, lineCounts, opts)
}
