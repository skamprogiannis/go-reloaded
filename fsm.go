package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func tokenize(text string) []string {
	var tokens []string
	var currentToken strings.Builder

	runes := []rune(text)
	n := len(runes)

	for i := 0; i < n; i++ {
		r := runes[i]

		if r == '\n' {
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			tokens = append(tokens, "\n")
			continue
		}

		if unicode.IsSpace(r) {
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			continue
		}

		if strings.ContainsRune(".,!?:;", r) {
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}

			pToken := string(r)
			j := i + 1
			for j < n && strings.ContainsRune(".,!?:;", runes[j]) {
				pToken += string(runes[j])
				j++
			}
			i = j - 1
			tokens = append(tokens, pToken)
			continue
		}

		if r == '\'' {
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			tokens = append(tokens, "'")
			continue
		}

		if r == '(' {
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}

			start := i
			for i < n && runes[i] != ')' {
				i++
			}

			if i < n {
				tokens = append(tokens, string(runes[start:i+1]))
			} else {
				tokens = append(tokens, string(runes[start:]))
			}
			continue
		}

		currentToken.WriteRune(r)
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}

func capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	runes := []rune(strings.ToLower(s))
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}

func isVowelOrH(r rune) bool {
	return strings.ContainsRune("aeiouhAEIOUH", r)
}

func Process(text string) string {
	tokens := tokenize(text)
	var processedTokens []string

	// Pass 1: Apply markers
	for _, token := range tokens {
		if strings.HasPrefix(token, "(") && strings.HasSuffix(token, ")") {
			content := token[1 : len(token)-1]
			parts := strings.Split(content, ",")
			command := strings.TrimSpace(parts[0])
			count := 1

			if len(parts) > 1 {
				c, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil {
					count = c
				}
			}

			apply := func(fn func(string) string) {
				n := len(processedTokens)
				start := n - count
				if start < 0 {
					start = 0
				}
				for k := start; k < n; k++ {
					processedTokens[k] = fn(processedTokens[k])
				}
			}

			switch command {
			case "hex":
				apply(func(s string) string {
					val, err := strconv.ParseInt(s, 16, 64)
					if err == nil {
						return fmt.Sprintf("%d", val)
					}
					return s
				})
				continue
			case "bin":
				apply(func(s string) string {
					val, err := strconv.ParseInt(s, 2, 64)
					if err == nil {
						return fmt.Sprintf("%d", val)
					}
					return s
				})
				continue
			case "up":
				apply(strings.ToUpper)
				continue
			case "low":
				apply(strings.ToLower)
				continue
			case "cap":
				apply(capitalize)
				continue
			}
		}

		processedTokens = append(processedTokens, token)
	}

	// Pass 2: Grammar (a -> an)
	for i := 0; i < len(processedTokens)-1; i++ {
		curr := processedTokens[i]
		next := processedTokens[i+1]

		if len(next) > 0 {
			nextStart := []rune(next)[0]
			if isVowelOrH(nextStart) {
				if curr == "a" {
					processedTokens[i] = "an"
				} else if curr == "A" {
					processedTokens[i] = "An"
				}
			}
		}
	}

	// Pass 3: Reconstruction (Spacing & Quotes)
	var result strings.Builder
	inQuote := false

	for i, token := range processedTokens {
		addSpace := true

		isPunct := len(token) > 0 && strings.ContainsRune(".,!?:;", []rune(token)[0])
		isQuote := token == "'"
		isNewline := token == "\n"

		// Determine current quote state changes
		isClosingQuote := false
		if isQuote {
			if inQuote {
				isClosingQuote = true
			}
		}

		// 1. Punctuation: No space before
		if isPunct {
			addSpace = false
		}

		// 2. Closing Quote: No space before
		if isClosingQuote {
			addSpace = false
		}

		// 3. Previous was Opening Quote
		if i > 0 {
			prev := processedTokens[i-1]
			if prev == "'" {
				if inQuote {
					// Prev was opening
					addSpace = false
				}
			}

			if processedTokens[i-1] == "\n" {
				addSpace = false
			}
		}

		if isNewline {
			addSpace = false
		}

		// First token logic (don't add space at very beginning)
		if i == 0 {
			addSpace = false
		}

		if addSpace {
			result.WriteRune(' ')
		}

		result.WriteString(token)

		// Update State
		if isQuote {
			inQuote = !inQuote
		}
	}

	return result.String()
}
