cmd: deps
	go build -ldflags "-X main.consumerKey '$(POCKET_CONSUMER_KEY)'" ./cmd/pocket

deps:
	go get ./...

test: testdeps
	go test ./...

testdeps:
	go get -t ./...

.PHONY: cmd deps
