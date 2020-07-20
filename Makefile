.PHONY: check
check: vet test

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test ./...
