# GITVERSION=$(shell git rev-parse FETCH_HEAD)
VERSION=$(shell git describe --tags $(git rev-list --tags --max-count=1) | sed 's/^v//')

BASENAME=control-panel-dns

all: help deb

help:
	@echo
	@echo "make deb     # attempt to build the .deb package using dpkg"
	@echo

deb: clean extract DEBIAN build

clean:
	rm -rf ../files
	rm -f ../*.deb
	rm -f ../*.tar.xz data.tar.xz
	rm -rf DEBIAN

extract:
	mkdir -p ../files/usr/bin
	mkdir -p ../files/usr/lib/control-panel-dns/
	-cp ~/go/src/go.wit.com/gui/toolkit/*.so ../files/usr/lib/control-panel-dns/
	cp  ../README.md ../files/usr/lib/control-panel-dns/
	cp  ../control-panel-dns ../files/usr/bin/

# makes the DEBIAN/ directory
DEBIAN:
	mkdir -p DEBIAN

	# make the md5sum file
	cd ../files/ && find -type f -exec md5sum '{}' \; |sort -k2 >../md5sums
	mv ../md5sums DEBIAN/

	# make the control there
	mkdir -p DEBIAN
	cp control DEBIAN/
	echo Version: ${VERSION} >>DEBIAN/control

	cp postinst DEBIAN

build:
	mv DEBIAN ../files/
	cd .. && dpkg-deb --build files ${BASENAME}_${VERSION}_amd64.deb
	@echo
	@echo '#######################'
	cd .. && dpkg-deb --info ${BASENAME}_${VERSION}_amd64.deb
	@echo '#######################'
	@echo
