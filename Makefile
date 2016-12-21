ifndef $(GOOS)
GOOS=$(shell uname | tr '[:upper:]' '[:lower:]')
endif
WORDNET_VERSION=3.1
DOCKER_RUN=docker run -e GOOS=$(GOOS) -e GITHUB_API_KEY -it -v $(shell pwd):/app/src/github.com/YoSmudge/server-names aws-server-names:latest

build: build/server-names

build/server-names: glide words/words.go
	gofmt -w $(shell glide novendor -x)
	go build -o build/server-names

test:
	go test -v -bench=. -benchtime=10s -benchmem $(shell glide novendor)

glide:
	glide install

words/words.go: build/wordlist-generator words/wordlist.json
	${GOPATH}/bin/go-bindata -o words/words.go -nomemcopy -pkg words -prefix words/ words/wordlist.json

build/wordlist-generator: glide
	go build -o build/wordlist-generator github.com/YoSmudge/server-names/wordlist-generator

words/wordlist.json: words/ tmp/dict/index.verb
	build/wordlist-generator --source tmp/dict/ --dest words/wordlist.json

container-build: container
	$(DOCKER_RUN) make build

tmp/:
	mkdir tmp/

words/:
	mkdir words/

tmp/dict/index.verb: tmp/wordnet-$(WORDNET_VERSION).tar.gz
	tar -xf tmp/wordnet-$(WORDNET_VERSION).tar.gz -C tmp/
	touch tmp/dict/index.verb

tmp/wordnet-$(WORDNET_VERSION).tar.gz: tmp/
	wget -q -O tmp/wordnet-$(WORDNET_VERSION).tar.gz http://wordnetcode.princeton.edu/wn$(WORDNET_VERSION).dict.tar.gz
	touch tmp/wordnet-$(WORDNET_VERSION).tar.gz

container:
	docker build . -t aws-server-names:latest

container-console: container
	$(DOCKER_RUN) bash

clean:
	rm -Rf tmp

release-container: container
	rm -Rf outputs/
	$(DOCKER_RUN) make release

release:
	rm -Rf outputs/ .goxc.local.json
	goxc
	@goxc -wlc publish-github -apikey=${GITHUB_API_KEY}
	goxc
