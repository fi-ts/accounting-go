
all: proto mocks test

.PHONY: proto
proto:
	rm -rf pkg/apis
	make -C proto protoc

.PHONY: mocks
mocks:
	rm -rf test/mocks
	docker run --rm --user $$(id -u):$$(id -g) -w /work -v ${PWD}:/work vektra/mockery:v3.7.2

.PHONY: test
test:
	go test -v ./...

.PHONY: check-diff
check-diff:
	git diff --exit-code
