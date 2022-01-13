package main

import (
	"strconv"
)

// Type Declaration
type token struct {
	tokenType string // The type of the token or CATEGORY
	tokenValue  string // The value of the token or VALUE
}

// Error Functions
func printParseError(line string, lineNumber int, token string) {
	functionMessage := ""
	functionMessage += "\n"
	functionMessage += "💀 Unexpected String Literal\n"
	functionMessage += `Unexpected String Literal '` + token + `' at line number ` + strconv.Itoa(lineNumber) + "\n"
	functionMessage += "👉 " + line + "\n"
	functionMessage += "\n"
	panicMessage += functionMessage
	panic("💀 ERROR 💀")
}

func printExpectedTokenError(line string, lineNumber int, expectedToken string) {
	functionMessage := ""
	functionMessage += "\n"
	functionMessage += "💀 Token was Expected\n"
	functionMessage += `Expected Token '` + expectedToken + `' at line number ` + strconv.Itoa(lineNumber) + "\n"
	functionMessage += "👉 " + line + "\n"
	functionMessage += "\n"
	panicMessage += functionMessage
	panic("💀 ERROR 💀")
}

func printError(line string, lineNumber int, i int) {
	functionMessage := ""
	functionMessage += "\n"
	functionMessage += "💀 Tokenization Error\n"
	functionMessage += `Unknown token: "` + string(line[i]) + `" at line number ` + strconv.Itoa(lineNumber) + "\n"
	functionMessage += line + "\n"
	for j := 0; j < i; j++ {
		functionMessage += " "
	}
	functionMessage += "👆\n"
	panicMessage += functionMessage
	panic("💀 ERROR 💀")
}

func runtimeError(line string, lineNumber int, errorMessage string) {
	functionMessage := ""
	functionMessage += "\n"
	functionMessage += "💀 Runtime Error\n"
	functionMessage += errorMessage + ` at line ` + strconv.Itoa(lineNumber) + "\n"
	functionMessage += "👉 " + line + "\n"
	functionMessage += "\n"
	panicMessage += functionMessage
	panic("💀 ERROR 💀")
}

