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

	for word, res := range previousAttempts {
		matchRes := getMatchRes(res)
		PruneWords(possibleWords, word, matchRes)
	}
	fmt.Println("After Pruning word count", len(possibleWords))

	fmt.Println(GetRepresentativeWord2(possibleWords))

	// for _, w := range allWords {
	// 	isValid := true
	// wordCheck:
	// 	for i, ch := range w {
	// 		if correctLetters[i] != '.' && byte(ch) != correctLetters[i] {
	// 			isValid = false
	// 			break wordCheck
	// 		}
	// 		for _, invalidChar := range invalidLetters {
	// 			if invalidChar == ch {
	// 				isValid = false
	// 				break wordCheck
	// 			}
	// 		}
	// 		for _, wrongPosChar := range wrongPositionLetters {
	// 			found := false
	// 			for _, ch := range w {
	// 				if ch == wrongPosChar {
	// 					found = true
	// 					break
	// 				}
	// 			}
	// 			if !found {
	// 				isValid = false
	// 				break wordCheck
	// 			}
	// 		}
	// 	}

	// 	if isValid {
	// 		fmt.Println("valid word", w)
	// 	}
	// }
}
