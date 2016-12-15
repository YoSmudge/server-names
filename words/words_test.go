package words

import (
	"encoding/json"
	"github.com/YoSmudge/server-names/wordlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWords(t *testing.T) {
	raw, err := Asset("wordlist.json")
	assert.Nil(t, err, "wordlist load returned error")

	if len(raw) < 1024 {
		assert.Fail(t, "wordlist had no content")
	}

	var words wordlist.WordList
	err = json.Unmarshal(raw, &words)
	if err != nil {
		assert.Nil(t, err, "wordlist json decode returned error")
	}

	if words.Len() < 1000 {
		assert.Fail(t, "wordlist was too short")
	}
}

func BenchmarkWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		raw, _ := Asset("wordlist.json")
		var words wordlist.WordList
		json.Unmarshal(raw, &words)
	}
}
