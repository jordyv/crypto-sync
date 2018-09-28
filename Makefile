
build-current:
	go build -o crypto-sync crypto-sync.go

build-all: build-linux-386 build-linux-amd64 build-linux-arm build-osx-386 build-osx-amd64

build-linux-386: crypto-sync.go
	env GOOS=linux GOARCH=386 go build -o bin/linux/386/crypto-sync crypto-sync.go

build-linux-amd64: crypto-sync.go
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/crypto-sync crypto-sync.go

build-linux-arm: crypto-sync.go
	env GOOS=linux GOARCH=arm go build -o bin/linux/arm/crypto-sync crypto-sync.go

build-osx-386: crypto-sync.go
	env GOOS=darwin GOARCH=386 go build -o bin/osx/amd64/crypto-sync crypto-sync.go

build-osx-amd64: crypto-sync.go
	env GOOS=darwin GOARCH=amd64 go build -o bin/osx/arm/crypto-sync crypto-sync.go

clean:
	rm -rf bin/*
