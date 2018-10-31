// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package wait

import (
	"fmt"
	"time"
)

type TimeoutError interface {
	Timeout() bool
	error
}

type errTimeout struct {
	error
}

func NewErrTimeout(err error) error {
	return &errTimeout{err}
}

func (e *errTimeout) Error() string {
	return e.error.Error()
}

func (e *errTimeout) Timeout() bool {
	return true
}

func WaitForIntervalWorkDone(
	name string, timeout time.Duration,
	intervalWork func() (done bool, err error),
) error {
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		if done, err := intervalWork(); done || err != nil {
			return err
		}
		time.Sleep(time.Second << uint(tries))
	}
	return NewErrTimeout(
		fmt.Errorf("pkg/wait: wait %s failed after %v", name, timeout),
	)
}

func WaitForOnetimeWork(
	name string, timeout time.Duration,
	onetimeWork func() error,
) (err error) {
	done := make(chan struct{})
	go func() {
		err = onetimeWork()
		close(done)
	}()

	select {
	case <-time.After(timeout):
		return NewErrTimeout(
			fmt.Errorf("pkg/wait: wait %s failed after %v", name, timeout),
		)
	case <-done:
		return err
	}
}
