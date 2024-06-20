BINARY = main.exe
OS = linux

build:
	go env -w GOOS=${OS}
	go env -w GOARCH=amd64
	go build -o ${BINARY} .

run: 
	go env -w GOOS=${OS}
	go env -w GOARCH=amd64
	go build -o ${BINARY} .
	./${BINARY}

clean:
	go clean
