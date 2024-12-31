package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"piscine/piscine"
)

// add  the result to the  file
func AddToFile(data []byte, name string) {
	err := os.WriteFile(name, data, 0o644)
	if err != nil {
		fmt.Println(" Error ")
	}
	fmt.Println("success")
}

// Read the file
func ReadFile(name string) string {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("There is no file with this name")
	}
	return string(data)
}

// Process all functrions
func ProcessAll() {
	table := addToTable()
	table = Proccedflag(table)

	table = withNumber(table)
	table = BinaryToDecimal(table)
	table = HexadecimalToDecimal(table)
	// table = Capitalized(table)
	table = Lower(table)
	table = Upper(table)
	table = punctuations(table)
	table = Avowel(table)
	table = Filter(table)
	text := piscine.TableToString(table)
	data := []byte(text)
	AddToFile(data, os.Args[2])
}

func Proccedflag(table []string) []string {
	for i := 0; i < len(table); i++ {
		switch {
		case table[i] == "(cap)":
			table[i-1] = Capitalized(table[i-1])
			table[i-1]=""
		}
	}

	return table
}

func HexadecimalToDecimal(table []string) []string {
	for i := 0; i < len(table); i++ {
		if i > 0 {
			if table[i] == "(hex)" {

				counter := 0

				counter = piscine.Ignoreempty(i, counter, table)

				decimalNumber, err := strconv.ParseInt(table[i-1-counter], 16, 0)
				if err != nil {
					fmt.Println("Error:", err)
					return table
				}
				table[i-1-counter] = fmt.Sprintf("%d", decimalNumber) // Convert to string and replace
			}
		}
	}
	result := piscine.DeleteCases("(hex)", table)

	return result
}

func BinaryToDecimal(table []string) []string {
	for i := 0; i < len(table); i++ {
		if i > 0 {
			if table[i] == "(bin)" {

				counter := 0

				counter = piscine.Ignoreempty(i, counter, table)

				decimalNumber, err := strconv.ParseInt(table[i-1-counter], 2, 0)
				if err != nil {
					fmt.Println("Error:", err)
					return table
				}
				table[i-1-counter] = fmt.Sprintf("%d", decimalNumber)
			}
		}
	}
	result := piscine.DeleteCases("(bin)", table)
	return result
}

func Capitalized(table string) string {
	str:=""
	pass:=true
	for _, char := range table {
		
		if unicode.IsLetter(char) && pass{
			str+=strings.ToUpper(string(char))
			pass=false
		}else{
			str+=strings.ToLower(string(char))
		}
	}
	return str
}

func Upper(table []string) []string {
	for i := 0; i < len(table); i++ {
		if i > 0 {
			if table[i] == "(up)" {
				counter := 0
				counter = piscine.Ignoreempty(i, counter, table)
				table[i-1] = strings.ToUpper(table[i-1-counter])
			}
		}
	}
	result := piscine.DeleteCases("(up)", table)
	return result
}

func Lower(table []string) []string {
	for i := 0; i < len(table); i++ {
		if i > 0 {
			if table[i] == "(low)" {
				counter := 0
				counter = piscine.Ignoreempty(i, counter, table)
				table[i-1-counter] = strings.ToLower(table[i-1-counter])
			}
		}
	}

	result := piscine.DeleteCases("(low)", table)

	return result
}

