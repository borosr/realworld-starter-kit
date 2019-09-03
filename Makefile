mod-reset:
	rm go.mod go.sum
	rm -rf vendor

mod-init:
	go mod init

mod-vendor:
	go mod vendor

mod-update:
	go get -u

mod-all: mod-reset mod-init mod-vendor
