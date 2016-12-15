package wordlist

import (
	"encoding/json"
	"fmt"
	"github.com/YoSmudge/server-names/words"
	"github.com/stretchr/testify/assert"
	"testing"
)

var wordlist *WordList

func TestWordList(t *testing.T) {
	raw, _ := words.Asset("wordlist.json")
	json.Unmarshal(raw, &wordlist)

	if wordlist.Len() < 1000 {
		assert.Fail(t, "wordlist did not load correctly", fmt.Sprintf("Expected at least %d results but had %d", 1000, wordlist.Len()))
	}
}

func testGrouping(t *testing.T, w []*Word, ty string) {
	if len(w) < 1000 {
		assert.Fail(t, fmt.Sprintf("wordlist did not return %s", ty), fmt.Sprintf("Expected at least %d results but had %d", 1000, len(w)))
	}

	for _, wd := range w {
		assert.Equal(t, ty, wd.Type, "word had invalid type")
	}
}

func TestWordNouns(t *testing.T) {
	testGrouping(t, wordlist.Nouns(), "noun")
}

func TestWordVerbs(t *testing.T) {
	testGrouping(t, wordlist.Verbs(), "verb")
}

func TestWordAdjectives(t *testing.T) {
	testGrouping(t, wordlist.Adjectives(), "adjective")
}

func TestWordByType(t *testing.T) {
	w := wordlist.byType("noun")
	testGrouping(t, w, "noun")
}

func BenchmarkWordByType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wordlist.byType("noun")
	}
}