// Lexer
func lex(line string, lineNumber int) []token {

	var tokens []token
	numeral := ""
	var i int

	for i = 0; i < len(line); i++ {
		if line[i] == '#' {
			break
		} else  if line[i] == ' ' {
			if numeral != "" {
				tokens = append(tokens, token{"NUMBER", numeral})
				numeral = ""
			}
			continue
		} else if line[i] == '+' {
			tokens = append(tokens, token{"ARITHMETIC", "+"}) // +
		} else if line[i] == '-' {
			tokens = append(tokens, token{"ARITHMETIC", "-"}) // -
		} else if line[i] == '*' {
			tokens = append(tokens, token{"ARITHMETIC", "*"}) // *
		} else if line[i] == '/' {
			tokens = append(tokens, token{"ARITHMETIC", "/"}) // /
		} else if line[i] == '=' {
			if line[i+1] == '=' {
				tokens = append(tokens, token{"COMPARISON", "=="}) // ==
				i++
			} else {
				tokens = append(tokens, token{"ASSIGNMENT", "="}) // =
			}
		} else if line[i] == '<' {
			if line[i+1] == '=' {
				tokens = append(tokens, token{"COMPARISON", "<="}) // <=
				i++
			} else {
				tokens = append(tokens, token{"COMPARISON", "<"}) // <
			}
		} else if line[i] == '>' {
			if line[i+1] == '=' {
				tokens = append(tokens, token{"COMPARISON", ">="}) // >=
				i++
			} else {
				tokens = append(tokens, token{"COMPARISON", ">"}) // >
			}
		} else if line[i] == '!' {
			if line[i+1] == '=' {
				tokens = append(tokens, token{"COMPARISON", "!="}) // !=
				i++
			}
		} else if line[i] == '|' {
			if line[i+1] == '|' {
				tokens = append(tokens, token{"COMPARISON", "||"}) // ||
				i++
			}
		} else if line[i] == '&' {
			if line[i+1] == '&' {
				tokens = append(tokens, token{"COMPARISON", "&&"}) // &&
				i++
			}
		} else if line[i] == 'c' {
			if line[i+1] == 'h' {
				if line[i+2] == 'a' {
					if line[i+3] == 'r' {
						tokens = append(tokens, token{"CHAR", "char"}) // char
						i = i + 3
					}
				}
			}
		} else if line[i] == 'f' {
			if line[i+1] == 'o' {
				if line[i+2] == 'r' {
					tokens = append(tokens, token{"LOOP", "for"}) // for
					i = i + 2
				}
			}
		} else if line[i] == 'i' {
			if line[i+1] == 'f' {
				tokens = append(tokens, token{"CONDITION", "if"}) // if
				i = i + 1
			}
		} else if line[i] == 'l' {
			if line[i+1] == 'e' {
				if line[i+2] == 'f' {
					if line[i+3] == 't' {
						tokens = append(tokens, token{"DIRECTION", "left"}) // left
						i = i + 3
					}
				}
			}
		} else if line[i] == 'm' {
			if line[i+1] == 'o' {
				if line[i+2] == 'v' {
					if line[i+3] == 'e' {
						tokens = append(tokens, token{"ACTION", "move"}) // move
						i = i + 3
					}
				}
			} else if line[i+1] == 'e' {
				if line[i+2] == 'm' {
					if line[i+3] == 'o' {
						if line[i+4] == 'r' {
							if line[i+5] == 'y' {
								tokens = append(tokens, token{"MEMORY", "memory"}) // memory
								i = i + 5
							}
						}
					}
				}
			}
		} else if line[i] == 'n' {
			if line[i+1] == 'e' {
				if line[i+2] == 'w' {
					tokens = append(tokens, token{"SPECIAL", "new"}) // new
					i = i + 2
				}
			}
		} else if line[i] == 'p' {
			if line[i+1] == 'r' {
				if line[i+2] == 'i' {
					if line[i+3] == 'n' {
						if line[i+4] == 't' {
							tokens = append(tokens, token{"PRINT", "print"}) // print
							i = i + 4
						}
					}
				}
			} else if line[i+1] == 'o' {
				if line[i+2] == 'i' {
					if line[i+3] == 'n' {
						if line[i+4] == 't' {
							if line[i+5] == 'e' {
								if line[i+6] == 'r' {
									tokens = append(tokens, token{"POINTER", "pointer"}) // pointer
									i = i + 6
								}
							}
						}
					}
				}
			}
		} else if line[i] == 'r' {
			if line[i+1] == 'i' {
				if line[i+2] == 'g' {
					if line[i+3] == 'h' {
						if line[i+4] == 't' {
							tokens = append(tokens, token{"DIRECTION", "right"}) // right
							i = i + 4
						}
					}
				}
			}
		} else if line[i] == 's' { 
			if line[i+1] == 'p' {
				if line[i+2] == 'a' {
					if line[i+3] == 'c' {
						if line[i+4] == 'e' {
							tokens = append(tokens, token{"SPECIAL", "space"}) // space
							i = i + 4
						}
					}
				}
			}
			
		} else if line[i] == 't' {
			if line[i+1] == 'a' {
				if line[i+2] == 'b' {
					tokens = append(tokens, token{"SPECIAL", "tab"}) // tab
					i = i + 2
				}
			}
		} else if line[i] == '0' || line[i] == '1' || line[i] == '2' || line[i] == '3' ||
			line[i] == '4' || line[i] == '5' || line[i] == '6' || line[i] == '7' ||
			line[i] == '8' || line[i] == '9' {
			numeral = numeral + string(line[i])
			if i == len(line)-1 {
				tokens = append(tokens, token{"NUMBER", numeral}) // numberal
				numeral = ""
			}
		} else {
			printError(line, lineNumber, i)
		}
	}

	return tokens
}

// Utility Functions
func checkArithmetic(tokens []token, line string, lineNumber int) []token {

	if tokens[0].tokenType == "MEMORY" {
		if tokens[1].tokenType == "ASSIGNMENT" {
			if tokens[2].tokenType == "MEMORY" {
				if tokens[3].tokenType == "ARITHMETIC" {
					if tokens[4].tokenType == "MEMORY" {
						return tokens[:5]
					} else if tokens[4].tokenType == "POINTER" {
						return tokens[:5]
					} else if tokens[4].tokenType == "NUMBER" {
						return tokens[:5]
					}
					printParseError(line, lineNumber, tokens[4].tokenValue)
				}
				printParseError(line, lineNumber, tokens[3].tokenValue)
			} else if tokens[2].tokenType == "POINTER" {
				if tokens[3].tokenType == "ARITHMETIC" {
					if tokens[4].tokenType == "MEMORY" {
						return tokens[:5]
					} else if tokens[4].tokenType == "POINTER" {
						return tokens[:5]
					} else if tokens[4].tokenType == "NUMBER" {
						return tokens[:5]
					}
					printParseError(line, lineNumber, tokens[4].tokenValue)
				}
			} else if tokens[2].tokenType == "NUMBER" {
				if tokens[3].tokenType == "ARITHMETIC" {
					if tokens[4].tokenType == "MEMORY" {
						return tokens[:5]
					} else if tokens[4].tokenType == "POINTER" {
						return tokens[:5]
					} else if tokens[4].tokenType == "NUMBER" {
						return tokens[:5]
					}
					printParseError(line, lineNumber, tokens[4].tokenValue)
				}
			}
			printParseError(line, lineNumber, tokens[2].tokenValue)
		}
		printParseError(line, lineNumber, tokens[1].tokenValue)
	}

	return nil
}

