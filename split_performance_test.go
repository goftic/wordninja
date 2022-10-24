package wordninja

import (
	"testing"
)

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("thereismassesoftextinformationofpeoplescommentswhichisparsedfromhtmlbuttherearenodelimitedcharactersinthemforexamplethumbgreenappleactiveassignmentweeklymetaphorapparentlytherearethumbgreenappleetcinthestringialsohavealargedictionarytoquerywhetherthewordisreasonablesowhatsthefastestwayofextractionthxalot")
	}
}
