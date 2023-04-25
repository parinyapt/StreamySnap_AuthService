install:
	go mod tidy

updateptgu:
	go get github.com/parinyapt/golang_utils@latest

run:
	go run cmd/main.go -mode=development

runprod:
	go run cmd/main.go -mode=production