GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

test: test-all test-bench

test-all:
	@go test -ldflags -s -v $(GOPACKAGES)

test-bench:
	@go test -v $(GOPACKAGES) -bench . -run=^Benchmark
