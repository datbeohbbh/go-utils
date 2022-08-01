DIR=$(shell pwd)

go-format:
	gofmt -l $(DIR) && gofmt -s -w $(DIR)

go-test-set:
	go test -v -timeout=20s -count=1 $(DIR)/set/...

go-test-map:
	go test -v -timeout=50s -count=1 $(DIR)/map/...

go-test-all: go-test-set go-test-map