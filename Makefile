
all: build-current build-arm

build-current: crypto-sync.go
	go build -o bin/crypto-sync crypto-sync.go

build-arm: crypto-sync.go
	env GOOS=linux GOARCH=arm go build -o bin/crypto-sync_arm crypto-sync.go
