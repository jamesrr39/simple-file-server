
.PHONY: install
install:
	go build -o ${GOBIN}/simple-file-server main.go
