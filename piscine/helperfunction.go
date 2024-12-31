package piscine

import (
	"fmt"
	"strconv"
)


func Ignoreempty(i int, counter int, table []string) int {
	for l := i - 1; l >= 0; l-- {
		if table[l] == "" {
			counter++
		} else {
			break
		}
	}
	return counter
}

// function to handle the ) in the exampels of (up, 2)
func MyAtoi(word string) int {
	correct := ""

	if word[len(word)-1] >= ')' && word[len(word)-2] >= '0' && word[len(word)-2] <= '9' {
		for i := 0; i < len(word)-1; i++ {
			correct += string(word[i])
		}
	}
	num1, err := strconv.Atoi(correct)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}

	return num1
}



// The name explane isSelfe
func TableToString(table []string) string {
	text := ""
	for i := 0; i < len(table); i++ {
		if table[i] != "" {

			text += string(table[i])
			if i < len(table)-1 {
				text += " "
			}
		}
	}
	return text
}

func DeleteCases(cas string, table []string) []string {
	var result []string
	for i := 0; i < len(table); i++ {
		if table[i] != cas {
			result = append(result, table[i])
		}
	}
	return result
}