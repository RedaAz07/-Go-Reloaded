package piscine

import (
	"strings"
)

func Nextchar(table []string, index int) (rune, bool) {
	if index < 0 || index >= len(table)-1 {
        return 0, false
    }
    for i := index + 1; i < len(table); i++ {
        if table[i] != "" {
            return rune(table[i][0]), true
        }
    }
    return 0, false
 
}


func Avoil( table[]string) []string{
for i := 0; i < len(table); i++ {
	

  if table[i] == "'a" || table[i] == "'A" || table[i] == "a" || table[i] == "A"{

 
	letter, found := Nextchar(table, i)
	if found && i+1 < len(table) && strings.ContainsRune("aeiouhAEIOUH", letter) {
		table[i] += "n"
	} 
}
}
return table 
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
		runes := []rune(word) 
		for i := 0; i < len(runes); i++ {
			char := runes[i]
			if char == '\'' {
				if i < len(runes)-1 && runes[i+1] == '\'' {
					if currentWord != "" {
						result = append(result, currentWord)
						currentWord = ""
					}
					for i < len(runes) && runes[i] == '\'' {
						result = append(result, "'")
						i++
					}
					i-- // Adjust index after inner loop
				} else if i > 0 && i < len(runes)-1 {
					currentWord += string(char)
				} else {
					if currentWord != "" {
						result = append(result, currentWord)
						currentWord = ""
					}
					result = append(result, string(char))
				}
			} else {
				currentWord += string(char)
			}
		}
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

		}
	}


	table = Avoil(table)
  	return table
}
