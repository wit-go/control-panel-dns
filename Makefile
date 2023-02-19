run: build
	./control-panel-dns

verbose: build
	./control-panel-dns --verbose --verbose-net --gui-debug --toolkit-debug

dns: build
	./control-panel-dns --verbose-dns

build-release:
	reset
	go get -v -u -x .
	go build

build:
	reset
	# GO111MODULE="off" go get -v -x .
	GO111MODULE="off" go build -v -o control-panel-dns

test:
	GO111MODULE="off" go test -v

update:
	GO111MODULE="off" go get -v -u -x .

clean:
	-rm control-panel-dns
	-rm -rf files/
	-rm *.deb

deb:
	cd debian && make
	-wit mirrors

netlink:
	GO111MODULE="off" go get -v -u github.com/vishvananda/netlink


####### MODULE STUFF DOWN HERE
#
#       What again is the 'right' way to do this?
#       It seems like it changes from year to year. This is better than 'vendor/' (that was a terrible hack)
#       maybe it's settled down finally. Use GO111MODULE="off" when you are developing. (?)
#       When you are ready to release, version this and all the packages correctly. (?)
#
#       At least, that is what I'm going to try to do as of Feb 18 2023.
#


build-with-custom-go.mod:
	go build -modfile=local.go.mod ./...

# module <yourname>
# go 1.18
# require (
#     github.com/versent/saml2aws/v2 v2.35.0
# )
# replace github.com/versent/saml2aws/v2 v2.35.0 => github.com/marcottedan/saml2aws/v2 master
# replace github.com/versent/saml2aws/v2 => /Users/dmarcotte/git/saml2aws/
