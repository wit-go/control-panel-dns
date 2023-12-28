run: build
	./control-panel-cloudflare

build-release:
	go get -v -u -x .
	go build
	./control-panel-cloudflare

build:
	GO111MODULE="off" go get -v -x .
	GO111MODULE="off" go build

update:
	GO111MODULE="off" go get -v -u -x .

log:
	reset
	tail -f /tmp/witgui.* /tmp/guilogfile

gocui: build
	./control-panel-cloudflare -gui gocui >/tmp/witgui.log.stderr 2>&1

quiet:
	./control-panel-cloudflare >/tmp/witgui.log.stderr 2>&1
