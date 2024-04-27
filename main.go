package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	vwl := []string{"a", "o", "u", "e", "A", "O", "U", "E" , "I" , "i"}
	if len(args) != 2 {
		fmt.Println("Error")
		os.Exit(1)
	} else {
		input, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}
		arr := strings.Split(string(input), " ")
		for i, v := range arr {
			if v == "(hex)" && i == 0 {
				fmt.Println("Error")
				os.Exit(1)
			} else if v == "(hex)" && i > 0 {
				arr[i-1] = Hex(arr[i-1])
				arr = append(arr[:i], arr[i+1:]...)
			} else if v == "(bin)" && i == 0 {
				fmt.Println("Err")
				os.Exit(1)
			} else if v == "(bin)" && i > 0 {
				arr[i-1] = Bin(arr[i-1])
				arr = append(arr[:i], arr[i+1:]...)
			} else if v == "(up)" && i == 0 {
				fmt.Println("Err")
				os.Exit(1)
			} else if v == "(up)" && i > 0 {
				arr[i-1] = strings.ToUpper(arr[i-1])
				arr = append(arr[:i], arr[i+1:]...)
			} else if v == "(low)" && i == 0 {
				fmt.Println("Err")
				os.Exit(1)
			} else if v == "(low)" && i > 0 {
				arr[i-1] = strings.ToLower(arr[i-1])
				arr = append(arr[:i], arr[i+1:]...)
			} else if v == "(cap)" && i == 0 {
				fmt.Println("err")
				os.Exit(1)
			} else if v == "(cap)" && i > 0 {
				arr[i-1] = Capital(arr[i-1])
				arr = append(arr[:i], arr[i+1:]...)
			} else if v == "(up," {
				arr[i+1] = strings.TrimRight(arr[i+1], ")")
				number, err := strconv.Atoi(arr[i+1])
				if err != nil {
					fmt.Println("Err")
					os.Exit(1)
				} else {
					if number <= i {
						for j := i - 1; number > 0; j-- {
							number--
							arr[j] = strings.ToUpper(arr[j])
						}
						arr = append(arr[:i], arr[i+2:]...)
					}
				}
			} else if v == "(low," {
				arr[i+1] = strings.TrimRight(arr[i+1], ")")
				number, err := strconv.Atoi(arr[i+1])
				if err != nil {
					fmt.Println("Err")
					os.Exit(1)
				} else {
					if number <= i {
						for j := i - 1; number > 0; j-- {
							number--
							arr[j] = strings.ToLower(arr[j])
						}
						arr = append(arr[:i], arr[i+2:]...)
					}
				}
			} else if v == "(cap," {
				arr[i+1] = strings.TrimRight(arr[i+1], ")")
				number, err := strconv.Atoi(arr[i+1])
				if err != nil {
					fmt.Println("Err")
					os.Exit(1)
				} else {
					if number <= i {
						for j := i - 1; number > 0; j-- {
							number--
							arr[j] = Capital(arr[j])
						}
						arr = append(arr[:i], arr[i+2:]...)
					}
				}
			} else if v == "a" || v == "A" {
				w := arr[i+1]
				if itHas(vwl, string(w[0])) {
					if v == "a" {
						arr[i] = "an"
					} else {
						arr[i] = "An"
					}
				}
			}
		}
		punc(arr)
		str := ""
		for _, v := range arr {
			if v != "" && v != " " {
				str += v + " "
			}
		}
		str = strings.TrimRight(str, " ")
		Write := os.WriteFile(args[1], []byte(str), 0o644)
		if Write != nil {
			fmt.Println("Err")
			os.Exit(1)
		} else {
			fmt.Println("Finish")
			os.Exit(0)
		}
	}
}

func punc(s []string) []string {
	punc := []string{".", ",", "!", "?", ":", ";"}
	q := true
	for i, v := range s {
		isGrouped := true
		for _, l := range v {
			if !itHas(punc, string(l)) {
				isGrouped = false
				break
			}
		}
		if v != "" {
			if itHas(punc, v) && !isGrouped {
				s[i-1] += v
				s[i] = ""
			} else if itHas(punc, string(v[0])) && !isGrouped {
				s[i-1] += string(v[0])
				s[i] = v[1:]
			} else if !isGrouped {
				for j, l := range v {
					if itHas(punc, string(l)) && j != len(v)-1 {
						s[i] = v[:j+1] + " " + v[j+1:]
					}
				}
			}
			if isGrouped {
				s[i-1] += v
				s[i] = ""
			}
			if q && v == "'" {
				if i+1 < len(s) {
					s[i+1] = "'" + s[i+1]
					s[i] = ""
					q = false
				}
			} else if !q && v == "'" {
				if i > 0 {
					if len(s[i-1]) == 0 {
						s[i-2] = s[i-2] + "'"
					} else {
						s[i-1] = s[i-1] + "'"
					}
					s[i] = ""
					q = true
				}
			}
		}
	}
	return s
}

func itHas(Array []string, s string) bool {
	for _, ww := range Array {
		if ww == s {
			return true
		}
	}
	return false
}

func Bin(z string) string {
	dcml, err := strconv.ParseInt(z, 2, 64)
	if err != nil {
		fmt.Println("Err")
		os.Exit(1)
	}
	return fmt.Sprint(dcml)
}

func Hex(z string) string {
	dcml, err := strconv.ParseInt(z, 16, 64)
	if err != nil {
		fmt.Println("Err")
		os.Exit(1)
	}
	return fmt.Sprint(dcml)
}

func Capital(z string) string {
	arr := []rune(z)
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
