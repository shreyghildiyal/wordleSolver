package wordle

import (
	"fmt"
	"strings"
)

type Wordle struct {
	rootWord string
	// wordCharSet map[rune]bool
}

type MatchType int

const (
	NOT_PRESENT    MatchType = 0
	WRONG_POSITION MatchType = 1
	CORRECT        MatchType = 2
)

func GetWordle(word string) Wordle {
	w := Wordle{rootWord: strings.ToLower(word)}

	return w
}

func (w *Wordle) Match(word string) []MatchType {
	if len(w.rootWord) != len(word) {
		fmt.Println("Input Length", len(word))
		fmt.Println("Root Word Length", len(w.rootWord))
		panic("The input string is incorrect length")
	}
	// mr := MatchResponse{CorrectCharacters: make([]bool, len(w.rootWord)), PresentCharacters: []rune{}}
	mr := make([]MatchType, len(word))

	for i := 0; i < len(w.rootWord); i++ {
		if w.rootWord[i] == word[i] {
			mr[i] = CORRECT
		} else if w.charInWord(word[i]) {
			mr[i] = WRONG_POSITION
		} else {
			mr[i] = NOT_PRESENT
		}
	}
	return mr
}

func (w *Wordle) charInWord(ch byte) bool {

	for _, c := range w.rootWord {
		if byte(c) == ch {
			return true
		}
	}

	return false
}

func (w *Wordle) GetWordLength() int {
	return len(w.rootWord)
}
