BINARY = main.exe

build:
	go env -w GOOS=windows
	go env -w GOARCH=amd64
	go build -o ${BINARY} .

run: 
	go env -w GOOS=windows
	go env -w GOARCH=amd64
	go build -o ${BINARY} .
	./${BINARY}

clean:
	go clean