func withNumber(table []string) []string {
	fmt.Println(table)

	count := 0
	for i := 0; i < len(table); i++ {
		if i > 0 {
			count++
			num1 := 0
			if string(table[i]) == "(low," && strings.HasSuffix(table[i+1], ")") {

				num1 = piscine.MyAtoi(table[i+1])
				if count < num1 {
					num1 = count
				}
				counter := 0
				counter = piscine.Ignoreempty(i, counter, table)
				fmt.Println(counter, "low counter")
				for j := 0; j <= num1; j++ {

					table[i-j-counter] = strings.ToLower(table[i-j-counter])
					table[i] = ""
					if i < len(table) {
						table[i] = ""
					}
					if i+1 < len(table) {
						table[i+1] = ""
					}

				}
			}
			if string(table[i]) == "(up," && strings.HasSuffix(table[i+1], ")") {
				num1 := 0
				num1 = piscine.MyAtoi(string(table[i+1]))
				if count < num1 {
					num1 = count
				}
				counter := 0
				counter = piscine.Ignoreempty(i, counter, table)

				for j := 0; j <= num1; j++ {

					table[i-j-counter] = strings.ToUpper(table[i-j-counter])
					table[i] = ""
					if i < len(table) {
						table[i] = ""
					}
					if i+1 < len(table) {
						table[i+1] = ""
					}
				}
			}
			if string(table[i]) == "(cap," && strings.HasSuffix(table[i+1], ")") {
				num1 = piscine.MyAtoi(string(table[i+1]))

				if count < num1 {
					num1 = count
				}

				counter := 0
				counter = piscine.Ignoreempty(i, counter, table)

				for j := 0; j <= num1; j++ {

					word := table[i-j-counter]
					if table[i-j-counter] != "" {

						runes := []rune(word)
						runes[0] = unicode.ToUpper(runes[0])
						for i := 1; i < len(runes); i++ {
							runes[i] = unicode.ToLower(runes[i])
						}
						table[i-j-counter] = string(runes)
						table[i] = ""
						if i < len(table) {
							table[i] = ""
						}
						if i+1 < len(table) {
							table[i+1] = ""
						}
					}
				}
			}
		} /* else if i < len(table)-1{
			table[i] = ""
			table[i+1] = ""

		} */
	}
	fmt.Println(table)
	var result []string
	for i := 0; i < len(table); i++ {
		if table[i] != "" {
			result = append(result, table[i])
		}
	}
	return result
}

func punctuations(table []string) []string {
	var result []string
	// i do this to handle the  case  of : ... .... ... ... . .
	for i := len(table) - 1; i >= 0; i-- {
		if i > 0 && strings.HasPrefix(table[i], ".") || strings.HasPrefix(table[i], ",") || strings.HasPrefix(table[i], "!") || strings.HasPrefix(table[i], "?") || strings.HasPrefix(table[i], ":") || strings.HasPrefix(table[i], ";") {
			count := 0
			for j := 0; j < len(table[i]); j++ {
				if table[i][j] == ',' || table[i][j] == '.' || table[i][j] == '!' || table[i][j] == '?' || table[i][j] == ':' || table[i][j] == ';' {
					count++
				}
			}
			if count == 1 {
				table[i-1] = table[i-1] + string(table[i][0])
				table[i] = table[i][1:]

			} else if i > 0 {
				table[i-1] = table[i-1] + string(table[i][0:count])
				table[i] = table[i][count:]
			}
		}
	}
	for i := 0; i < len(table); i++ {
		if table[i] == "." || table[i] == "," || table[i] == "!" || table[i] == "?" || table[i] == ":" || table[i] == ";" {
			if i > 0 {
				for j := 0; j < i; j++ {
					if table[i-1-j] != "" {
						table[i-1-j] += table[i]
						table[i] = ""
					}
				}
			}
		} else if strings.HasPrefix(table[i], ".") || strings.HasPrefix(table[i], ",") || strings.HasPrefix(table[i], "!") || strings.HasPrefix(table[i], "?") || strings.HasPrefix(table[i], ":") || strings.HasPrefix(table[i], ";") {
			if len(table[i]) > 0 && (table[i][0] == '.' || table[i][0] == ',' || table[i][0] == '!' || table[i][0] == '?' || table[i][0] == ':' || table[i][0] == ';') {
				count := 0
				for j := 0; j < len(table[i]) && (table[i][j] == '.' || table[i][j] == ',' || table[i][j] == '!' || table[i][j] == '?' || table[i][j] == ':' || table[i][j] == ';'); j++ {
					count++
				}
				fmt.Printf("%q", table)
				if i > 0 && table[i-1] != "" {

					table[i-1] += table[i][:count]
					table[i] = table[i][count:]
				} else if i == 0 {

					count := 0

					for j := 0; j < len(table[i]) && (table[i][j] == '.' || table[i][j] == ',' || table[i][j] == '!' || table[i][j] == '?' || table[i][j] == ':' || table[i][j] == ';'); j++ {
						count++
					}

					table = append([]string{table[i][:count], table[i][count:]}, table[i+1:]...)

				}

			}
		} else if strings.ContainsAny(table[i], ".,!?;:") {
			fmt.Println("scond ")

			index := strings.IndexAny(table[i], ".,!?;:")
			if index > 0 && index < len(table[i])-1 {
				punctuation := table[i][index:]
				remaining := table[i][:index]

				table[i] = remaining
				table = append(table[:i+1], append([]string{punctuation}, table[i+1:]...)...)
			}
		}
	}

	for i := 0; i < len(table); i++ {
		if table[i] != "" {
			result = append(result, table[i])
		}
	}
	return result
}

