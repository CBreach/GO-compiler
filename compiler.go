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

import "log"

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
/*

 The parser...

For our parser were going to take our array of tokens and turn it into an AST

[{type: 'paren', value: '('},...] => {type: 'Program, body: [...]'}

we will define our type node here. within node are pointers types to what would otherwise be recursive types in GO.

*/
type node struct{
	kind string
	value string
	name string 
	callee *node
	expression *node
	body []node
	params []node
	arguments *[]node
	context *[]node
}
/*type ast is just an allias type for node. we do this because it makes things clearer dude to the the hih number of reference there are to 'node'*/
type ast node

//this is the counter variable that will be used for counting
var pc int

// this variable will store our slice of 'token's inside of it
var pt []token

/*
now we can defin our parser function, it will take a slice of tokens to create the ast
*/
func parser (tokens []token) ast {
	pc = 0 
	pt = tokens

	//this is the root node for our AST, our AST will have a Program node as the root
	ast := ast{
		kind: "Program",
		body: []node{},
	}

	for pc < len(pt){
		ast.body = append(ast.body, walk())
	}	
	return ast
}
func walk() node{

	//we begin by grabbing the current token
	token := pt[pc]
	
	//we are going to split each type of token off into a different code path, starting off with 'number' tokns
	if token.kind == "number"{
		pc++

		//and we will return a new AST called 'NumberLiteral' and settig its value to the value of our token
		return node{
			kind: "NumberLiteral",
			value: token.value,
		}
	}
	//next were going to look for calledExpessions. we start this off when we encounter an opening parenthesis
	if token.kind == "paren" && token.value == "("{
		//we will increment 'current' to skip the parenthsis since we don't care about it in this AST
		pc++
		token = pt[pc]

		/*
			we create a base node with the type CalledExpression, and we're going to set the name as the curren't token's val
			since the next token after th eopen parenthesis is the name of the function
		*/
		n := node{
			kind: "CallExpression",
			name: token.value,
			params: []node{},
		}
		//we increment the 'curr' again to skip the name token
		pc++
		token = pt[pc]

		for token.kind != "paren" || (token.kind == "paren" && token.value != ")"){
			n.params = append(n.params, walk())
			token = pt[pc]
		}
		pc++
		return n
	}
	//if we don't recognize the token kind we throw an error
	log.Fatal(token.kind)
	return node{}
}

/*
	The traverser--> now that we have our AST, we wan't to be able to visit different nodes within it and most importantly we want 
	to be able to call methods on the visitor node whenever we encounter a node of a given type 

	for this we'll utilize some interesting syntax basically, we'll create a map of functions where the key is the name of our function
*/
//so for what i understand visitor is now a map that takes a string as a key and has a functon associated with that key.. this way we can perform different actions on different node types
type visitor map[string]func(n *node, p node)

func traverser(a ast, v visitor){
	//we kickstart the traverser by calling traverseNode with our ast 
	//no parent is passed because well... the root node has no parent
	
}
func traverseArray(a []node, p node, v visitor){
	for _, child := range a {
		traverseNode(child, p, v)
	}
}
func traverseNode(n, p node, v visitor){
	//we start by testing for the existance of a method on the visitor with a matching type
	for key, value := range v{
		if key == n.kind{
			value(&n,p)
		}
	}
	switch n.kind{
		case "Program":
			traverseArray(n.body,n,v)
			break
		case "CallExpression":
			traverseArray(n.params, n,v)
			break
		case "NumberLiteral":
			break
		default:
			log.Fatal(n.kind)
	
	}
}
