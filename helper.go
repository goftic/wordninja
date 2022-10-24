package wordninja

import (
	"errors"
	"math"
	"regexp"
	"strings"
)

type match struct {
	cost float64
	idx  int
}

type text struct {
	s string
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

// min return the minimal cost of the matches.
func minCost(matches []match) (match, error) {
	if len(matches) == 0 {
		return match{}, errors.New("match.len")
	}
	r := matches[0]
	for _, m := range matches {
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
