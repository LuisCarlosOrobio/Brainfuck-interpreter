package main

import (
	"fmt"
	"os"
)

func brainfuckInterpreter(code, input string) string {
	data := make([]byte, 30000) // Memory tape initialized to zero
	var output string
	var codePtr, dataPtr, inputPtr int
	loopStack := []int{}

	for codePtr < len(code) {
		switch code[codePtr] {
		case '>':
			dataPtr++
		case '<':
			dataPtr--
		case '+':
			data[dataPtr]++
		case '-':
			data[dataPtr]--
		case '.':
			output += string(data[dataPtr])
		case ',':
			if inputPtr < len(input) {
				data[dataPtr] = input[inputPtr]
				inputPtr++
			}
		case '[':
			if data[dataPtr] == 0 {
				for openBrackets := 1; openBrackets > 0; codePtr++ {
					if code[codePtr+1] == '[' {
						openBrackets++
					} else if code[codePtr+1] == ']' {
						openBrackets--
					}
				}
			} else {
				loopStack = append(loopStack, codePtr)
			}
		case ']':
			if data[dataPtr] != 0 {
				codePtr = loopStack[len(loopStack)-1]
			} else {
				loopStack = loopStack[:len(loopStack)-1]
			}
		}
		codePtr++
	}

	return output
}

func main() {
	// Example Brainfuck code that outputs 'A'
	bfCode := "++++++++[>++++++++<-]>+."
	input := "" // No input in this example
	fmt.Println(brainfuckInterpreter(bfCode, input)) // Should print 'A'

	// Run with custom Brainfuck code from command line argument
	if len(os.Args) > 1 {
		fmt.Println(brainfuckInterpreter(os.Args[1], input))
	}
}
