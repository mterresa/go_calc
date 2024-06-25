package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var romanMap = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var arabMap = map[int]string{
	100: "C",
	90:  "XC",
	50:  "L",
	40:  "XL",
	10:  "X",
	9:   "IX",
	5:   "V",
	4:   "IV",
	1:   "I",
}

var mySort = [9]int{
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// var text string = "X * X"
	fmt.Println(parser(text))
}

func parser(str string) string {
	operator := ""
	text := strings.Split(strings.TrimSpace(strings.ToUpper(str)), "")
	for _, val := range text {
		if val == "+" || val == "-" || val == "*" || val == "/" {
			operator += string(val)
			break
		}
	}
	var args_f []string = strings.Split(strings.TrimSpace(strings.ToUpper(str)), operator)
	if len(args_f) > 2 {
		return "Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	}
	if len(args_f) <= 1 || !is_operation(operator) {
		return "Выдача паники, так как строка не является математической операцией."
	}
	var args [3]string
	args[0] = strings.TrimSpace(args_f[0])
	args[2] = strings.TrimSpace(args_f[1])
	args[1] = operator
	args[2] = strings.TrimSpace(args[2])
	var arg1, err1 = strconv.Atoi(args[0])
	var arg2, err2 = strconv.Atoi(args[2])
	if err1 == nil && err2 == nil {
		if 0 < arg1 && arg1 < 11 && 0 < arg2 && arg2 < 11 {
			return strconv.Itoa(calculation(arg1, arg2, args[1]))
		}
		return "Выдача паники, некорректное число."
	}
	if err1 != nil && err2 != nil {
		r_arg1 := romanToInt(args[0])
		r_arg2 := romanToInt(args[2])
		if 0 < r_arg1 && r_arg1 < 11 && 0 < r_arg2 && r_arg2 < 11 {
			return intToRoman(calculation(r_arg1, r_arg2, args[1]))
		}
		return "Выдача паники, некорректное число."
	}
	return "Выдача паники, так как используются одновременно разные системы счисления."
}

func intToRoman(arg int) string {
	if arg < 1 {
		return "Выдача паники, так как в римской системе нет отрицательных чисел."
	}
	var ans string = ""
	n := arg
	for n > 0 {
		for _, k := range mySort {
			if k <= n {
				ans += arabMap[k]
				n -= k
				break
			}
		}
	}
	return ans
}

func romanToInt(arg string) int {
	_, isPresent := romanMap[arg]
	if isPresent {
		return romanMap[arg]
	}
	return 11
}

func is_operation(oper string) bool {
	operands := []string{"+", "-", "*", "/"}
	return slices.Contains(operands, oper)
}

func calculation(arg1 int, arg2 int, oper string) int {
	switch oper {
	case "+":
		return arg1 + arg2
	case "-":
		return arg1 - arg2
	case "*":
		return arg1 * arg2
	default:
		return arg1 / arg2
	}
}

// func calculation(arg1 int, arg2 int, oper string) string {
// 	switch oper {
// 	case "+":
// 		return strconv.Itoa(arg1 + arg2)
// 	case "-":
// 		return strconv.Itoa(arg1 - arg2)
// 	case "*":
// 		return strconv.Itoa(arg1 * arg2)
// 	default:
// 		return strconv.Itoa(arg1 / arg2)
// 	}
// }
