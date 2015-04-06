test: main.go bindata.go
	go run main.go bindata.go --help

build: main.go bindata.go
	go build -o godown main.go bindata.go

bindata.go:
	go-bindata -nomemcopy=true -ignore=.DS_Store resources/...

clean:
	rm bindata.go
