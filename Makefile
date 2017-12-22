# Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

GO_LDFLAGS:=--ldflags '-w -s -extldflags "-static"'
GO:=docker run --rm -v $(shell go env GOPATH):/go -w /go/src/$(shell go list) golang:1.9.2-alpine3.6 go

default: generate gen-pkgdoc
	@echo "ok"

build-docker: Dockerfile
	docker build -t chai2010/qingcloud-go .
	@docker image prune -f 1>/dev/null 2>&1

run-docker:
	docker run --rm -it -v `pwd`:/root -w /root chai2010/qingcloud-go qcli

init-vendor:
	govendor init
	govendor add +external
	@echo "ok"

update-vendor:
	govendor update +external
	govendor list
	@echo "ok"

remove-unused-vendor:
	govendor remove +u
	@echo "ok"

tools:
	go get go get github.com/golang/protobuf/protoc-gen-go
	go get github.com/devnev/godoc2ghmd
	docker pull golang:alpine
	@echo "ok"

generate:
	cd ./api && make
	go generate ./pkg/...
	go fmt  ./...
	go test ./...

gen-pkgdoc:
	godoc2ghmd github.com/chai2010/qingcloud-go/pkg/client    > ./pkg/client/README.md
	godoc2ghmd github.com/chai2010/qingcloud-go/pkg/pbutil    > ./pkg/pbutil/README.md
	godoc2ghmd github.com/chai2010/qingcloud-go/pkg/signature > ./pkg/signature/README.md
	godoc2ghmd github.com/chai2010/qingcloud-go/pkg/status    > ./pkg/status/README.md
	godoc2ghmd github.com/chai2010/qingcloud-go/pkg/version   > ./pkg/version/README.md
	godoc2ghmd github.com/chai2010/qingcloud-go/pkg/wait      > ./pkg/wait/README.md

	godoc2ghmd github.com/chai2010/qingcloud-go/pkg/cmd/qcli > ./pkg/cmd/qcli/README.md

	godoc2ghmd github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud > ./pkg/cmd/protoc-gen-qingcloud/README.md
	godoc2ghmd github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud/utils > ./pkg/cmd/protoc-gen-qingcloud/utils/README.md

lake-graph:
	go install ./cmd/qcli
	qcli make -g > lakefile-graph.dot
	qcli make -g | dot -Tpng > lakefile-graph.png

fmt:
	go fmt ./...

test:
	go generate ./pkg/...
	go fmt  ./...
	go vet ./...
	go test ./...

dist:
	go generate ./pkg/...

	mkdir -p ./_dist/qingcloud-cli-windows-amd64
	mkdir -p ./_dist/qingcloud-cli-linux-amd64
	mkdir -p ./_dist/qingcloud-cli-darwin-amd64

	GOOS=windows GOARCH=amd64 go build -o ./_dist/qingcloud-cli-windows-amd64/qcli.exe ./cmd/qcli
	GOOS=linux   GOARCH=amd64 go build -o ./_dist/qingcloud-cli-linux-amd64/qcli       ./cmd/qcli
	GOOS=darwin  GOARCH=amd64 go build -o ./_dist/qingcloud-cli-darwin-amd64/qcli      ./cmd/qcli

	GOOS=windows GOARCH=amd64 go build -o ./_dist/qingcloud-cli-windows-amd64/qlua.exe ./cmd/qlua
	GOOS=linux   GOARCH=amd64 go build -o ./_dist/qingcloud-cli-linux-amd64/qlua       ./cmd/qlua
	GOOS=darwin  GOARCH=amd64 go build -o ./_dist/qingcloud-cli-darwin-amd64/qlua      ./cmd/qlua

	cp cmd/qlake/qlake.lua ./_dist/qingcloud-cli-windows-amd64/
	cp cmd/qlake/qlake.lua ./_dist/qingcloud-cli-linux-amd64/
	cp cmd/qlake/qlake.lua ./_dist/qingcloud-cli-darwin-amd64/

	cp README.md ./_dist/qingcloud-cli-windows-amd64/
	cp README.md ./_dist/qingcloud-cli-linux-amd64/
	cp README.md ./_dist/qingcloud-cli-darwin-amd64/

	cp CHANGES.md ./_dist/qingcloud-cli-windows-amd64/
	cp CHANGES.md ./_dist/qingcloud-cli-linux-amd64/
	cp CHANGES.md ./_dist/qingcloud-cli-darwin-amd64/

	cp LICENSE ./_dist/qingcloud-cli-windows-amd64/
	cp LICENSE ./_dist/qingcloud-cli-linux-amd64/
	cp LICENSE ./_dist/qingcloud-cli-darwin-amd64/

	cd ./_dist && zip -r qingcloud-cli-windows-amd64.zip qingcloud-cli-windows-amd64
	cd ./_dist && zip -r qingcloud-cli-linux-amd64.zip   qingcloud-cli-linux-amd64
	cd ./_dist && zip -r qingcloud-cli-darwin-amd64.zip  qingcloud-cli-darwin-amd64

	@echo "ok"

clean:
	make -C ./pkg/api clean
	make -C ./pkg/cmd/qcli/api clean
