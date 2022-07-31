DIR=$(shell pwd)

go-format:
	gofmt -l $(DIR) && gofmt -s -w $(DIR)

go-test:
	go test -v -timeout=100s $(DIR)/...