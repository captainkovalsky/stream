install:
	go get -v ./...
	go build -o app
run: 
	go run main.go --help
