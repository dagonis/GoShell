.PHONY: mac windows linux

mac:
	go build

windows: 
	GOOS=windows GOARCH=amd64 go build

linux:
	GOOS=linux GOARCH=amd64 go build