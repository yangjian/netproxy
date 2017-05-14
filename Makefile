.PHONY: native

native:
	go build -o netproxy

linux:
	GOOS=linux GOARCH=amd64 go build -o netproxy.linux.amd64

clean:
	rm -f netproxy netproxy.linux.amd64
