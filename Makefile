build:
	go build -o .

update-dependencies:
	go mod tidy

download-dependencies:
	go mod download