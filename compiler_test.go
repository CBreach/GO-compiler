package main

import (
	"slices"
	"testing"
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
