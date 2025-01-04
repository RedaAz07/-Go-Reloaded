package main

import (
	"fmt"
	"os"
	"strings"

	"piscine/piscine"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("  go run . (sample0.txt) (result.txt) ")
		return
	}

	if !strings.HasSuffix(os.Args[2], ".txt") {
		fmt.Println("Nice try  bro ")
		return
	}

	textFile, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	lines := strings.Split(string(textFile), "\n")
	var table []string
	for _, line := range lines {
		word := strings.Fields(line)

		lastResult := []string{}
		filtering := []string{}
		last := []string{}
		for {
			previous := strings.Join(word, " ")
			word = piscine.Flags(word)
			filtering = []string{}
			for _, w := range word {
				if w != "" {
					filtering = append(filtering, w)
				}
			}
			current := strings.Join(filtering, " ")
			if current == previous {
				break
			}
			word = filtering

		}

		filtering = piscine.SplitePunc(filtering)

		last = []string{}
		for _, w := range filtering {
			if w != "" {
				last = append(last, w)
			}
		}

		last = piscine.Filter(last)


		lastResult = []string{}
		for _, w := range last {
			if w != "" {
				lastResult = append(lastResult, w)
			}
		}

		for {
			previous := strings.Join(lastResult, " ")
			lastResult = piscine.Flags(lastResult)
			filtering = []string{}
			for _, w := range lastResult {
				if w != "" {
					filtering = append(filtering, w)
				}
			}
			current := strings.Join(filtering, " ")
			if current == previous {
				break
			}
			lastResult = filtering

		}




		table = append(table, strings.Join(lastResult, " "))
	}

	err = os.WriteFile(os.Args[2], []byte(strings.Join(table, "\n")), 0o644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
	}
}
