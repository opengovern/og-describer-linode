.PHONY: build

local-build: clean
	CC=/usr/bin/musl-gcc GOPRIVATE="github.com/opengovern" GOOS=linux GOARCH=amd64 go build -a -v -mod=mod -ldflags "-linkmode external -extldflags '-static' -s -w" -tags musl -o ./local/og-describer-linode main.go

build-cli: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -ldflags "-w -extldflags -static" -o ./build/og-linode-cli ./command/main.go

clean:
	rm -rf ./local/og-describer-linode ./build/og-linode-cli
