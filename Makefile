
GOPATH=$(PWD)/.go

echo:
	@echo "$(GOPATH)"
	find . -path ./.go -prune -o -name "*.go" -exec gofmt -w {} \;
	find . -path ./.go -prune -o -name "*.i2pkeys" -exec rm {} \;

build:
	go build -a -tags netgo -tags webface -ldflags '-w -extldflags "-static"'

deps:
	go get -u golang.org/x/time/rate
	go get -u github.com/eyedeekay/sam-forwarder/manager

test:
	go test -tags webface

cssfile:
	cd css && ./snippet.sh

jsfile:
	cd js && ./snippet.sh

cssget:
	wget -q -O css/styles.css localhost:7957/css/styles.css

jsget:
	wget -q -O js/script.js localhost:7957/js/script.js

index:
	wget -q -O index.html localhost:7957/index
