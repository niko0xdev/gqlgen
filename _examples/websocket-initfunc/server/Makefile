bin_name=server

build:
	@echo "building binary..."
# go generate gives missing go sum entry for module errors
# https://github.com/niko0xdev/gqlgen/issues/1483
# you will need to first do a go get -u github.com/niko0xdev/gqlgen
	go run -mod=mod github.com/niko0xdev/gqlgen generate .
	go build -o ${bin_name} server.go
