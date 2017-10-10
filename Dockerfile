# Copyright 2017 ChaiShushan<chaishushan@gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

FROM scratch
ADD qingcloud-cli.linux.exe /
ENTRYPOINT ["/qingcloud-cli.linux.exe"]
CMD []
