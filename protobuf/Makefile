all:
	protoc --version
	make dnsmessage.pb.go

clean:
	rm -f *.pb.go

dnsmessage.pb.go: dnsmessage.proto
	protoc --go_out=. dnsmessage.proto

compile:
	protoc --go_out=. *.proto

deps:
	apt install golang-goprotobuf-dev
	apt install protobuf-compiler

push:
	git pull
	git add --all
	git commit -a -s
	git push