func ChecKer(text []string, index int) bool {
	if text[index] == "'" {
		if index+1 < len(text) {
			for j := index + 1; j < len(text); j++ {
				if text[j] == "'" {
					return true
				}
			}
		}
	}

	return false
}

func CheckNextLetter(slice []string, index int) (rune, bool) {
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

func filltquot(str []string) []string {
	result := []string{}
	currentWord := ""

	for _, word := range str {
		for i := 0; i < len(word); i++ {
			char := word[i]

			if char == '\'' {
				if i > 0 && i < len(word)-1 {
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
	str = filltquot(str)
	isinside := false

	for i := 0; i < len(str); i++ {
		correct := true

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

			if i+1 < len(str) && ChecKer(str, i) && !isinside {
				str[i] = "'" + str[i+1]
				if str[i+1] == "'" {
					correct = false
				} else {
					isinside = true
				}
				str[i+1] = ""
			} else if i-x > 0 && isinside && correct {
				str[i-x-1] += "'"
				str[i] = ""
				isinside = false
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
			letter, found := CheckNextLetter(str, i)
			if found && i+1 < len(str) && (letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' || letter == 'h') || (letter == 'A' || letter == 'E' || letter == 'I' || letter == 'O' || letter == 'U' || letter == 'H') {
				str[i] += "n"
			}
		}
	}
	return str
}

func Avowel(table []string) []string {
	for i := 0; i < len(table); i++ {
		if i < len(table)-1 {
			if (table[i] == "a" || table[i] == "A") && (strings.HasPrefix(table[i+1], "a") || strings.HasPrefix(table[i+1], "e") || strings.HasPrefix(table[i+1], "i") || strings.HasPrefix(table[i+1], "o") || strings.HasPrefix(table[i+1], "u") || strings.HasPrefix(table[i+1], "h") || strings.HasPrefix(table[i+1], "A") || strings.HasPrefix(table[i+1], "E") || strings.HasPrefix(table[i+1], "I") || strings.HasPrefix(table[i+1], "O") || strings.HasPrefix(table[i+1], "U") || strings.HasPrefix(table[i+1], "H")) {
				table[i] += "n"
			}
		}
	}
	return table
}

func addToTable() []string {
	name := os.Args[1]
	txt := ReadFile(name)

	// Split the content of the file into words
	word := ""
	var table []string
	for _, i := range txt {
		if i != ' ' {
			word += string(i)
		} else {
			if word != "" {
				table = append(table, word)
			}
			word = ""
		}
	}
	if word != "" {
		table = append(table, word)
	}
	return table
}

func main() {
	// fmt.Println(MyAtoi("10)"))

	if len(os.Args) != 3 {
		fmt.Println("bro you should to write just the sample, result file names   ")
	} else {
		if !strings.HasSuffix(os.Args[2], ".txt") {
			fmt.Println(" Nice try bro (^_-) ")
		} else {
			ProcessAll()
		}
	}
}
