package namer

import (
	"encoding/json"
	"fmt"
	"github.com/YoSmudge/server-names/wordlist"
	"github.com/YoSmudge/server-names/words"
	"math/rand"
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

	nouns := words.Nouns()
	prefixes := append(words.Adjectives(), words.Verbs()...)

	var maxId int64 = int64(len(nouns)) * int64(len(prefixes))

	if strings.HasPrefix(serverId, "i-") {
		serverId = strings.TrimLeft(serverId, "i-")
	}

	serverIdDecoded, err := strconv.ParseUint(serverId, 16, 64)
	if err != nil {
		return "", fmt.Errorf("Could not decode provided server ID to integer", err)
	}

	r := rand.New(rand.NewSource(int64(serverIdDecoded)))
	picker := uint64(r.Int63n(maxId))

	quo := picker / uint64(len(nouns))
	mod := picker % uint64(len(nouns))

	pickedNoun := nouns[mod]
	pickedPrefix := prefixes[quo]

	return fmt.Sprintf("%s-%s", pickedPrefix.Word, pickedNoun.Word), nil
}
