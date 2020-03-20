all: meta server client

.PHONY: meta

meta:
	cd meta;protoc --go_out=plugins=grpc:. *.proto;cd -
server:
	go build -o bin/$@ cmd/$@/main.go

client:
	go build -o bin/$@ cmd/$@/main.go

clean:
	@rm -rf bin



