/*
 as cool as it would be i don't have the technical skills to develop a full on compiler yet....

 this version of "The super tiny compiler" will hopefully serve as a way to strengthen some skills and fill some knowledge gaps that i have...

 this is more of a learning oriented project than something with actual use but it will hopefully aid me gain more confidence for future low level projects
*/

/**
 * Most compilers break down into three primary stages: Parsing, Transformation,
 * and Code Generation
 *
 * 1. *Parsing* is taking raw code and turning it into a more abstract
 *    representation of the code.
 *
 * 2. *Transformation* takes this abstract representation and manipulates to do
 *    whatever the compiler wants it to.
 *
 * 3. *Code Generation* takes the transformed representation of the code and
 *    turns it into new code.
 */

package main

//"fmt"
//"log"
//"strings"

//lets first define our token structure

type token struct {
	kind  string
	value string
}

//awesome, now we can start with our tokenizer

func tokenizer(input string) []token {
	// we append a new line to the passed input
	input += "\n"

	current := 0

	//initilize an empty slice of our token type to append to
	tokens := []token{}

	for current < len([]rune(input)) {
		char := string([]rune(input)[current]) //stores current character

		//first thing we want to check for are open parenthesis

		//there is definetly a more efficient way of doing this but we'll optimize later
		if char == "(" {
			//if our current char happens to be an opennign parenthesis we append a new token to our slice
			tokens = append(tokens, token{
				kind:  "paren",
				value: "(",
			})

			//we increment current to move onto the next character in our input
			current++
			continue

		}
		// the process above will pretty much repeat for the all of the tokens we consider... next down the line would be a closing parenthesis
		if char == ")" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: ")",
			})
			current++
			continue
		}

		//whitespaces however are interesting we care that they are there to separate the characters in queston but its not important for us to store it as a token
		if char == " " {
			current++
			continue
		}

		//numbers as well as strings differently --> they could be of any given length and its important for us to capture the entire entry as one single token
		if isNumber(char) {
			val := "" //this is a str that will be used to capture the entirety of the number being passed

			for isNumber(char) {
				val += char
				current++
				char = string([]rune(input)[current])
			}

			//once the loop exits we known that we've stored every digit and we can append
			tokens = append(tokens, token{
				kind:  "number",
				value: val,
			})
			continue
		}

		//the last type of token we'll need for this tiny tiny tiny compiler is a 'name' token. This is a sequence of letters instead of numbers that are the names of functions in our lisp syntax
		if isLetter(char) {
			//we follow the same principle we did for the number token
			val := ""

			for isLetter(char) {
				val += char
				current++
				char = string([]rune(input)[current])
			}
			tokens = append(tokens, token{
				kind:  "name",
				value: val,
			})
			continue
		}

		break //if none of the conditions are met whatever we are encountering is not a valid token so we can break our of the iteration
	}

	return tokens

}

func isNumber(char string) bool {

	if char == "" {
		return false
	}

	n := []rune(char)[0]

	if n >= '0' && n <= '9' {
		return true
	}

	return false
}
func isLetter(char string) bool {
	if char == "" {
		return false
	}

	n := []rune(char)[0]

	if (n >= 'a' && n <= 'z') || (n >= 'A' && n <= 'Z') {
		return true
	}

	return false
}
