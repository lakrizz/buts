package buts

import (
	"testing"
	"time"
)

func TestInitCorrect(t *testing.T) {
	_, err := NewBoundedTimeoutStack(1*time.Minute, 5)
	if err != nil {
		t.Error(err)
	}
}

func TestInitWithZeroTimeout(t *testing.T) {
	_, err := NewBoundedTimeoutStack(0, 5)
	if err == nil {
		t.Error(err)
	}
}

func TestInitWithZeroCapacity(t *testing.T) {
	_, err := NewBoundedTimeoutStack(1*time.Minute, 0)
	if err == nil {
		t.Error(err)
	}
}

func TestCapMap(t *testing.T) {
	cap := 5
	b, err := NewBoundedTimeoutStack(1*time.Minute, cap)
	if err != nil {
		t.Error(err)
	}

	items := []int{1, 2, 3, 4, 5, 6}
	for _, v := range items {
		err := b.Push(v)
		if err != nil {
			t.Error(err)
		}
	}

	if len(b.GetItemsMap()) > cap {
		t.Error("stack is over capacity")
	}
}

func TestCapSlice(t *testing.T) {
	cap := 5
	b, err := NewBoundedTimeoutStack(1*time.Minute, cap)
	if err != nil {
		t.Error(err)
	}

	items := []int{1, 2, 3, 4, 5, 6}
	for _, v := range items {
		err := b.Push(v)
		if err != nil {
			t.Error(err)
		}
	}

	if len(b.GetItemsSlice()) > cap {
		t.Error("stack is over capacity")
	}
}

func TestDuplicate(t *testing.T) {
	b, err := NewBoundedTimeoutStack(1*time.Minute, 5)
	if err != nil {
		t.Error(err)
	}

	err = b.Push(5)
	if err != nil {
		t.Error(err)
	}

	err = b.Push(5)
	if err == nil { // this should throw an error because the 5 is already in, but don't panic
		t.Error(err)
	}
}

func TestDuplicatePushedOut(t *testing.T) {
	b, err := NewBoundedTimeoutStack(1*time.Minute, 5)
	if err != nil {
		t.Error(err)
	}

	err = b.Push(5)
	if err != nil {
		t.Error(err)
	}

	// now add 5 more things and readd the 5
	items := []int{9, 8, 7, 6, 10}
	for _, v := range items {
		err := b.Push(v)
		if err != nil {
			t.Error(err)
		}
	}

	err = b.Push(5)
	if err != nil { // this should throw an error because the 5 is already in, but don't panic
		t.Error(err)
	}
}

func TestTimeoutPush(t *testing.T) {
	dur := 10 * time.Millisecond
	b, err := NewBoundedTimeoutStack(dur, 5)
	if err != nil {
		t.Error(err)
	}

	err = b.Push(1)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(2 * dur)
	err = b.Push(1)
	if err != nil {
		t.Error(err)
	}
}

func TestTimeoutEmpty(t *testing.T) {
	dur := 10 * time.Millisecond
	b, err := NewBoundedTimeoutStack(dur, 5)
	if err != nil {
		t.Error(err)
	}

	err = b.Push(1)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(2 * dur)
	if len(b.GetItemsSlice()) != 0 {
		t.Error("stack should be empty")
	}
}

func TestShouldContain(t *testing.T) {
	b, err := NewBoundedTimeoutStack(1*time.Minute, 5)
	if err != nil {
		t.Error(err)
	}

	b.Push(1)
	if !b.Contains(1) {
		t.Error("stack should contain item but doesn't")
	}
}

func TestShouldNotContain(t *testing.T) {
	b, err := NewBoundedTimeoutStack(1*time.Minute, 5)
	if err != nil {
		t.Error(err)
	}

	b.Push(1)
	if b.Contains(2) {
		t.Error("stack should not contain item but does")
	}
}

func TestShouldNotContainAnymore(t *testing.T) {
	b, err := NewBoundedTimeoutStack(1*time.Minute, 5)
	if err != nil {
		t.Error(err)
	}

	items := []int{9, 8, 7, 6, 10, 1}
	for _, v := range items {
		err := b.Push(v)
		if err != nil {
			t.Error(err)
		}
	}

	if b.Contains(9) {
		t.Error("stack should not contain item but does")
	}
}