func checkAssignment(tokens []token, line string, lineNumber int) []token {

	if tokens[0].tokenType == "MEMORY" {
		if tokens[1].tokenType == "ASSIGNMENT" {
			if tokens[2].tokenType == "POINTER" {
				return tokens[:3]
			} else if tokens[2].tokenType == "NUMBER" {
				return tokens[:3]
			}
			printParseError(line, lineNumber, tokens[2].tokenValue)
		}
		printParseError(line, lineNumber, tokens[1].tokenValue)
	}

	return nil
}

func checkPrint(tokens []token, line string, lineNumber int) []token {

	if tokens[0].tokenType == "PRINT" {
		if tokens[1].tokenType == "MEMORY" {
			return tokens[:2]
		} else if tokens[1].tokenType == "POINTER" {
			return tokens[:2]
		}

		if tokens[1].tokenType == "CHAR" {
			if tokens[2].tokenType == "POINTER" {
				return tokens[:3]
			} else if tokens[2].tokenType == "MEMORY" {
				return tokens[:3]
			} else if tokens[2].tokenType == "NUMBER" {
				return tokens[:3]
			}

			printParseError(line, lineNumber, tokens[2].tokenValue)
		}

		if tokens[1].tokenType == "NUMBER" {
			return tokens[:2]
		}

		if tokens[1].tokenType == "SPECIAL" {
			return tokens[:2]
		}
		printParseError(line, lineNumber, tokens[1].tokenValue)
	}

	return nil
}

func checkAction(tokens []token, line string, lineNumber int) []token {

	if tokens[0].tokenType == "ACTION" {
		if tokens[2].tokenType == "DIRECTION" {
			if tokens[1].tokenType != "POINTER" {
				printParseError(line, lineNumber, tokens[1].tokenValue)
			}

			if len(tokens) == 4 {
				if tokens[3].tokenType == "NUMBER" {
					return tokens[:4]
				}
			}

			return tokens[:3]

		}

		printParseError(line, lineNumber, tokens[2].tokenValue)
	}

	return nil
}

func checkCondition(tokens []token, line string, lineNumber int) []token {

	if tokens[0].tokenType == "CONDITION" {
		if tokens[1].tokenType == "MEMORY" {
			if tokens[2].tokenType == "COMPARISON" {
				if tokens[3].tokenType == "POINTER" {
					return tokens[:4]
				} else if tokens[3].tokenType == "MEMORY" {
					return tokens[:4]
				} else if tokens[3].tokenType == "NUMBER" {
					return tokens[:4]
				}
				printParseError(line, lineNumber, tokens[3].tokenValue)
			}
			printParseError(line, lineNumber, tokens[2].tokenValue)
		} else if tokens[1].tokenType == "POINTER" {
			if tokens[2].tokenType == "COMPARISON" {
				if tokens[3].tokenType == "POINTER" {
					return tokens[:4]
				} else if tokens[3].tokenType == "MEMORY" {
					return tokens[:4]
				} else if tokens[3].tokenType == "NUMBER" {
					return tokens[:4]
				}
				printParseError(line, lineNumber, tokens[3].tokenValue)
			}
			printParseError(line, lineNumber, tokens[2].tokenValue)
		} else if tokens[1].tokenType == "NUMBER" {
			if tokens[2].tokenType == "COMPARISON" {
				if tokens[3].tokenType == "POINTER" {
					return tokens[:4]
				} else if tokens[3].tokenType == "MEMORY" {
					return tokens[:4]
				} else if tokens[3].tokenType == "NUMBER" {
					return tokens[:4]
				}
				printParseError(line, lineNumber, tokens[3].tokenValue)
			}
			printParseError(line, lineNumber, tokens[2].tokenValue)
		}
		printParseError(line, lineNumber, tokens[1].tokenValue)
	}
	return nil
}

func checkLoop(tokens []token, line string, lineNumber int) []token {

	if tokens[0].tokenType == "LOOP" {
		if tokens[1].tokenType == "POINTER" {
			return tokens[:2]
		} else if tokens[1].tokenType == "MEMORY" {
			return tokens[:2]
		} else if tokens[1].tokenType == "NUMBER" {
			return tokens[:2]
		}
	}
	printParseError(line, lineNumber, tokens[1].tokenValue)
	return nil
}

