package utils

import (
	"errors"
	"testing"
	"time"
)

func TestWaitForSpecificOrError(t *testing.T) {

	waitInterval := 100 * time.Millisecond
	timeout := 10*waitInterval + waitInterval/2
	times := 0
	err := WaitForSpecificOrError(func() (bool, error) {
		times++
		println("times", times)
		if times == 3 {
			return true, nil
		}
		return false, nil
	}, timeout, waitInterval)
	tAssert(t, err == nil)
	tAssert(t, 3 == times)

	times = 0
	err = WaitForSpecificOrError(func() (bool, error) {
		times++
		println("times", times)
		if times == 3 {
			return false, errors.New("error")
		}
		return false, nil
	}, timeout, waitInterval)
	tAssert(t, err != nil)
	tAssert(t, 3 == times)

	times = 0
	err = WaitForSpecificOrError(func() (bool, error) {
		times++
		println("times", times)
		return false, nil
	}, timeout, waitInterval)
	tAssert(t, err != nil)
	tErr, ok := err.(*TimeoutError)
	tAssert(t, ok)
	tAssert(t, timeout == tErr.timeout)
	tAssert(t, 10 == times)
}
