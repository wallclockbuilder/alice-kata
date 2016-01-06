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
	"strings"
)

type word struct {
	name  string
	count int
}

type byCount []word

// func Len(), func Less, func Swap are needed so that byCount implements
// the requisite sort.Interface for type word{}
func (l byCount) Len() int           { return len(l) }
func (l byCount) Less(i, j int) bool { return l[i].count < l[j].count }
func (l byCount) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func insert(dict map[string]int, word string) map[string]int {
	// insert adds the input word to the input dictionary
	// and return the modified dictionary.
	if _, ok := dict[word]; !ok {
		dict[word] = 1
		return dict
	}
	dict[word] += 1
	return dict
}

func words(s string) []string {
	// words returns a slice of the words from the input string.
	regex := regexp.MustCompile("\\w+")
	words := regex.FindAllString(s, -1)
	return words
}

func assertNil(err error) {
	// assertNil makes sure theres no error. log panics when there is.
	if err != nil {
		log.Panic(err)
	}
}

func list(m map[string]int) []word {
	// list converts the supplied map to a slice of words
	var list []word
	for k, v := range m {
		list = append(list, word{k, v})
	}
	return list
}

func dictionary(dict map[string]int, words []string) map[string]int {
	// dictionary adds the words in the supplied list to the
	// supplied dictionary and returns the dictionary.
	for _, w := range words {
		dict = insert(dict, strings.ToLower(w))
	}
	return dict
}

func first(wordFrequencyList []word, n int) []word {
	// first returns the first n words from supplied list.
	return wordFrequencyList[:n]
}

func main() {
	// main prints out top 7 frequent words of input file.

	dict := make(map[string]int)
	f, err := os.Open("alice-in-wonderland.txt")
	assertNil(err)

	r := bufio.NewReader(f)
	for {
		switch line, err := r.ReadString('\n'); err {
		case nil:
			dict = dictionary(dict, words(line))
		case io.EOF:
			// convert dict to slice for sorting in O(n) time.
			// maps are horrible for sorting because the go
			// runtime randomizes map iteration order.
			// So we use a different data structure(slice) to maintain order.
			wordFrequencyList := list(dict)

			// sort orders the slice from smallest to biggest
			// sort.reverse ensures an ordering from big to small
			sort.Sort(sort.Reverse(byCount(wordFrequencyList)))

			fmt.Println(first(wordFrequencyList, 7))
			return
		default:
			return
		}
	}

}
