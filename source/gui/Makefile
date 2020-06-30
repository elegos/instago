.DEFAULT_GOAL := build

vendor/github.com/therecipe/qt:
	go mod vendor

vendor/github.com/therecipe/env_linux_amd64_513:
	git clone https://github.com/therecipe/env_linux_amd64_513.git vendor/github.com/therecipe/env_linux_amd64_513

setup: vendor/github.com/therecipe/qt vendor/github.com/therecipe/env_linux_amd64_513
	qtsetup

build: vendor/github.com/therecipe/qt vendor/github.com/therecipe/env_linux_amd64_513
	mkdir -p ../../out
	qtdeploy build linux .

clean:
	rm -rf vendor
	rm -rf ../../out/qt-instago
