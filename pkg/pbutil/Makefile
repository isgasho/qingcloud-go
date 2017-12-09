# Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

default:
	protoc -I. --go_out=. testdata.proto
	mv testdata.pb.go testdata_test.go
	@echo "ok"

clean:
	@echo "ok"
