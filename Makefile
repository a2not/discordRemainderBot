.PHONY: run git

run:
	go run main.go

git:
	git diff
	git add .
	git commit
	git push

