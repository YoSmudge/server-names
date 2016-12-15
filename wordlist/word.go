package wordlist

import (
	"sort"
	"strings"
)

type WordList struct {
	Words []Word
}

type Word struct {
	Type string
	Word string
}

func (w *WordList) Add(wd Word) {
	w.Words = append(w.Words, wd)
}

func (w *WordList) byType(ty string) []*Word {
	sort.Sort(w)
	var n []*Word
	for i, _ := range w.Words {
		wd := w.Words[i]
		if wd.Type == ty {
			n = append(n, &wd)
		}
	}
	return n
}

func (w *WordList) Nouns() []*Word {
	return w.byType("noun")
}

func (w *WordList) Verbs() []*Word {
	return w.byType("verb")
}

func (w *WordList) Adjectives() []*Word {
	return w.byType("adjective")
}

func (w *WordList) Len() int {
	return len(w.Words)
}

func (w *WordList) Swap(a, b int) {
	w.Words[a], w.Words[b] = w.Words[b], w.Words[a]
}

func (w *WordList) Less(a, b int) bool {
	if strings.Compare(w.Words[a].Word, w.Words[b].Word) == 1 {
		return true
	}
	return false
}
