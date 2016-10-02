package main

import (
	"fmt"
	"os"
)

type token struct {
	kind  string
	value string
}

func tokenizer(input string) []token {
	current := 0
	tokens := []token{}
	kind := ""
	value := ""

	for current < len([]rune(input)) {
		if kind != "" {
			tokens = append(tokens, token{
				kind:  kind,
				value: value,
			})
		}
		kind = ""
		value = ""
		char := string([]rune(input)[current])

		if []rune(input)[current] == 13 {
			current += 2
			continue
		}
		if []rune(input)[current] == 9 {
			current++
			continue
		}
		if char == " " {
			current++
			continue
		}
		if char == "+" || char == "-" {
			kind = "op-suma"
			value = char
			current++
			continue
		}
		if char == "/" {
			value = char
			current++
			if string([]rune(input)[current]) == "*" {
				kind = "comment"
				value += "*"
				current++
				for string([]rune(input)[current]) != "*" && string([]rune(input)[current+1]) != "/" && current < len([]rune(input)) {
					value += string([]rune(input)[current])
					current++
				}
			} else if string([]rune(input)[current]) == "/" {
				kind = "comment-l"
				value += "/"
				for ([]rune(input)[current] != 13 || string([]rune(input)[current]) == "\n") && current < len([]rune(input)) {
					current++
				}
			} else {
				kind = "op-mult"
			}
			continue
		}
		if char == "*" {
			value = char
			current++
			kind = "op-mult"
			continue
		}
		if char == "^" {
			kind = "op-exp"
			value = char
			current++
			continue
		}
		if char == "(" || char == ")" {
			kind = "paren"
			value = char
			current++
			continue
		}
		if char == "<" || char == ">" {
			kind = "comparator"
			value = char
			current++
			if string([]rune(input)[current]) == "=" {
				value += "="
				current++
			}
			continue
		}
		if char == "=" {
			kind = "equals"
			value = char
			current++
			if string([]rune(input)[current]) == "=" {
				value += "="
				kind = "comparator"
				current++
			}
			continue
		}
		if char == "!" {
			value = char
			current++
			if string([]rune(input)[current]) == "=" {
				kind = "comparator"
				value += "="
				current++
			}
			continue
		}
		if char == ";" {
			kind = "p-comma"
			value = ";"
			current++
			continue
		}
		if char == "," {
			kind = "comma"
			value = ","
			current++
			continue
		}
		if char == "{" || char == "}" {
			kind = "llave"
			value = char
			current++
			continue
		}
		if char == "[" || char == "]" {
			kind = "corche"
			value = char
			current++
			continue
		}
		if isNumber(char) {
			dec := true
			for isNumber(char) {
				value += char
				current++
				char = string([]rune(input)[current])
				if char == "." && dec && isNumber(string([]rune(input)[current+1])) {
					dec = false
					value += char
				}
			}
			kind = "numero"
			continue
		}
		if isLetter(char) {
			kind = "identificador"
			for isLetter(char) || isNumber(char) {
				value += char
				current++
				if isNumber(char) {
					kind = "identificador"
				}
				char = string([]rune(input)[current])
			}
			if isReserved(value) {
				kind = "reserved"
			}
			continue
		}

		//if current+1 != len([]rune(input)) {
		fmt.Print("RUTA/PROGRAMA.EXT: syntax error: unexpected ", string([]rune(input)[current]), "\n")
		/*}
		current++
		value = ""
		*/
		os.Exit(1)
		break
	}
	return tokens
}

func isNumber(char string) bool {
	n := []rune(char)[0]
	if n >= '0' && n <= '9' {
		return true
	}
	return false
}

func isLetter(char string) bool {
	n := []rune(char)[0]
	if (n >= 'a' && n <= 'z') || (n >= 'A' && n <= 'Z') {
		return true
	}
	return false
}

func isReserved(word string) bool {
	switch word {
	case "program", "if", "else", "fi", "do", "until", "while", "read", "write": //functions
		return true
	case "float", "int", "bool", "not", "and", "or", "true", "false": //types and op-log
		return true
	}
	return false
}
