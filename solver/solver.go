package solver

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/shreyghildiyal/wordleSolver/wordle"
)

func Solve(w wordle.Wordle) {

	possibleWords := getWordsOfLength(w.GetWordLength())
	// tryWord := "slate"
	for i := 0; ; i++ {
		tryWord := GetRepresentativeWord(possibleWords)
		// fmt.Println("possible words Count", len(possibleWords))
		fmt.Println("trying", tryWord, "attempt", i+1)
		matchRes := w.Match(tryWord)
		fmt.Println("result", getResultString(matchRes))
		allMatch := true
		for _, res := range matchRes {
			if res != wordle.CORRECT {
				allMatch = false
			}
		}
		if allMatch {
			break
		} else {
			possibleWords = PruneWords(possibleWords, tryWord, matchRes)
		}

	}

}

func getResultString(res []wordle.MatchType) string {
	resRunes := make([]string, len(res))
	for i, r := range res {
		if r == wordle.CORRECT {
			resRunes[i] = "C"
		} else if r == wordle.NOT_PRESENT {
			resRunes[i] = "N"
		} else if r == wordle.WRONG_POSITION {
			resRunes[i] = "W"
		}
	}
	return strings.Join(resRunes, "")
}

func IsValidWord(word, prevWord string, prevRes []wordle.MatchType) bool {
	// isValid := true

	for i, res := range prevRes {
		char := prevWord[i]
		if res == wordle.CORRECT {
			if word[i] != char {
				return false

			}
		} else if res == wordle.WRONG_POSITION {
			if word[i] == char {
				return false

			}
			foundChar := false
			for _, ch := range word {
				if byte(ch) == char {
					foundChar = true
				}
			}
			if !foundChar {
				return false

			}

		} else if res == wordle.NOT_PRESENT {
			for _, ch := range word {
				if byte(ch) == char {
					return false

				}
			}
		}
	}

	return true
}

func PruneWords(words []string, prevWord string, prevRes []wordle.MatchType) []string {

	newWords := []string{}

	for _, word := range words {
		// isValid := true
		isValid := IsValidWord(word, prevWord, prevRes)
		// wordChecks:

		// 	for i, res := range prevRes {
		// 		char := prevWord[i]
		// 		if res == wordle.CORRECT {
		// 			if word[i] != char {
		// 				isValid = false
		// 				break wordChecks
		// 			}
		// 		} else if res == wordle.WRONG_POSITION {
		// 			if word[i] == char {
		// 				isValid = false
		// 				break wordChecks
		// 			}
		// 			foundChar := false
		// 			for _, ch := range word {
		// 				if byte(ch) == char {
		// 					foundChar = true
		// 				}
		// 			}
		// 			if !foundChar {
		// 				isValid = false
		// 				break wordChecks
		// 			}

		// 		} else if res == wordle.NOT_PRESENT {
		// 			for _, ch := range word {
		// 				if byte(ch) == char {
		// 					isValid = false
		// 					break wordChecks
		// 				}
		// 			}
		// 		}
		// 	}
		if isValid {
			newWords = append(newWords, word)
		}
	}

	return newWords
}

func GetRepresentativeWord(words []string) string {
	charactersCount := 26
	charFrequencyMat := [][]int{}

	if len(words) == 0 {
		panic("The slice is empty")
	}

	for i := 0; i < len(words[0]); i++ {
		charFrequencyMat = append(charFrequencyMat, make([]int, charactersCount))
	}

	for _, word := range words {
		for i, ch := range word {
			charFrequencyMat[i][ch-'a'] += 1
		}
	}

	representativeWord := ""
	score := 0
	for _, word := range words {
		wordScore := 0
		distinctChars := map[rune]bool{}
		for i, ch := range word {
			wordScore += charFrequencyMat[i][ch-'a']
			distinctChars[ch] = true
		}
		score = score * len(distinctChars)
		if wordScore > score {
			representativeWord = word
			score = wordScore
		}
	}
	return representativeWord
}

func getWordsOfLength(l int) []string {

	inputFileName := "cleanedWords.txt"

	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	wantedList := []string{}
	regexString := fmt.Sprintf("^([a-z]|[A-Z]){%d}$", l)
	r, _ := regexp.Compile(regexString)

	for scanner.Scan() {
		word := scanner.Text()
		word = strings.Trim(word, "\n")
		a := r.MatchString(word)
		if a {
			wantedList = append(wantedList, strings.ToLower(word))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Wanted List length", len(wantedList))
	return wantedList
}

func Solve2(w wordle.Wordle) {
	possibleWords := getWordsOfLength(w.GetWordLength())

	for i := 0; ; i++ {

		tryWord := GetRepresentativeWord2(possibleWords)
		// fmt.Println("possible words Count", len(possibleWords))
		fmt.Println("trying", tryWord, "attempt", i+1)
		matchRes := w.Match(tryWord)
		fmt.Println("result", getResultString(matchRes))
		allMatch := true
		for _, res := range matchRes {
			if res != wordle.CORRECT {
				allMatch = false
			}
		}
		if allMatch {
			break
		} else {
			possibleWords = PruneWords(possibleWords, tryWord, matchRes)
		}

	}
}

func GetRepresentativeWord2(words []string) string {
	// charactersCount := 26
	// charFrequencyMat := [][]int{}

	representativeWord := words[0]
	representativeScore := len(words)

	for _, word := range words {
		score := GetSplitScore(word, words)
		// fmt.Println(word, score)
		if score < representativeScore {
			representativeWord = word
			representativeScore = score
		}
	}
	fmt.Println(representativeWord, representativeScore)
	return representativeWord
}

func GetSplitScore(word string, words []string) int {
	possibleMatchresults := GetPossibleMatchResults(len(word))
	// for _, res := range possibleMatchresults {
	// 	fmt.Println(getResultString(res))
	// }
	maxScore := 0
	for _, possibleMatchres := range possibleMatchresults {
		score := 0
		for _, w := range words {
			if IsValidWord(w, word, possibleMatchres) {
				score++
			}
		}
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}

func GetPossibleMatchResults(num int) [][]wordle.MatchType {
	res := [][]wordle.MatchType{}
	currentRes := make([]int, num)
mainLoop:
	for {
		r := make([]wordle.MatchType, num)
		for i, c := range currentRes {
			r[i] = wordle.MatchType(c)
		}
		res = append(res, r)
		Increment(currentRes, 3)
		for _, j := range currentRes {
			if j != 0 {
				continue mainLoop
			}
		}
		break

	}
	return res
}

func Increment(current []int, max int) {

	current[0]++
	for i := 0; i < len(current); i++ {
		if current[i] >= max {
			current[i] = 0
			if i+1 < len(current) {
				current[i+1] = current[i+1] + 1
			}
		}
	}
}
