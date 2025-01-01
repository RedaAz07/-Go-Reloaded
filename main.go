package main

import (
	"fmt"
	"os"
	"strings"

	"piscine/piscine"
)




func main() {



fmt.Println(piscine.SpliteQuot([]string{"'''red'a'","annis"}))




	if len(os.Args) != 3 {
		fmt.Println("  go run . (sample0.txt) (result.txt) ")
		return
	}

	if !strings.HasSuffix(os.Args[2], ".txt") {
		fmt.Print("Nice try  bro ")
		return
	}

	textFile, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

lines := strings.Split(string(textFile), "\n")

var table[]string
for i := 0; i < len(lines); i++ {
	

	worlds := strings.Fields(lines[i])



	lastResult := []string{}
	filteredWorlds := []string{}
	last := []string{}
	worlds=piscine.Flags(worlds)
	for _, word := range worlds {
		if word != "" {
			filteredWorlds = append(filteredWorlds, word)
		}
	}

	filteredWorlds = piscine.SplitePunc(filteredWorlds)
	for _, word := range filteredWorlds {
		if word != "" {
			last = append(last, word)
		}
	}

	last = piscine.Filter(last)
	for _, word := range last {
		if word != "" {
			lastResult = append(lastResult, word)
		}
	}
	table = append(table,strings.Join(lastResult," "))

}
	err = os.WriteFile(os.Args[2], []byte(strings.Join(table, "\n")), 0o644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
	}
	
}