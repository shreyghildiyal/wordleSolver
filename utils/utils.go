package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func GenerateFileWithLength(l int) {
	inputFileName := "allWords.txt"
	outputFileName := fmt.Sprintf("wordLen%d.txt", l)
	fmt.Println("OutputFile:", outputFileName)

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
			wantedList = append(wantedList, word)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Wanted List length", len(wantedList))
	writeSliceToFile(wantedList, outputFileName)
}

func writeSliceToFile(slice []string, fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range slice {
		_, _ = datawriter.WriteString(strings.ToLower(data) + "\n")
	}

	datawriter.Flush()
	file.Close()
}

func IsRealWord(word string) bool {
	fmt.Println("Checking word", word)

	// url := fmt.Sprintf("https://dictionary.cambridge.org/dictionary/english/%s", word)
	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)
	fmt.Println("URL:", url)
	resp, err := http.Get(url)

	if err != nil {
		return false
	} else {

		fmt.Println("Content Length", resp.ContentLength)

		var bodyBytes []byte
		if resp.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(resp.Body)
		}
		if strings.Contains(string(bodyBytes), "Sorry pal, we couldn't find definitions for the word you were looking for.") {
			return false
		}

	}

	return true
}

type DictWord struct {
	Word string `json:"word"`
}

func CreateVerifiedWordsList() {
	inputFileName := "EDMTDictionary.json"
	outputFileName := "cleanedWords.txt"
	fmt.Println("OutputFile:", outputFileName)

	jsonFile, err := os.Open(inputFileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened input file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	writefile, err := os.OpenFile(outputFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer writefile.Close()
	datawriter := bufio.NewWriter(writefile)

	dictWords := []DictWord{}

	// inputDat := []byte{}
	inputDat, err := ioutil.ReadAll(jsonFile)
	// _, err = jsonFile..Read(inputDat)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(inputDat, &dictWords)
	if err != nil {
		log.Fatal(err)
	}

	// actualWords := make([]string, len(dictWords))

	foundWords := map[string]bool{}

	for i, word := range dictWords {
		// actualWords[i] = word.Word
		if i%50 == 0 {
			fmt.Println(i)
		}
		if _, found := foundWords[word.Word]; !found {

			_, err = datawriter.WriteString(strings.ToLower(word.Word) + "\n")
			if err != nil {
				log.Fatal(err)
			}
			datawriter.Flush()
			foundWords[word.Word] = true
		}
	}

	// for scanner.Scan() {
	// 	word := scanner.Text()
	// 	word = strings.Trim(word, "\n")

	// 	// wantedList = append(wantedList, word)

	// 	_, err = datawriter.WriteString(strings.ToLower(word) + "\n")
	// 	if err != nil {
	// 		fmt.Println("Error writing string", err)
	// 	}
	// 	datawriter.Flush()

	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

}
