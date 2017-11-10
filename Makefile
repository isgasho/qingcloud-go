# Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

GO_LDFLAGS:=--ldflags '-w -s -extldflags "-static"'
GO:=docker run --rm -v $(shell go env GOPATH):/go -w /go/src/$(shell go list) golang:alpine go

default:
	go fmt ./...
	go test ./...

build-docker: Dockerfile
	$(GO) fmt ./...
	docker build -t qingcloud-cli .
	@docker image prune -f 1>/dev/null 2>&1

run-docker:
	docker run --rm -it -v `pwd`:/root -w /root qingcloud-cli

hello:
	go fmt ./...
	go run hello.go

demo_mgo:
	go run examples/mongo/DescribeMongos/DescribeMongos.go

demo_cluster:
	go run examples/cluster/DescribeClusterNodes/DescribeClusterNodes.go

# go run examples/cluster/DescribeClusters/DescribeClusters.go

demo:
	go run examples/nic/DescribeNics/DescribeNics.go

init-vendor:
	govendor init
	govendor add +external
	@echo "ok"

update-vendor:
	govendor update +external
	govendor list
	@echo "ok"

tools:
	go get github.com/kardianos/govendor
	docker pull golang:alpine
	@echo "ok"

generate:
	cd ./api && make
	go fmt  ./...
	go test ./...

fmt:
	go fmt ./...

test:
	go fmt  ./...
	go test ./...

clean:
