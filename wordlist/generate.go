package wordlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Generate(source string, dest string) error {
	words, err := loadWords(source)
	if err != nil {
		return err
	}
	fmt.Printf("Loaded words, total %d words matching filters\n", len(words.Words))
	fmt.Printf(" - Nouns: %d, Verbs: %d, Adjectives: %d\n", len(words.Nouns()), len(words.Verbs()), len(words.Adjectives()))

	fileData, err := json.Marshal(words)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dest, fileData, 0644)
	if err != nil {
		return err
	}

	return nil
}
