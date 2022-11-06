clean:
	rm docker/listapi || true
run:
	go build -o docker/listapi cmd/listapi/main.go
	./docker/listapi