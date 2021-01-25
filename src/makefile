BIN_NAME=../bin/iotc
WIN_BIN_NAME=${BIN_NAME}.exe

buildlinux:
	env GOOS=linux GOARCH=amd64 go build -o ${BIN_NAME} -v
build:
	env GOOS=windows GOARCH=amd64 go build -o ${WIN_BIN_NAME} -v
run:
	go run *.go
clean:
	go clean
	rm -f ${BIN_NAME}
	rm -f ${WIN_BIN_NAME}
