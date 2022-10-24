package wordninja

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test(t *testing.T) {
	var testCases = []struct {
		Name     string
		Text     string
		Expected []string
	}{
		{
			Name:     "Test_Simple",
			Text:     "derekanderson",
			Expected: []string{"derek", "anderson"},
		},
		{
			Name:     "Test_Upper_Case",
			Text:     "DEREKANDERSON",
			Expected: []string{"DEREK", "ANDERSON"},
		},
		{
			Name:     "Test_Digits",
			Text:     "win32intel",
			Expected: []string{"win", "32", "intel"},
		},
		{
			Name:     "Test_Apostrophes",
			Text:     `"that'sthesheriff'sbadge"`,
			Expected: []string{"that's", "the", "sheriff's", "badge"},
		},
		{
			Name:     "Test_With_Chinese",
			Text:     "you你aresuch真awonderful好ma",
			Expected: []string{"you", "are", "such", "a", "wonderful", "ma"},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.Expected, Split(tc.Text))
		})
	}
}
