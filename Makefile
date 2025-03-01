
.PHONY: help
help:
	less Makefile

.PHONE: run
run: 
	go run main.go

.PHONY: test
test:
	go test ./src/...

.PHONY: debug
debug:
	dlv debug 