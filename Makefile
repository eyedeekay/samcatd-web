
GOPATH=$(PWD)/.go

echo:
	@echo "$(GOPATH)"
	find . -path ./.go -prune -o -name "*.go" -exec gofmt -w {} \;
	find . -path ./.go -prune -o -name "*.i2pkeys" -exec rm {} \;

build:
	go build -a -tags netgo -tags webface -ldflags '-w -extldflags "-static"'

deps:
	go get -u github.com/eyedeekay/sam-forwarder/manager

test:
	go test -tags webface
