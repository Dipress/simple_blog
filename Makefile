SHELL := /bin/sh

dev:
	docker-compose -f kit/docker/docker-compose.yaml up

test:
	go test -v -race `go list ./...`

cover: 
	go test --race `go list ./... | grep -v /vendor | grep -v /cmd/blog ` -coverprofile cover.out.tmp && \
	cat cover.out.tmp > cover.out && \
	go tool cover -func cover.out && \
	rm cover.out.tmp && \
	rm cover.out