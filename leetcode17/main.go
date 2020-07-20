package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("27"))
}

var (
	phone = map[byte][]byte{
		'2':[]byte{'a','b','c'},
		'3':[]byte{'d', 'e', 'f'},
		'4':[]byte{'g', 'h', 'i'},
		'5':[]byte{'j', 'k', 'l'},
		'6':[]byte{'m', 'n', 'o'},
		'7':[]byte{'p', 'q', 'r', 's'},
		'8':[]byte{'t', 'u', 'v'},
		'9':[]byte{'w', 'x', 'y', 'z'},
	}
	output = make([]string, 0)
	)

func letterCombinations(digits string) []string {

	if len(digits) !=0 {
		backtrack("", digits)
	}
	return output
}

func backtrack(combination string, digits string) {
	if len(digits) == 0{
		output = append(output, combination)
	} else {
		num:=digits[0]
		letters := phone[num]
		for _, letter:=range letters {
			backtrack(combination + string(letter), digits[1:])
		}
	}
}