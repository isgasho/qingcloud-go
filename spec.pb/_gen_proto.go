// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// https://github.com/google/protobuf/releases
// go install github.com/chai2010/qingcloud-go/protoc-gen-go

//go:generate protoc --go_out=plugins=qingcloud:. base.proto instances.proto nic.proto snapshot.proto zone.proto cache.proto             job.proto               rdb.proto               tag.proto dns.proto               key_pair.proto          router.proto            user_data.proto eip.proto               load_balancer.proto     security_group.proto    volume.proto image.proto             mongo.proto             shared_storage.proto    vxnet.proto

package spec
