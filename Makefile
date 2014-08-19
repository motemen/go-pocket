cmd: deps
	go build ./cmd/pocket

deps:
	go get ./...

test: testdeps
	go test ./...

testdeps:
	go get -t ./...

.PHONY: cmd deps test tesdeps
