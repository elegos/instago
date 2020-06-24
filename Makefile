.DEFAULT_GOAL := build

build:
	mkdir -p out
	go mod download
	go build -o out/instago cmd.go
	cp InstaGo.desktop out/InstaGo.desktop

clean:
	rm -rf out/*