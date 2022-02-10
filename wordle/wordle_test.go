package wordle_test

import (
	"testing"
	"wordleSolver/wordle"
)

func TestMatch(t *testing.T) {

	w := wordle.GetWordle("soler")

	resp := w.Match("sopre")

	expectedResp := []wordle.MatchType{wordle.CORRECT, wordle.CORRECT, wordle.NOT_PRESENT, wordle.WRONG_POSITION, wordle.WRONG_POSITION}

	for i := 0; i < len(resp); i++ {
		if resp[i] != expectedResp[i] {
			t.Error("The expected response is incorrect")
		}
	}
}
