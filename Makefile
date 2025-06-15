VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`
HASH=`git log -1 --format=%H`
AUTHOR=`git log -1 --format=%ce`
LDFLAGS=-ldflags "-w -s -X cmd.Version=${VERSION}  -X cmd.Build=${HASH}"

build:
	go mod tidy
	go build -o bin/ ./...

clean:
	go clean