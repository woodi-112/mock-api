
all: help

.PHONY: test
test:
		go test ./...

test-cover: prepare-cover
		go test ./... -coverprofile=.cover/cover.out

prepare-cover: clean
	mkdir .cover/

clean:
	rm -rf .cover/
