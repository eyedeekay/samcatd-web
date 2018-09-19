
GOPATH=$(PWD)/.go

echo:
	@echo "$(GOPATH)"
	find . -name "*.go" -exec gofmt -w {} \;
	find . -name "*.i2pkeys" -exec rm {} \;

build:
	go build -a -tags netgo -tags webface -ldflags '-w -extldflags "-static"'

deps:
	go get -u github.com/eyedeekay/sam-forwarder/manager
