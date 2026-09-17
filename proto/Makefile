MAKEFLAGS += --no-print-directory
BUF_VERSION := 1.72.0

.PHONY: generate-buf-gen-yaml
generate-buf-gen-yaml:
	go run ./generate.go

_buf:
	docker run --rm \
		--entrypoint sh \
		-v $(PWD)/..:/workspace \
		-w /workspace/proto \
		bufbuild/buf:$(BUF_VERSION) \
		-c "buf $(CMD) && chown -R $(shell id -u):$(shell id -g) /workspace"

.PHONY: protolint
protolint: generate-buf-gen-yaml
	@$(MAKE)	_buf	CMD="format -w"
	@$(MAKE)	_buf	CMD="lint -v"

.PHONY: protoc
protoc: protolint
	@$(MAKE)	_buf	CMD="generate -v"
