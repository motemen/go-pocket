cmd: deps
	go build -ldflags "-X main.consumerKey '$(POCKET_CONSUMER_KEY)'" ./cmd/pocket

deps:
	go get ./...

.PHONY: cmd deps
