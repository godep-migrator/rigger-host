test:
	go list ./... | xargs -n1 go test

cov:
	gocov test ./... | gocov-html > /tmp/coverage.html
	open /tmp/coverage.html

update_deps:
	godep save -copy=false

deps:
	@echo "--> Installing dependencies"
	@go get ./...

all:
	@mkdir -p bin/
	@bash --norc -i ./build.sh

.PHONY: all cov test update_deps