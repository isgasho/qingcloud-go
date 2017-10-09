# Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

default:
	go fmt ./...
	go test ./...

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
	cd ./spec.pb && make
	go fmt  ./...
	go test ./...

fmt:
	go fmt ./...

test:
	go fmt  ./...
	go test ./...

clean:
