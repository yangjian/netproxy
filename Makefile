.PHONY: native

build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o netproxy.darwin.amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o netproxy.linux.amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o netproxy.linux.arm7

clean:
	rm -f netproxy.*
