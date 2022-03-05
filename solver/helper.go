package solver

import (
	"fmt"

	"github.com/shreyghildiyal/wordleSolver/wordle"
)

func getMatchRes(res string) []wordle.MatchType {
	matchRes := make([]wordle.MatchType, len(res))

	charMap := map[byte]wordle.MatchType{
		'G': wordle.CORRECT,
		'Y': wordle.WRONG_POSITION,
		'B': wordle.NOT_PRESENT,
	}

	for i, ch := range res {
		matchRes[i] = charMap[byte(ch)]
	}

	return matchRes
}

func HelpSolve(previousAttempts map[string]string) {

	wordLen := 0
	for word, res := range previousAttempts {
		if len(word) == len(res) {
			wordLen = len(word)
		}

	}

	possibleWords := getWordsOfLength(wordLen)
	fmt.Println("Before Pruning word count", len(possibleWords))

	// minLetterFrequency := map[byte]int{}
	for word, res := range previousAttempts {
		// updateMinLetterFrequency
		matchRes := getMatchRes(res)
		PruneWords(possibleWords, word, matchRes)
	}
	fmt.Println("After Pruning word count", len(possibleWords))

	count := 0
	for word := range possibleWords {
		fmt.Printf("%s, ", word)
		count++
		if count >= 30 {
			break
		}
	}
	fmt.Println("")

	fmt.Println(GetRepresentativeWord2(possibleWords))

}