func pointerMovements(actionTokens []token, line string, lineNumber int, p int, m int) (int, int) {
	numberOfTokens := len(actionTokens)
	if numberOfTokens == 4 {
		points := actionTokens[3].tokenValue
		pointsToMove, _ := strconv.Atoi(points)

		if actionTokens[2].tokenValue == "right" {
			p = p + pointsToMove
		} else if actionTokens[2].tokenValue == "left" {

			// DO NOT ALLOW P TO BE NEGATIVE
			if p-pointsToMove < 0 {
				runtimeError(line, lineNumber, "Pointer can't point to a negative box")
			} else {
				p = p - pointsToMove
			}
		}
	} else if numberOfTokens == 3 {
		if actionTokens[2].tokenValue == "right" {
			p = p + 1
		} else if actionTokens[2].tokenValue == "left" {
			// DO NOT ALLOW P TO BE NEGATIVE
			if p - 1 < 0 {
				runtimeError(line, lineNumber, "Pointer can't point to a negative box")
			} else {
				p = p - 1
			}
		}
	}

	return p, m
}

func printStuff(printTokens []token, p int, m int) (int, int) {

	functionMessage := ""

	if printTokens[1].tokenType == "MEMORY" {
		functionMessage += strconv.Itoa(m)
	} else if printTokens[1].tokenType == "POINTER" {
		functionMessage += strconv.Itoa(p)
	} else if printTokens[1].tokenType == "NUMBER" {
		functionMessage += printTokens[1].tokenValue
	} else if printTokens[1].tokenType == "CHAR" {
		if printTokens[2].tokenType == "POINTER" {
			functionMessage += returnASCII(p)
		} else if printTokens[2].tokenType == "MEMORY" {
			functionMessage += returnASCII(m)
		} else if printTokens[2].tokenType == "NUMBER" {
			points := printTokens[2].tokenValue
			number, _ := strconv.Atoi(points)
			functionMessage += returnASCII(number)
		}
	} else if printTokens[1].tokenType == "SPECIAL" {
		if printTokens[1].tokenValue == "new" {
			functionMessage += "\n"
		} else if printTokens[1].tokenValue == "tab" {
			functionMessage += "\t"
		} else if printTokens[1].tokenValue == "space" {
			functionMessage += " "
		}
	}

	message += functionMessage

	return p, m
}

func doArithmetic(arithmeticTokens []token, line string, lineNumber int, p int, m int) (int, int) {
	var firstOperand int
	var secondOperand int

	if arithmeticTokens[2].tokenType == "MEMORY" {
		firstOperand = m
	} else if arithmeticTokens[2].tokenType == "POINTER" {
		firstOperand = p
	} else if arithmeticTokens[2].tokenType == "NUMBER" {
		firstOperand, _ = strconv.Atoi(arithmeticTokens[2].tokenValue)
	}

	if arithmeticTokens[4].tokenType == "MEMORY" {
		secondOperand = m
	} else if arithmeticTokens[4].tokenType == "POINTER" {
		secondOperand = p
	} else if arithmeticTokens[4].tokenType == "NUMBER" {
		secondOperand, _ = strconv.Atoi(arithmeticTokens[4].tokenValue)
	}

	if arithmeticTokens[3].tokenValue == "+" {
		m = firstOperand + secondOperand
	} else if arithmeticTokens[3].tokenValue == "-" {
		if firstOperand-secondOperand < 0 {
			runtimeError(line, lineNumber, "Negative number. Memory Cannot Hold Negative Numbers")
		} else {
			m = firstOperand - secondOperand
		}
	} else if arithmeticTokens[3].tokenValue == "*" {
		m = firstOperand * secondOperand
	} else if arithmeticTokens[3].tokenValue == "/" {
		if secondOperand == 0 {
			runtimeError(line, lineNumber, "Cannot Divide Entity by Zero")
		} else {
			testValue := float32(firstOperand) / float32(secondOperand)
			if isIntegral(testValue) {
				m = int(testValue)
			} else {
				runtimeError(line, lineNumber, "Memory Cannot Have Decimal Values")
			}
		}
	}

	return p, m
}

func assignMemory(assignmentTokens []token, p int) (int, int) {

	var memory int
	if assignmentTokens[2].tokenType == "POINTER" {
		memory = p
	} else {
		number := assignmentTokens[2].tokenValue
		memory, _ = strconv.Atoi(number)
	}

	return p, memory
}

