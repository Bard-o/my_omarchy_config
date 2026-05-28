.PHONY: build test clean install all

build:
	go build -o waybar-calendar .

test:
	go test -cover ./...

clean:
	rm -f waybar-calendar

install: build
	cp waybar-calendar ~/.local/bin/

all: build test