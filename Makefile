echo:
	@echo "$(GOPATH)"
	find . -name "*.go" -exec gofmt -w {} \;
	find . -name "*.i2pkeys" -exec rm {} \;
