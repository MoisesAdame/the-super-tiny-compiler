package compiler

// - Helper Functions: To detect if a given byte is either a number or a letter.
func IsNumber(s byte) bool {
	if s >= '0' && s <= '9' {
		return true
	}

	return false
}

func IsLetter(s byte) bool {
	if s >= 'a' && s <= 'z' {
		return true
	}

	return false
}

// - Token: Stores every token with their respective kind and value as strings
type Token struct {
	kind  string
	value string
}

func (token *Token) ToString() string {
	res := "{ kind: " + token.kind + ", value: " + token.value + " },"
	return res
}

// - Tokenizer: Iterate over the raw code and return an array of Token
func Tokenizer(input string) []Token {
	// Appending new line to the raw code
	input += "\n"

	// Instantiating cursor and token array
	cursor := 0
	tokens := []Token{}

	// Looping trough the raw code and detecting tokens
	for cursor < len(input) {
		// Check if the token is a parenthesis and append them
		if input[cursor] == '(' || input[cursor] == ')' {
			tokens = append(tokens, Token{
				kind:  "paren",
				value: string(input[cursor]),
			})

			cursor++
			continue
		}

		// Check if the token is a number and append a string of it
		if IsNumber(input[cursor]) {
			value := string(input[cursor])
			cursor++
			for IsNumber(input[cursor]) {
				value += string(input[cursor])
				cursor++
			}

			tokens = append(tokens, Token{
				kind:  "number",
				value: value,
			})

			continue
		}

		// Check if the token is a word and append it
		if IsLetter(input[cursor]) {
			value := string(input[cursor])
			cursor++

			for IsLetter(input[cursor]) {
				value += string(input[cursor])
				cursor++
			}

			tokens = append(tokens, Token{
				kind:  "name",
				value: value,
			})
			continue
		}

		// If there are spaces, ignore them
		if input[cursor] == ' ' {
			cursor++
			continue
		}

		break
	}

	return tokens
}

func ToString(tokens []Token) string {
	res := "[\n"
	for _, val := range tokens {
		res += "\t" + val.ToString()
	}

	return res + "\n]"
}
