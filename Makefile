default: build

build:
	go build -o bin/ssh-sync-keys sync.go

install:
	cp bin/ssh-sync-keys  /usr/local/bin/ssh-sync-keys

purge:
	rm /usr/local/bin/ssh-sync-keys