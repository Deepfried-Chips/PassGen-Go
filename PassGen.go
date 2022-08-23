package main

import (
	"flag"
	passwordGenerator "github.com/theTardigrade/golang-passwordGenerator"
	"golang.design/x/clipboard"
	"os"
)

var (
	Length           int
	IncludeUpperCase bool
	IncludeLowerCase bool
	IncludeNumbers   bool
	ExcludeAmbiguous bool
)

func init() {
	flag.IntVar(&Length, "L", 128, "Length of the password")
	flag.BoolVar(&IncludeUpperCase, "u", true, "Include upper case letters")
	flag.BoolVar(&IncludeLowerCase, "l", true, "Include lower case letters")
	flag.BoolVar(&IncludeNumbers, "n", true, "Include numbers")
	flag.BoolVar(&ExcludeAmbiguous, "a", true, "Exclude ambiguous characters")
	flag.Parse()
}

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	generate, err2 := passwordGenerator.New(passwordGenerator.Options{
		Len: Length,
		IncludeRunesList: []rune{
			'!', '?', '-', '_', '=', '@', '$',
			'#', '(', ')', '[', ']', '{', '}',
			'<', '>', '+', '/', '*', '\\', '/',
			':', ';', '&', '\'', '"', '%', '^',
		},
		IncludeUpperCaseLetters: IncludeUpperCase,
		IncludeLowerCaseLetters: IncludeLowerCase,
		IncludeDigits:           IncludeNumbers,
		ExcludeAmbiguousRunes:   ExcludeAmbiguous,
	}).Generate()
	if err2 != nil {
		return
	}
	clipboard.Write(clipboard.FmtText, []byte(generate))

	os.Exit(0)
}
