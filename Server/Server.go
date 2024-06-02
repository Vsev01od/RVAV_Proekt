package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//1

		var data map[string]interface{}

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		v, ok := data["message"].(string)
		if !ok {
			http.Error(w, "invalid data type", http.StatusBadRequest)
			return
		}
		str := v
		fmt.Println(str)

		//2

		i := infixToPostfix(str)
		result := evaluatePostfix(i)
		fmt.Println(result)

	})
	http.ListenAndServe("localhost:8080", nil)
}

func precedence(operator rune) int {
	if operator == '*' || operator == '/' {
		return 2
	} else if operator == '+' || operator == '-' {
		return 1
	}
	return 0
}

func infixToPostfix(expression string) string {
	var result string
	var stack []rune
	for _, char := range expression {
		switch {
		case char >= '0' && char <= '9':
			result += string(char) + " "
		case char == '(':
			stack = append(stack, char)
		case char == ')':
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				result += string(stack[len(stack)-1]) + " "
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		default:
			for len(stack) > 0 && precedence(stack[len(stack)-1]) >= precedence(char) {
				result += string(stack[len(stack)-1]) + " "
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, char)
		}
	}
	for len(stack) > 0 {
		result += string(stack[len(stack)-1]) + " "
		stack = stack[:len(stack)-1]
	}
	return result
}
func evaluatePostfix(postfix string) int {
	var stack []int
	for _, char := range postfix {
		if char >= '0' && char <= '9' {
			digit, _ := strconv.Atoi(string(char))
			stack = append(stack, digit)
		} else if char == '+' {
			op1 := stack[len(stack)-2]
			op2 := stack[len(stack)-1]
			stack[len(stack)-2] = op1 + op2
			stack = stack[:len(stack)-1]
		} else if char == '-' {
			op1 := stack[len(stack)-2]
			op2 := stack[len(stack)-1]
			stack[len(stack)-2] = op1 - op2
			stack = stack[:len(stack)-1]
		} else if char == '*' {
			op1 := stack[len(stack)-2]
			op2 := stack[len(stack)-1]
			stack[len(stack)-2] = op1 * op2
			stack = stack[:len(stack)-1]
		}
	}
	return stack[0]
}
