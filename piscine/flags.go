package piscine

import (
	"fmt"
	
	"strconv"
	"strings"
	"unicode"

)


func hex(word string) string {
	correctNumber, err := strconv.ParseInt(word, 16, 0)
	if err != nil {
		return word
	}
	return fmt.Sprintf("%d", correctNumber)
}

func bin(word string) string {
	correctNumber, err := strconv.ParseInt(word, 2, 0)
	if err != nil {
		return word
	}
	return fmt.Sprintf("%d", correctNumber)
}

func up(world string) string {
	return strings.ToUpper(world)
}

func low(world string) string {
	return strings.ToLower(world)
}

func cap(world string) string {
	x := ""
	firstRune := true
	for _, chr := range world {
		if unicode.IsLetter(chr) && firstRune {
			x += strings.ToUpper(string(chr))
			firstRune = false
		} else {
			x += strings.ToLower(string(chr))
		}
	}
	return x
}

func Iscorrect(nbr string) (int, error) {
	if len(nbr) > 0 && nbr[len(nbr)-1] == ')' {
		return strconv.Atoi(nbr[:len(nbr)-1])
	}
	return strconv.Atoi(nbr)
}


func Flags (str []string) []string {
	for i := 0; i < len(str); i++ {
		switch {
		case str[i] == "(hex)":
			for j := i - 1; j >= 0; j-- {
				if str[j] != "" {
					str[j] = hex(str[j])
					break
				}
			}
			str[i] = ""

		case str[i] == "(bin)":
			for j := i - 1; j >= 0; j-- {
				if str[j] != "" {
					str[j] = bin(str[j])
					break
				}
			}
			str[i] = ""
		case str[i] == "(up)":
			if i > 0 {
				for j := i - 1; j >= 0; j-- {
					if str[j] != "" {
						str[j] = up(str[j])
						str[i] = ""
						break
					}
				}
			} else {
				str[i] = ""
			}

		case str[i] == "(low)":
			if i > 0 {
				for j := i - 1; j >= 0; j-- {
					if str[j] != "" {
						str[j] = low(str[j])
						str[i] = ""
						break
					}
				}
			} else {
				str[i] = ""
			}
		case str[i] == "(cap)":
			if i > 0 {
				for j := i - 1; j >= 0; j-- {
					if str[j] != "" {
						str[j] = cap(str[j])
						str[i] = ""
						break
					}
				}
			} else {
				str[i] = ""
			}
		case str[i] == "(up,":

			if i+1 < len(str) {
				countStr := str[i+1]
				count, err := Iscorrect(countStr)
				var emptyctr int
				if err != nil {
					fmt.Println("Error: invalid count or not enough words to transform")
				} else if str[i+1][len(str[i+1])-1] == ')' {
					for k := i - 1; k >= 0; k-- {
						if str[k] == "" {
							emptyctr++
						} else {
							break
						}
					}
					if count+emptyctr >= i {
						count = i - emptyctr
					}
					end := i - count - emptyctr
					for j := i - 1; j >= end; j-- {
						if str[j] != "" {
							str[j] = up(str[j])
						}
					}
					str[i] = ""
					str[i+1] = ""
				}

			}

		case str[i] == "(low,":
			if i < len(str) {
				countStr := str[i+1]
				count, err := Iscorrect(countStr)
				var emptyctr int
				if err != nil {
					fmt.Println("Error")
				} else if str[i+1][len(str[i+1])-1] == ')' {
					for k := i - 1; k >= 0; k-- {
						if str[k] == "" {
							emptyctr++
						} else {
							break
						}
					}
					if count+emptyctr >= i {
						count = i - emptyctr
					}

					for j := 0; j < count+emptyctr; j++ {
						if str[i-1-j] != "" {
							str[i-1-j] = low(str[i-1-j])
						}
					}
					str[i] = ""
					str[i+1] = ""
				}

			}
		case str[i] == "(cap,":

			if i < len(str) {

				countStr := str[i+1]
				count, err := Iscorrect(countStr)
				var emptyctr int
				if err != nil {
					fmt.Println("Error")
				} else if str[i+1][len(str[i+1])-1] == ')' {
					for k := i - 1; k >= 0; k-- {
						if str[k] == "" {
							emptyctr++
						} else {
							break
						}
					}
					if count+emptyctr >= i {
						count = i - emptyctr
					}
					for j := 0; j < count+emptyctr; j++ {
						if str[i-1-j] != "" {
							str[i-1-j] = cap(str[i-1-j])
						}
					}
					str[i] = ""
					str[i+1] = ""
				}
			}

		}
	}
	return str

}





