package solver

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/shreyghildiyal/wordleSolver/wordle"
)

func getResultString(res []wordle.MatchType) string {
	resRunes := make([]string, len(res))
	for i, r := range res {
		if r == wordle.CORRECT {
			resRunes[i] = "G"
		} else if r == wordle.NOT_PRESENT {
			resRunes[i] = "B"
		} else if r == wordle.WRONG_POSITION {
			resRunes[i] = "Y"
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

func PruneWords(words map[string]int32, prevWord string, prevRes []wordle.MatchType) {

	// newWords := map[string]int32{}

	for word := range words {
		// isValid := true
		isValid := IsValidWord(word, prevWord, prevRes)

		if !isValid {
			delete(words, word)
		}
	}

	// return newWords
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

func getCleanedWordsOfLength(l int) map[string]int32 {

	inputFileName2 := "cleanedWords.txt"

	file2, err := os.Open(inputFileName2)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	wantedList := map[string]int32{}

	scanner := bufio.NewScanner(file2)

	for scanner.Scan() {
		line := scanner.Text()
		// parts := strings.Split(line, ",")
		// if _, found := wantedList[parts[0]]; found {
		// 	freq, err := strconv.ParseInt(parts[1], 10, 32)
		// 	if err == nil {
		// 		wantedList[parts[0]] = int32(freq)
		// 	}
		// }
		if len(line) == l {
			wantedList[line] = 0
		}

	}

	return wantedList
}

func populateWordFrequencies(wantedList map[string]int32) {
	inputFileName2 := "unigram_freq.csv"

	file2, err := os.Open(inputFileName2)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	scanner := bufio.NewScanner(file2)

	for scanner.Scan() {
		line := scanner.Text()
		// word = strings.Trim(word, "\n")
		// a := r.MatchString(word)
		parts := strings.Split(line, ",")
		if _, found := wantedList[parts[0]]; found {
			freq, err := strconv.ParseInt(parts[1], 10, 32)
			if err == nil {
				wantedList[parts[0]] = int32(freq)
			}

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Wanted List length", len(wantedList))
	// return wantedList
}

func removeUnusedWords(wantedList map[string]int32) {
	for word, freq := range wantedList {
		if freq == 0 {
			delete(wantedList, word)
		}
	}
}

func getWordsOfLength(l int) map[string]int32 {

	wantedList := getCleanedWordsOfLength(l)
	fmt.Println("cleanedWords count", len(wantedList))
	populateWordFrequencies(wantedList)
	removeUnusedWords(wantedList)
	fmt.Println("After populating the frequencies", len(wantedList))

	// inputFileName1 := "cleanedWords.txt"

	// wantedList := map[string]int32{}

	// optionally, resize scanner's capacity for lines over 64K, see next example

	// regexString := fmt.Sprintf("^([a-z]|[A-Z]){%d}$", l)
	// r, _ := regexp.Compile(regexString)
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
			PruneWords(possibleWords, tryWord, matchRes)
			if len(possibleWords) < 40 {
				fmt.Println("POSSIBLE", possibleWords)
			}
		}

	}
}

func GetRepresentativeWord2(words map[string]int32) string {
	// charactersCount := 26
	// charFrequencyMat := [][]int{}

	representativeWord := ""
	var representativeFreq int32

	for w, f := range words {
		representativeWord = w
		representativeFreq = f
		break
	}

	representativeScore := len(words)

	totalWords := len(words)
	i := 0

	for word, freq := range words {
		// startTime := time.Now()
		score := GetSplitScore(word, words)
		// fmt.Println("time to get score for one word", time.Since(startTime).Milliseconds())
		// fmt.Println(word, score)
		if score < representativeScore {
			representativeWord = word
			representativeScore = score
			representativeFreq = freq
			// fmt.Println(representativeWord, representativeScore, representativeFreq)
		} else if score == representativeScore && freq > representativeFreq {
			representativeWord = word
			representativeScore = score
			representativeFreq = freq
			// fmt.Println(representativeWord, representativeScore, representativeFreq)
		}
		i++
		if i%100 == 0 {
			fmt.Println(i, "/", totalWords)
			fmt.Println(representativeWord, representativeScore, representativeFreq)
		}
	}
	fmt.Println(representativeWord, representativeScore)
	return representativeWord
}

func GetSplitScore(word string, words map[string]int32) int {
	possibleMatchresults := GetPossibleMatchResults(len(word))
	// for _, res := range possibleMatchresults {
	// 	fmt.Println(getResultString(res))
	// }
	maxScore := 0
	for _, possibleMatchres := range possibleMatchresults {
		score := 0
		for w := range words {
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
