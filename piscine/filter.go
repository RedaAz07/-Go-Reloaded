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
        for i := 0; i < len(word); i++ {
            char := word[i]

            if char == '\'' {
                if i < len(word)-1 && word[i+1] == '\'' {
                    if currentWord != "" {
                        result = append(result, currentWord)
                        currentWord = ""
                    }
                    for i < len(word) && word[i] == '\'' {
                        result = append(result, "'")
                        i++
                    }
                    i-- 
                } else if i > 0 && i < len(word)-1 {
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
	str = SpliteQuot(str)
	inside := false

	for i := 0; i < len(str); i++ {
		correc := true

		switch {
		case str[i] == "'":
			x := 0
			for k := i - 1; k >= 0; k-- {
				if str[k] == "" {
					x++
				} else {
					break
				}
			}

			if i+1 < len(str) && CheckQuot(str, i) && !inside {
				str[i] = "'" + str[i+1]
				if str[i+1] == "'" {
					correc = false
				} else {
					inside = true
				}
				str[i+1] = ""
			} else if i-x > 0 && inside && correc {
				str[i-x-1] += "'"
				str[i] = ""
				inside = false
				break

			}
		case str[i] == "." || str[i] == "," || str[i] == "!" || str[i] == "?" || str[i] == ":" || str[i] == ";":
			if i > 0 {
				for j := 0; j < i; j++ {
					if str[i-1-j] != "" {
						str[i-1-j] += str[i]
						str[i] = ""
					}
				}
			}
		case strings.HasPrefix(str[i], ".") || strings.HasPrefix(str[i], ",") || strings.HasPrefix(str[i], "!") || strings.HasPrefix(str[i], "?") || strings.HasPrefix(str[i], ":") || strings.HasPrefix(str[i], ";"):
			if len(str[i]) > 0 && (str[i][0] == '.' || str[i][0] == ',' || str[i][0] == '!' || str[i][0] == '?' || str[i][0] == ':' || str[i][0] == ';') {
				count := 0
				for j := 0; j < len(str[i]) && (str[i][j] == '.' || str[i][j] == ',' || str[i][j] == '!' || str[i][j] == '?' || str[i][j] == ':' || str[i][j] == ';'); j++ {
					count++
				}

				if i > 0 && str[i-1] != "" {
					str[i-1] += str[i][:count]
					str[i] = str[i][count:]
				}
			}

		case str[i] == "a" || str[i] == "A":
			letter, found := Nextchar(str, i)
			if found && i+1 < len(str) && (letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' || letter == 'h') || (letter == 'A' || letter == 'E' || letter == 'I' || letter == 'O' || letter == 'U' || letter == 'H') {
				str[i] += "n"
			} 
		}
	}
	return str
}
