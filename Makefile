test: main.go bindata.go
	go run main.go bindata.go --help

install_deps_osx:
	sudo cp deps/osx/*.dylib /usr/local/lib/
	sudo cp -R deps/osx/*.framework /Library/Frameworks

build: main.go bindata.go
	go build -o godown main.go bindata.go

bindata.go:
	go-bindata -nomemcopy=true -ignore=.DS_Store resources/...

clean:
	rm bindata.go
