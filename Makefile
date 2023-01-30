buildAt=`date +%Y/%m/%d-%H:%M:%S`
gitHash=`git rev-parse HEAD`
version=`git rev-parse --abbrev-ref HEAD | grep -v HEAD || git describe --exact-match HEAD || git rev-parse HEAD`  ## todo: use current release git tag
flags="-X 'atom/utils.Version=${version}' -X 'atom/utils.BuildAt=${buildAt}' -X 'atom/utils.GitHash=${gitHash}'"

GOPATH:=$(shell go env GOPATH)

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: dist
dist:
	@go build -ldflags=${flags} -o bin/atom
	@cp config.toml bin/

.PHONY: test
test:
	@go test -v ./... -cover