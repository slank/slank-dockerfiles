bin/setup: bin setup.go
	GOOS=linux GOARCH=amd64 go build -o bin/setup setup.go

bin:
	mkdir bin

bin/echo: bin/setup
	cd bin && ln -s setup echo
