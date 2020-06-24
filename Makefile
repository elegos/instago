.DEFAULT_GOAL := build

build:
	mkdir -p out
	go mod download
	go build -o out/instago source/cmd/cmd.go
	cp res/InstaGo.desktop out/InstaGo.desktop
	cp res/config.yml out/config.yml

clean:
	rm -rf out/*