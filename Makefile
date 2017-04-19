cmd/remarkgo/remarkgo: *.go template/* cmd/remarkgo/main.go
	go-bindata -pkg remark template
	cd cmd/remarkgo && go build

.PHONY: setup
setup:
	go get -u github.com/jteeuwen/go-bindata/...
