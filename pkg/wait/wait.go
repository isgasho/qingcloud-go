// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package wait

import (
	"fmt"
	"time"
)

func WaitForIntervalWorkDone(name string, timeout time.Duration, intervalWork func() (done bool, err error)) error {
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		if done, err := intervalWork(); done || err != nil {
			if time.Now().After(deadline) {
				return fmt.Errorf("pkg/wait: wait %s failed after %v", name, timeout)
			}
			return err
		}
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("pkg/wait: wait %s failed after %v", name, timeout)
}

func WaitForOnetimeWork(name string, timeout time.Duration, onetimeWork func() error) (err error) {
	done := make(chan struct{})
	go func() {
		err = onetimeWork()
		close(done)
	}()

	select {
	case <-time.After(timeout):
		return fmt.Errorf("pkg/wait: wait %s failed after %v", name, timeout)
	case <-done:
		return err
	}
}
