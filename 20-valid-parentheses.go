package main

import (
	"fmt"
)

func isClosing(b byte) bool {
	CLOSING := []byte{'}', ']', ')'}
	for i:=0; i<len(CLOSING); i++ {
		if b == CLOSING[i] {
			return true
		}
	}
	return false
}

func isClosingFor(closing byte, open byte) bool {
	if closing == ')' && open == '(' {return true}
	if closing == '}' && open == '{' {return true}
	if closing == ']' && open == '[' {return true}
	return false
}

func isValid(s string) bool {
	const STACK_SIZE = 10000
	stack := make([]byte, STACK_SIZE)
	stack_ptr := 0
	for _, sym := range s {
		if !isClosing(byte(sym)) {
			stack[stack_ptr] = byte(sym)
			stack_ptr++
		} else {
			stack_ptr--
			if stack_ptr < 0 {return false}
			if !isClosingFor(byte(sym), stack[stack_ptr]) {return false}
		}
	}
	if stack_ptr > 0 {
		return false
	}
	return true
}

func main() {
	test_cases := []string{
		"()",
		"()[]{}",
		"(]",
	}
	for _, test_case := range test_cases {
		fmt.Println(isValid(test_case))
	}

}
