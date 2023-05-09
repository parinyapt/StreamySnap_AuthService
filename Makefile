install:
	go mod tidy

updateptgu:
	go get github.com/parinyapt/golang_utils@latest

run:
	go run cmd/main.go -mode=development

runprod:
	go run cmd/main.go 
	
dockerbuild:
	docker build -t streamysnap-authservice .

dockerpush:
	docker tag streamysnap-authservice:latest parinyapt/streamysnap-authservice:v2.5 && docker push parinyapt/streamysnap-authservice:v2.5