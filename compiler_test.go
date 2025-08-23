package main

import (
	"slices"
	"testing"
	"fmt"
	"reflect"
)

func TestIsDigit(t *testing.T) {
	if !isNumber("5") {
		t.Fatal("expected '5' to be a digit")
	}
	if isNumber("w") {
		t.Fatal("expected 'w' to be reported as not a number")
	}
	if isNumber("") {
		t.Fatal("expected \"\" to be reported as not a number")
	}
}

func TestIsLetter(t *testing.T) {
	if !isLetter("a") {
		t.Fatal("expected 'a' to be reported as a letter")
	}
	if isLetter("3") {
		t.Fatal("3 is not a letter")
	}
	if isLetter("") {
		t.Fatal("an empty char is not a letter")
	}
}

func TestTokenizer(t *testing.T) {
	input := "(add 2 (substract 4 2))"

	tokens := tokenizer(input)

	expected := []token{
		{kind: "paren", value: "("},
		{kind: "name", value: "add"},
		{kind: "number", value: "2"},
		{kind: "paren", value: "("},
		{kind: "name", value: "substract"},
		{kind: "number", value: "4"},
		{kind: "number", value: "2"},
		{kind: "paren", value: ")"},
		{kind: "paren", value: ")"},
	}
	if !slices.Equal(tokens, expected) {
		t.Fatal("tokenizer mismatch")
	}

}
func generateTestAst() ast{
	var testAst = ast{
		kind: "Program",
		body: []node{
			node{
				kind: "CallExpression",
				name: "add",
				params: []node{
					node{
						kind: "NumberLiteral",
						value: "2",
					},
					node{
						kind: "CallExpression",
						name: "substract",
						params: []node{
							node{
								kind: "NumberLiteral",
								value: "4",
							},
							node{
								kind: "NumberLiteral",
								value: "2",
							},
						},
					},
				},
			},
		},
	}
	return testAst
}
func TestParser(t *testing.T){
	expected := []token{
		{kind: "paren", value: "("},
		{kind: "name", value: "add"},
		{kind: "number", value: "2"},
		{kind: "paren", value: "("},
		{kind: "name", value: "substract"},
		{kind: "number", value: "4"},
		{kind: "number", value: "2"},
		{kind: "paren", value: ")"},
		{kind: "paren", value: ")"},
	}

	ast := parser(expected)
	testAst  := generateTestAst()
	if !reflect.DeepEqual(ast, testAst){
		t.Error("\nExpected:",testAst, "\nGot:", ast)
	}
	fmt.Println(ast)
		
}
