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

func main() {
	dict := make(map[string]int)
	list := []word{}
	f, err := os.Open("alice-in-wonderland.txt")
	if err != nil {
		log.Panic(err)
	}
	r := bufio.NewReader(f)

	for {
		switch s, err := r.ReadString('\n'); err {
		case nil:
			words := extractWordsFromLine(s)

			for _, w := range words {
				dict = insert(dict, w)
			}

		case io.EOF:
			fmt.Println("EOF")
			//convert dict to slice for sorting in O(n) time.
			// slices are horrible for sorting
			for k, v := range dict {
				list = append(list, word{k, v})
			}

			sort.Sort(sort.Reverse(byCount(list)))

			// print out top 7 most common words
			fmt.Println(list[:7])
			return

		default:
			return
		}
	}

}
