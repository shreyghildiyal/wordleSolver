package main

import (
	"fmt"

	"github.com/shreyghildiyal/wordleSolver/solver"
	"github.com/shreyghildiyal/wordleSolver/wordle"
)

func main() {
	fmt.Println("Hello World")

	// utils.GenerateFileWithLength(5)

	word := "robin"

	wdl := wordle.GetWordle(word)

	solver.Solve2(wdl)

	// solver.HelpSolve(
	// map[string]string{
	// "raise": "YYBBB",
	// "vertu": "YYBBB",
	// "buyer": "BYBGG",
	// "gluer": "BGYGG",
	// 	},
	// )

	// utils.CreateVerifiedWordsList()
	// word := "point"
	// fmt.Println(word, utils.IsRealWord(word))

	// writeTest()

	// solver.GetSplitScore("abcd", []string{"aaaa", "bbbb", "cccc"})

	// seleniumTry.Example()
}

/*
{
	"word":"A",
	"type":"()",
	"description":"The first letter of the English and of many other alphabets. The capital A of the alphabets of Middle and Western Europe, as also the small letter (a), besides the forms in Italic, black letter, etc., are all descended from the old Latin A, which was borrowed from the Greek Alpha, of the same form; and this was made from the first letter (/) of the Phoenician alphabet, the equivalent of the Hebrew Aleph, and itself from the Egyptian origin. The Aleph was a consonant letter, with a guttural breath sound that was not an element of Greek articulation; and the Greeks took it to represent their vowel Alpha with the a sound, the Phoenician alphabet having no vowel symbols."},
*/
