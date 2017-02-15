package namer

import (
	"encoding/json"
	"fmt"
	"github.com/YoSmudge/server-names/wordlist"
	"github.com/YoSmudge/server-names/words"
	"github.com/spaolacci/murmur3"
	"sort"
	"strconv"
	"strings"
)

func Name(serverId string) (string, error) {
	wordlistRaw, err := words.Asset("wordlist.json")
	if err != nil {
		return "", fmt.Errorf("Could not open wordlist from embedded data! %s", err)
	}

	var words wordlist.WordList
	err = json.Unmarshal(wordlistRaw, &words)
	if err != nil {
		return "", fmt.Errorf("Could not open wordlist from embedded data! %s", err)
	}

	sort.Sort(&words)

	nouns := words.Nouns()
	prefixes := append(words.Adjectives(), words.Verbs()...)

	var maxId uint64 = uint64(len(nouns)) * uint64(len(prefixes))

	if strings.HasPrefix(serverId, "i-") {
		serverId = strings.TrimLeft(serverId, "i-")
	}

	_, err = strconv.ParseUint(serverId, 16, 64)
	if err != nil {
		return "", fmt.Errorf("Could not decode provided server ID to integer", err)
	}

	hash := murmur3.Sum64([]byte(serverId))

	picker := hash % maxId

	quo := picker / uint64(len(nouns))
	mod := picker % uint64(len(nouns))

	pickedNoun := nouns[mod]
	pickedPrefix := prefixes[quo]

	return fmt.Sprintf("%s-%s", pickedPrefix.Word, pickedNoun.Word), nil
}
