DIR=$(shell pwd)

format:
	gofmt -l $(DIR) && gofmt -s -w $(DIR)

test-set:
	go test -v -timeout=20s -count=1 $(DIR)/set/...

test-map:
	go test -v -timeout=50s -count=1 $(DIR)/map/...

test-linked-list:
	go test -v -timeout=10s -count=1 $(DIR)/linked-list/...

test-queue:
	go test -v -timeout=10s -count=1 $(DIR)/queue/...

test-all: test-set test-map test-linked-list test-queue