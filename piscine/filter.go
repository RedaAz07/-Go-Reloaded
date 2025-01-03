package piscine

import (
	
	"strings"
	
)

func Nextchar(slice []string, index int) (rune, bool) {
	if index < 0 || index >= len(slice)-1 {
		return 0, false
	}

	for i := index + 1; i < len(slice); i++ {
		if slice[i] != "" {
			return rune(slice[i][0]), true
		}
	}

	return 0, false
} 


func SplitePunc(str []string) []string {
	result := []string{}
	currentWord := ""

	for _, word := range str {
		for _, char := range word {
			if !(char == '.' || char == ',' || char == '!' || char == '?' || char == ':' || char == ';') {
				currentWord += string(char)
			} else {

				if currentWord != "" {
					result = append(result, currentWord)
					currentWord = ""
				}

				result = append(result, string(char))
			}
		}

		if currentWord != "" {
			result = append(result, currentWord)
			currentWord = ""
		}
	}

	return result
}
func CheckQuot(word []string, index int) bool {
	if word[index] == "'" {
		if index+1 < len(word) {
			for j := index + 1; j < len(word); j++ {
				if word[j] == "'" {
					return true
				}
			}
		}
	}

	return false
}

func SpliteQuot(str []string) []string {
	result := []string{}
	currentWord := ""

	for _, word := range str {
		for i, char := range word {
			if char == '\'' {
				// Handle consecutive quotes by splitting them
				if i < len(word)-1 && word[i+1] == '\'' {
					if currentWord != "" {
						result = append(result, currentWord)
						currentWord = ""
					}
					// Append each single quote as a separate token
					for i < len(word) && word[i] == '\'' {
						result = append(result, "'")
						i++
					}
					i-- // Adjust for loop increment
				} else if i > 0 && i < len(word)-1 {
					// Quote in the middle of a word
					currentWord += string(char)
				} else {
					// Single quote at start or end of word
					if currentWord != "" {
						result = append(result, currentWord)
						currentWord = ""
					}
					result = append(result, string(char))
				}
			} else {
				// Append normal characters to currentWord
				currentWord += string(char)
			}
		}

		// Append the remaining word after finishing the current word
		if currentWord != "" {
			result = append(result, currentWord)
			currentWord = ""
		}
	}

	return result
}


func Filter(str []string) []string {
	table := SpliteQuot(str)
	inside := false

	for i := 0; i < len(table); i++ {
		correc := true

		switch {

		case table[i] == "'":

			x := 0
			for k := i - 1; k >= 0; k-- {
				if table[k] == "" {
					x++
				} else {
					break
				}
			}

			if i+1 < len(table) && CheckQuot(table, i) && !inside {
				table[i] = "'" + table[i+1]
				if table[i+1] == "'" {
					correc = false
				} else {
					inside = true
				}
				table[i+1] = ""
			} else if i-x > 0 && inside && correc {
				table[i-x-1] += "'"
				table[i] = ""
				inside = false
				break

			}
		case table[i] == "." || table[i] == "," || table[i] == "!" || table[i] == "?" || table[i] == ":" || table[i] == ";":
			if i > 0 {
				for j := 0; j < i; j++ {
					if table[i-1-j] != "" {
						table[i-1-j] += table[i]
						table[i] = ""
					}
				}
			}
		case strings.HasPrefix(table[i], ".") || strings.HasPrefix(table[i], ",") || strings.HasPrefix(table[i], "!") || strings.HasPrefix(table[i], "?") || strings.HasPrefix(table[i], ":") || strings.HasPrefix(table[i], ";"):
			if len(table[i]) > 0 && (table[i][0] == '.' || table[i][0] == ',' || table[i][0] == '!' || table[i][0] == '?' || table[i][0] == ':' || table[i][0] == ';') {
				count := 0
				for j := 0; j < len(table[i]) && (table[i][j] == '.' || table[i][j] == ',' || table[i][j] == '!' || table[i][j] == '?' || table[i][j] == ':' || table[i][j] == ';'); j++ {
					count++
				}

				if i > 0 && table[i-1] != "" {
					table[i-1] += table[i][:count]
					table[i] = table[i][count:]
				}
			}

		case table[i] == "a" || table[i] == "A":
			letter, found := Nextchar(table, i)
			if found && i+1 < len(table) && (letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' || letter == 'h') || (letter == 'A' || letter == 'E' || letter == 'I' || letter == 'O' || letter == 'U' || letter == 'H') {
				table[i] += "n"
			} 
		}
	}
	return table
}
