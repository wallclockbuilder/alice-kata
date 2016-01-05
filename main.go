package main

// Print out top 7 frequent words
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
)

type word struct {
	name  string
	count int
}

type byCount []word

func (l byCount) Len() int           { return len(l) }
func (l byCount) Less(i, j int) bool { return l[i].count < l[j].count }
func (l byCount) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func insert(dict map[string]int, word string) map[string]int {
	if _, ok := dict[word]; !ok {
		dict[word] = 1
		return dict
	}
	dict[word] += 1
	return dict
}

func extractWordsFromLine(s string) []string {
	regex := regexp.MustCompile("\\w+")
	words := regex.FindAllString(s, -1)
	return words
}

func assertNil(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func convertToList(m map[string]int) []word {
	// slice converts a map to a slice of words
	list := []word{}
	for k, v := range m {
		list = append(list, word{k, v})
	}
	return list
}

func addWordsToDictionary(dict map[string]int, words []string) map[string]int {
	// addaddWordsToDictionary adds the words in the supplied list to the
	// supplied dictionary and returns the dictionary.
	for _, w := range words {
		dict = insert(dict, w)
	}
	return dict
}

func mostFrequent(wordFrequencyList []word, n int) []word {
	return wordFrequencyList[:n]
}

func main() {
	dict := make(map[string]int)
	wordFrequencyList := []word{}
	f, err := os.Open("alice-in-wonderland.txt")

	assertNil(err)

	r := bufio.NewReader(f)
	for {
		switch s, err := r.ReadString('\n'); err {
		case nil:
			words := extractWordsFromLine(s)
			dict = addWordsToDictionary(dict, words)
		case io.EOF:
			// convert dict to slice for sorting in O(n) time.
			// maps are horrible for sorting because the go
			// runtime randomizes map iteration order.
			// So we use a different data structure(slice) to maintain order.
			wordFrequencyList = convertToList(dict)
			sort.Sort(sort.Reverse(byCount(wordFrequencyList)))
			fmt.Println(mostFrequent(wordFrequencyList, 7))
			return
		default:
			return
		}
	}

}
