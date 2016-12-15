package wordlist

import (
	"bufio"
	"os"
	"path"
	"regexp"
	"strings"
)

const MIN_WORD_LENGTH int = 3
const MAX_WORD_LENGTH int = 16

var sourceFiles []string = []string{"index.noun", "index.adj", "index.verb"}
var wordMatch *regexp.Regexp = regexp.MustCompile("^[a-z]+$")

var types map[string]string = map[string]string{
	"n": "noun",
	"v": "verb",
	"a": "adjective",
}

func loadWords(source string) (WordList, error) {
	var words WordList

	for _, fn := range sourceFiles {
		f, err := os.Open(path.Join(source, fn))
		if err != nil {
			return words, err
		}
		defer f.Close()

		s := bufio.NewScanner(f)
		for s.Scan() {
			l := s.Text()
			if strings.HasPrefix(l, "  ") {
				continue
			}
			p := strings.Split(l, " ")
			word := p[0]
			if len(word) <= MIN_WORD_LENGTH || len(word) > MAX_WORD_LENGTH || strings.Contains(word, "_") || !wordMatch.Match([]byte(word)) {
				continue
			}

			wordType := types[p[1]]
			if wordType == "" {
				continue
			}

			w := Word{wordType, word}
			words.Add(w)
		}

		err = s.Err()
		if err != nil {
			return words, err
		}

		f.Close()
	}

	return words, nil
}
