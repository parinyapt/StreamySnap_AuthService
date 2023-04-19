install:
	go mod tidy

updateptgu:
	go get github.com/parinyapt/golang_utils@latest

run:
	go run main.go -mode=development