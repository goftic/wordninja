package wordninja

import (
	"bufio"
	"embed"
	"errors"
	"io"
	"math"
	"regexp"
	"strings"
	"unicode"
)

var maxLenWord int
var wordCost map[string]float64

//go:embed dictionary
var wordFile embed.FS

type match struct {
	cost float64
	idx  int
}

type text struct {
	s string
}

func init() {
	words := loadWords()
	generateSplitWordMap(words)
}

func Split(s string) []string {
	eng := getEnglishText(s)
	return SplitEnglish(eng)
}

// SplitEnglish return the best matched words with spliting the English string `s`.
func SplitEnglish(eng string) []string {
	costs := []float64{0}
	text := text{s: eng}

	for i := 1; i < len(eng)+1; i++ {
		if m, err := text.bestMatch(costs, i); err == nil {
			costs = append(costs, m.cost)
		}
	}

	var out []string
	i := len(eng)

	for i > 0 {
		m, err := text.bestMatch(costs, i)
		if err != nil {
			continue
		}

		newToken := true

		//ignore a lone apostrophe
		if !(eng[i-m.idx:i] == "'") {
			if len(out) > 0 {
				//re-attach split 's and split digits or digit followed by digit.
				if out[len(out)-1] == "'s" ||
					(unicode.IsDigit(rune(eng[i-1])) && unicode.IsDigit(rune(out[len(out)-1][0]))) {
					// combine current token with previous token.
					out[len(out)-1] = eng[i-m.idx:i] + out[len(out)-1]
					newToken = false
				}
			}
		}

		if newToken {
			word := eng[i-m.idx : i]
			out = append(out, word)
		}
		i -= m.idx
	}

	return reverse(out)
}

// load the english split words from file, return list of words.
func loadWords() []string {
	words, err := readFileByLine()
	if err != nil {
		panic("load english split word failed," + err.Error())
	}

	return words
}

// readFileByLine returns a list by reading file line by line.
func readFileByLine() (lines []string, err error) {
	f, err := wordFile.Open("dictionary/wordninja_words.txt")
	if err != nil {
		return lines, err
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if line == "\n" {
			continue
		}

		line = strings.Replace(line, "\n", "", -1)
		lines = append(lines, line)
	}

	return lines, nil
}

// generateSplitWordMap initialize wordCost map and the length of the longest word.
// the cost of a word is calculated by `log(log(len(words))*idx)`.
func generateSplitWordMap(words []string) {
	wordCost = make(map[string]float64)
	var wordLen int
	logLen := math.Log(float64(len(words)))

	for idx, word := range words {
		wordLen = len(word)
		if wordLen > maxLenWord {
			maxLenWord = wordLen
		}

		wordCost[word] = math.Log(logLen * float64(idx+1))
	}
}

// bestMatch will return the minimal cost and its appropriate character's index.
func (s *text) bestMatch(costs []float64, i int) (match, error) {
	var matches []match
	var l, h float64 = 0, float64(i - maxLenWord)
	candidates := costs[int(math.Max(l, h)):i]
	k := 0

	for j := len(candidates) - 1; j >= 0; j-- {
		cost := getWordCost(strings.ToLower(s.s[i-k-1:i])) + float64(candidates[j])
		matches = append(matches, match{cost: cost, idx: k + 1})
		k++
	}

	return minCost(matches)
}

// getWordCost return cost of word from the wordCost map.
// if the word is not exist in the map, it will return `9e99`.
func getWordCost(word string) float64 {
	if v, ok := wordCost[word]; ok {
		return v
	}

	return 9e99
}

// min return the minimal cost of the matchs.
func minCost(matchs []match) (match, error) {
	if len(matchs) == 0 {
		return match{}, errors.New("match.len ")
	}
	r := matchs[0]
	for _, m := range matchs {
		if m.cost < r.cost {
			r = m
		}
	}

	return r, nil
}

// reverse returns reversed list of `dst`
func reverse(dst []string) []string {
	length := len(dst)
	for i := 0; i < length/2; i++ {
		dst[i], dst[length-i-1] = dst[length-i-1], dst[i]
	}

	return dst
}

// getEnglishText return all the English characters of string `s`.
func getEnglishText(s string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9']+")
	return strings.Join(reg.Split(s, -1), "")
}
