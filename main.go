package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 2 {
		Text, _ := os.ReadFile(args[0])

		array := strings.Split(string(Text), " ")
		for i, chr := range array {
			if chr == "(up)" {
				if i == 0 {
					fmt.Println("Error")
					os.Exit(1)
				} else {
					array[i-1] = strings.ToUpper(array[i-1])
					array = append(array[:i], array[i+1:]...)
				}
			} else if chr == "(low)" {
				if i == 0 {
					fmt.Println("Error")
					os.Exit(1)
				} else {
					array[i-1] = strings.ToLower(array[i-1])
					array = append(array[:i], array[i+1:]...)
				}
			} else if chr == "(cap)" {
				if i == 0 {
					fmt.Println("Error")
					os.Exit(1)
				} else {
					array[i-1] = capitalize(array[i-1])
					array = append(array[:i], array[i+1:]...)
					fmt.Println(array)
				}
			} else if chr == "(hex)" {
				if i == 0 {
					fmt.Println("Error")
					os.Exit(1)
				} else {
					array[i-1] = hex(array[i-1])
					array = append(array[:i], array[i+1:]...)
				}
			} else if chr == "(bin)" {
				array[i-1] = BintoInt(array[i-1])
				array = append(array[:i], array[i+1:]...)

			} else if chr == "(up," {
				if i == 0 {
					fmt.Println("Error")
					os.Exit(1)
				} else {

					b := strings.TrimRight(string(chr[i+1]), ")")
					number, _ := strconv.Atoi(string(b))
					for j := 1; j <= number; j++ {
						array[i-j] = strings.ToUpper(array[i-j])
					}
					array = append(array[:i], array[i+2:]...)

				}
			} else if chr == "(low," {
				if i == 0 {
					fmt.Println("Error")
					os.Exit(1)
				} else {
					b := strings.TrimRight(string(array[i+1]), ")")
					number, _ := strconv.Atoi(string(b))
					for j := 1; j <= number; j++ {
						array[i-j] = strings.ToLower(array[i-j])
					}
					array = append(array[:i], array[i+2:]...)

				}
			} else if chr == "(cap," {
				b := strings.TrimRight(string(chr[i+1]), ")")
				number, _ := strconv.Atoi(string(b))
				for j := 1; j <= number; j++ {
					array[i-j] = capitalize(array[i-j])
				}
				array = append(array[:i], array[i+2:]...)
			}
		}

		punc(array)
		str := ""

		for _, v := range array {
			if v != "" && v != " " {
				str += v + " "
			}
		}
		str = strings.TrimRight(str, " ")
		Write := os.WriteFile(args[1], []byte(str), 0644)
		if Write != nil {
			fmt.Println("Error In Output")
		} else {
			fmt.Println("Done")
			os.Exit(0)
		}

	} else {
		fmt.Println("Missing file")
	}
}

func hex(hex string) string {
	number, _ := strconv.ParseInt(hex, 16, 64)
	return fmt.Sprint(number)
}

func BintoInt(bin string) string {
	number, _ := strconv.ParseInt(bin, 2, 64)
	return fmt.Sprint(number)
}
func capitalize(chr string) string {
	s := []rune(chr)
	for i, v := range s {
		if i == 0 && (v >= 'a' && v <= 'z') {
			s[i] -= 32
		} else if i > 0 && (v >= 'A' && v <= 'Z') {
			s[i] += 32
		}
	}
	return string(s)
}

func punc(s []string) []string {
	punctuations := []string{".", ",", "!", "?", ":", ";"}
	firstQ := true
	for i, v := range s {
		isGroup := true
		for _, l := range v {
			if !Contain(punctuations, string(l)) {
				isGroup = false
				break
			}
		}
		if v != "" {
			if Contain(punctuations, v) && !isGroup {
				s[i-1] += v
				s[i] = ""
				//fmt.Print(s[i])
			} else if Contain(punctuations, string(v[0])) && !isGroup {
				s[i-1] += string(v[0])
				s[i] = v[1:]
			} else if !isGroup {
				for j, l := range v {
					if Contain(punctuations, string(l)) && j != len(v)-1 {
						s[i] = v[:j+1] + " " + v[j+1:]
					}
				}
			}
			if isGroup {
				s[i-1] += v
				s[i] = ""
			}
			if firstQ && v == "'" {
				if i+1 < len(s) {
					s[i+1] = "'" + s[i+1]
					s[i] = ""
					firstQ = false
				}
			} else if !firstQ && v == "'" {
				if i > 0 {
					if len(s[i-1]) == 0 {
						s[i-2] = s[i-2] + "'"
					} else {
						s[i-1] = s[i-1] + "'"
					}
					s[i] = ""
					firstQ = true
				}
			}
		}
	}
	return s
}

func Contain(arr []string, a string) bool {
	for _, v := range arr {
		if v == a {
			return true
		}
	}
	return false
}

func Cap(hex string) string {
	arr := []rune(hex)
	for i, v := range arr {
		if i == 0 && (v >= 97 && v <= 122) {
			arr[i] -= 32
		}
		if i > 0 && (v >= 65 && v <= 90) {
			arr[i] += 32
		}
	}
	return string(arr)
}

func HexToDec(hex string) string {
	decimal, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		ErrorExit("Can't Convert Hex to Decimal")
	}
	return fmt.Sprint(decimal)

}
func BinToDec(hex string) string {
	decimal, err := strconv.ParseInt(hex, 2, 64)
	if err != nil {
		ErrorExit("Can't Convert Bin to Decimal")
	}
	return fmt.Sprint(decimal)
}

func ErrorExit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
