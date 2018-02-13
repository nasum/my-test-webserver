build: main.go handler/hello_world.go
	dep ensure
	GOOS=linux GOARCH=amd64 go build
	docker build ./ -t nasum/my-test-webserver
