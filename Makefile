run: main.go bindata.go
	go run main.go bindata.go

bindata.go: clean
	go-bindata -nomemcopy=true -ignore=.DS_Store resources/...

clean:
	rm bindata.go
