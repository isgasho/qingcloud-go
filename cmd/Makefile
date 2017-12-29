# Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

default:

init-vendor:
	govendor init
	govendor add +external
	govendor remove ./vendor/github.com/chai2010/qingcloud-go/...
	mkdir -p ./vendor/github.com/chai2010/qingcloud-go
	ln -s ../../../../../pkg ./vendor/github.com/chai2010/qingcloud-go/pkg
	@echo "ok"

update-vendor:
	govendor update +external
	govendor remove ./vendor/github.com/chai2010/qingcloud-go/...
	mkdir -p ./vendor/github.com/chai2010/qingcloud-go
	ln -s ../../../../../pkg ./vendor/github.com/chai2010/qingcloud-go/pkg
	govendor list
	@echo "ok"

remove-unused-vendor:
	govendor remove +u
	@echo "ok"

clean:
	-rm -rf ./vendor
