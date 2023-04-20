all:
	./gengen.py > generate.go
	go generate
	go build ./...

.PHONY: all