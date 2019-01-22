GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

test: test-all test-bench

test-all:
	@go test -ldflags -s -v $(GOPACKAGES)

test-bench:
# DEBUG=false bash -c "go test -v nbd-gitlab/NBD/MBA/rest-api/route/routelogin -bench=. -run BenchmarkLoginHandler"
	@go test -v $(GOPACKAGES) -bench . -run=^Benchmark