func doConditionalCheck(tokens []token, conditionTokens []token, line string, lineNumber int, p int, m int) (int, int) {

	var shouldWeDo bool = false
	var firstOperand int = 0
	var secondOperand int = 0

	if conditionTokens[1].tokenType == "NUMBER" {
		firstOperand, _ = strconv.Atoi(conditionTokens[1].tokenValue)
	} else if conditionTokens[1].tokenType == "MEMORY" {
		firstOperand = m
	} else if conditionTokens[1].tokenType == "POINTER" {
		firstOperand = p
	}

	if conditionTokens[3].tokenType == "NUMBER" {
		secondOperand, _ = strconv.Atoi(conditionTokens[3].tokenValue)
	} else if conditionTokens[3].tokenType == "MEMORY" {
		secondOperand = m
	} else if conditionTokens[3].tokenType == "POINTER" {
		secondOperand = p
	}

	if conditionTokens[2].tokenValue == ">" {
		if firstOperand > secondOperand {
			shouldWeDo = true
		}
	} else if conditionTokens[2].tokenValue == "<" {
		if firstOperand < secondOperand {
			shouldWeDo = true
		}
	} else if conditionTokens[2].tokenValue == "==" {
		if firstOperand == secondOperand {
			shouldWeDo = true
		}
	} else if conditionTokens[2].tokenValue == "!=" {
		if firstOperand != secondOperand {
			shouldWeDo = true
		}
	} else if conditionTokens[2].tokenValue == ">=" {
		if firstOperand >= secondOperand {
			shouldWeDo = true
		}
	} else if conditionTokens[2].tokenValue == "<=" {
		if firstOperand <= secondOperand {
			shouldWeDo = true
		}
	}

	if shouldWeDo {
		unTouchedTokens := tokens[4:]
		p, m = parser(unTouchedTokens, line, lineNumber, p, m)
	}

	return p, m
}

func doLoops (tokens []token, loopTokens []token, line string, lineNumber int, p int, m int ) (int, int) {
	var timesToLoop int

	if loopTokens[1].tokenType == "NUMBER" {
		timesToLoop, _ = strconv.Atoi(loopTokens[1].tokenValue)
	} else if loopTokens[1].tokenType == "MEMORY" {
		timesToLoop = m
	} else if loopTokens[1].tokenType == "POINTER" {
		timesToLoop = p
	}

	tokensToLoop := tokens[len(loopTokens):]
	
	for i := 0; i < timesToLoop; i++ {
		p, m = parser(tokensToLoop, line, lineNumber, p, m)
	}

	return p, m
}


func returnASCII(num int) string {
	return string(rune(num))
}


func isIntegral(val float32) bool {
	return val == float32(int(val))
}

// Parse Function
func parse(line string, lineNumber int, p int, m int) (int, int) {
	tokens := lex(line, lineNumber)
	p, m = parser(tokens, line, lineNumber, p, m)

	return p, m
}

// Parser
func parser(tokens []token, line string, lineNumber int, p int, m int) (int, int) {

	if len(tokens) == 0 {
		return p, m
	}

	if tokens[0].tokenType == "ACTION" {
		actionTokens := checkAction(tokens, line, lineNumber)
		if actionTokens != nil {
			p, m = pointerMovements(actionTokens, line, lineNumber, p, m) // Movement of the pointer
		}
	} else if tokens[0].tokenType == "PRINT" {
		printTokens := checkPrint(tokens, line, lineNumber)
		if printTokens != nil {
			p, m = printStuff(printTokens, p, m)

		}
	} else if tokens[0].tokenType == "MEMORY" {

		if len(tokens) == 3 {
			assignmentTokens := checkAssignment(tokens, line, lineNumber)
			if assignmentTokens != nil {
				p, m = assignMemory(assignmentTokens, p)
			}
		} else if len(tokens) == 5 {
			arithmeticTokens := checkArithmetic(tokens, line, lineNumber)
			if arithmeticTokens != nil {
				p, m = doArithmetic(arithmeticTokens, line, lineNumber, p, m)
			}
		} else {
			printExpectedTokenError(line, lineNumber, "pointer or memory or number")
		}

	} else if tokens[0].tokenType == "CONDITION" {
		conditionTokens := checkCondition(tokens, line, lineNumber)
		if conditionTokens != nil {
			p, m = doConditionalCheck(tokens, conditionTokens, line, lineNumber, p, m)
		}

	} else if tokens[0].tokenType == "LOOP" {
		loopTokens := checkLoop(tokens, line, lineNumber)
		if loopTokens != nil {
			p, m = doLoops(tokens, loopTokens, line, lineNumber, p, m)
		}

	} else {
		printParseError(line, lineNumber, "")
	}

	return p, m
}