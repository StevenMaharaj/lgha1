module_name = $(shell basename $(CURDIR))

build:
	go build -o bin/$(module_name) cmd/$(module_name)/main.go

git:
	git add .
	git commit -m "$(m)"
	git push

test:
	go test -v ./tests/
