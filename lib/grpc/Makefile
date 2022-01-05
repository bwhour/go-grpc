all: vet test testrace

build: deps
	go build github.com/bwhour/go-grpc/lib/grpc/...

clean:
	go clean -i github.com/bwhour/go-grpc/lib/grpc/...

deps:
	go get -d -v github.com/bwhour/go-grpc/lib/grpc/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate github.com/bwhour/go-grpc/lib/grpc/...

test: testdeps
	go test -cpu 1,4 -timeout 7m github.com/bwhour/go-grpc/lib/grpc/...

testappengine: testappenginedeps
	goapp test -cpu 1,4 -timeout 7m github.com/bwhour/go-grpc/lib/grpc/...

testappenginedeps:
	goapp get -d -v -t -tags 'appengine appenginevm' github.com/bwhour/go-grpc/lib/grpc/...

testdeps:
	go get -d -v -t github.com/bwhour/go-grpc/lib/grpc/...

testrace: testdeps
	go test -race -cpu 1,4 -timeout 7m github.com/bwhour/go-grpc/lib/grpc/...

updatedeps:
	go get -d -v -u -f github.com/bwhour/go-grpc/lib/grpc/...

updatetestdeps:
	go get -d -v -t -u -f github.com/bwhour/go-grpc/lib/grpc/...

vet: vetdeps
	./vet.sh

vetdeps:
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	deps \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testdeps \
	testrace \
	updatedeps \
	updatetestdeps \
	vet \
	vetdeps