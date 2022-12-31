# /bin/sh

fileName := dns_updater
version := $(shell git describe --tags --abbrev=0)
revision := $(shell git rev-parse --short HEAD)

mod_tidy:
	go mod tidy
build_linux: clean
	GOOS=linux GOARCH=amd64 go build -ldflags "-X dns_updater/version.Version=$(version) -X  dns_updater/version.Revision=$(revision)" -o ${fileName} main.go
clean:
	rm -f ${fileName}
test:
	./${fileName} -dest ../ -src ../
