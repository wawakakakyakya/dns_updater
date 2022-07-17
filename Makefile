# /bin/sh

fileName := dns_updater

# all: all-test export-all-coverage

mod_tidy:
	go mod tidy
build_linux: clean
	GOOS=linux GOARCH=amd64 go build -o ${fileName} main.go
clean:
	rm -f ${fileName}
test:
	./${fileName} -dest ../ -src ../
