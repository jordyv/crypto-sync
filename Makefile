
build-current:
	go build -o crypto-sync crypto-sync.go

build-all: build-linux-386 build-linux-amd64 build-linux-arm build-osx-386 build-osx-amd64

build-linux-386: crypto-sync.go
	env GOOS=linux GOARCH=386 go build -o bin/crypto-sync-linux-386 crypto-sync.go

build-linux-amd64: crypto-sync.go
	env GOOS=linux GOARCH=amd64 go build -o bin/crypto-sync-linux-amd64 crypto-sync.go

build-linux-arm: crypto-sync.go
	env GOOS=linux GOARCH=arm go build -o bin/crypto-sync-linux-arm crypto-sync.go

build-osx-386: crypto-sync.go
	env GOOS=darwin GOARCH=386 go build -o bin/crypto-sync-osx-386 crypto-sync.go

build-osx-amd64: crypto-sync.go
	env GOOS=darwin GOARCH=amd64 go build -o bin/crypto-sync-osx-amd64 crypto-sync.go

clean:
	rm -rf bin/*
