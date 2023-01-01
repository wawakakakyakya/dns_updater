# /bin/sh

fileName := dns_updater
version := $(shell git describe --tags --abbrev=0)
revision := $(shell git rev-parse --short HEAD)
arch := nil

mod_tidy:
	go mod tidy
set_amd64:
	$(eval arch := amd64)
set_arm:
	$(eval arch := arm)

build:
	GOOS=linux GOARCH=$(arch) go build -ldflags "-X dns_updater/version.Version=$(version) -X  dns_updater/version.Revision=$(revision)" -o ${fileName}_for_$(arch) main.go
build_linux: clean set_amd64 build
build_arm: clean set_arm build

clean:
	rm -f ${fileName}
test:
	./${fileName} -dest ../ -src ../
