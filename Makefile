.PHONY: run build git

run:
	go run main.go

build:
	go build

git: build
	git diff
	git add .
	git commit
	git push

