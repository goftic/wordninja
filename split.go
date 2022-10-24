package wordninja

import (
	"unicode"

	r "github.com/goftic/wordninja/reader"
)

/*
I did not author this code, tweaked it from:
    - https://github.com/keredson/wordninja
    - https://github.com/willsmil/go-wordninja
*/

var maxLenWord int
var wordCost map[string]float64

func init() {
	words := r.LoadWords()
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
