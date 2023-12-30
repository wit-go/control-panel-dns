# export GO111MODULE="off"
run: build
	./control-panel-digitalocean --gui-debug

build-release:
	go get -v -u -x .
	go build
	./control-panel-digitalocean

build:
	go get -v -x .
	go build

update:
	go get -v -u -x .

log:
	reset
	tail -f /tmp/witgui.* /tmp/guilogfile

gocui: build
	./control-panel-digitalocean -gui gocui

quiet:
	./control-panel-digitalocean >/tmp/witgui.log.stderr 2>&1
