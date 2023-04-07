run: build
	./control-panel-dns >/tmp/witgui.log.stderr 2>&1

install:
	go install -v go.wit.com/control-panel-dns@latest
	# go install -v git.wit.com/wit/control-panel-dns@latest

debug: build
	./control-panel-dns --verbose --verbose-net --gui-debug

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

# ./control-panel-dns.v1: /lib/x86_64-linux-gnu/libc.so.6: version `GLIBC_2.34' not found (required by ./control-panel-dns.v1)
# ./control-panel-dns.v1: /lib/x86_64-linux-gnu/libc.so.6: version `GLIBC_2.32' not found (required by ./control-panel-dns.v1)
# compiling with CGO disabled means it compiles but then plugins don't load
GLIBC_2.34-error:
	GO111MODULE="off" CGO_ENABLED=0 go build -v -o control-panel-dns

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
#
check-cert:
	reset
	# https://crt.sh/?q=check.lab.wit.org
	# # https://letsencrypt.org/certificates/
	# openssl s_client -connect check.lab.wit.org:443 -showcerts
	openssl s_client -CApath /etc/ssl/certs/ -connect check.lab.wit.org:443 -showcerts
	# openssl s_client -CApath /etc/ssl/certs/ -connect check.lab.wit.org:443 -showcerts -trace -debug
	# openssl s_client -CAfile isrgrootx1.pem -connect check.lab.wit.org:443 -showcerts
	# cat isrgrootx1.pem lets-encrypt-r3.pem > full-chain.pem
	# full-chain.pem
	# openssl s_client -CAfile /etc/ssl/certs/wit-full-chain.pem -connect check.lab.wit.org:443 -showcerts

ssl-cert-hash:
	openssl x509 -hash -noout -in wit-full-chain.pem
	# cd /etc/ssl/certs && ln -s wit-full-chain.pem 4042bcee.0
	openssl x509 -hash -noout -in isrgrootx1.pem
	openssl x509 -hash -noout -in lets-encrypt-r3.pem

sudo-cp:
	sudo cp -a lets-encrypt-r3.pem 8d33f237.0 /etc/ssl/certs/

go-get:
	go install -v check.lab.wit.org/gui

log:
	reset
	tail -f /tmp/witgui.* /tmp/guilogfile
