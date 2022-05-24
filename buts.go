package buts

import (
	"errors"
	"sort"
	"time"
)

type BoundedUniqueTimeoutStack struct {
	timeout time.Duration
	items   map[any]time.Time
	bounds  int
}

//  NewBoundedTimeoutStack() creates a new instance of a buts with a timeout and a cap
func NewBoundedTimeoutStack(timeout time.Duration, bounds int) (*BoundedUniqueTimeoutStack, error) {
	if timeout == 0 {
		return nil, errors.New("timeout should not be empty, items will be discarded instantly")
	}
	if bounds == 0 {
		return nil, errors.New("bounds should not be empty, items can't be added then")
	}

	return &BoundedUniqueTimeoutStack{timeout: timeout, items: make(map[any]time.Time, 0), bounds: bounds}, nil
}

// GetItemssMap() returns an unordered (by spec) map of valid items and their creation-date
func (buts *BoundedUniqueTimeoutStack) GetItemsMap() map[any]time.Time {
	now := time.Now()
	for k, v := range buts.items {
		if now.After(v) { // this creates the Timeout effect upon reading the stack
			delete(buts.items, k)
		}
	}
	return buts.items
}

// GetItemsSlice() returns an ordered (oldest item first) slice of all valid items in the stack
func (buts *BoundedUniqueTimeoutStack) GetItemsSlice() []any {
	// only collect the keys here
	s := make([]any, 0)
	for k := range buts.GetItemsMap() {
		s = append(s, k)
	}

	// sort it by date
	buts.sortslice(s)

	return s
}

// Push() pushes a new item on the stack and returns an error if it already exists on the stack
func (buts *BoundedUniqueTimeoutStack) Push(item any) error {
	if item == nil {
		return errors.New("can't add nil as an item")
	}

	if _, exists := buts.GetItemsMap()[item]; exists {
		return errors.New("item already exists, this is a unique set")
	}

	buts.items[item] = time.Now().Add(buts.timeout)

	if len(buts.items) >= buts.bounds {
		// limit the base stack to the givens bounds
		order := buts.getOrder()
		for i := buts.bounds; i <= len(buts.items)-1; i++ {
			// delete this item as it's too old/over bounds
			delete(buts.items, order[i])
		}
	}

	return nil
}

// Pop() returns the oldest valid item on the stack or nil if the stack is empty
func (buts *BoundedUniqueTimeoutStack) Pop() any {
	sl := buts.GetItemsSlice()
	if len(sl) == 0 {
		return nil
	}
	pop := sl[0]
	delete(buts.items, pop)

	return pop
}

// Contains() compares the given item with all items on the stack and returns true upon a match or false if it's now part of the stack
func (buts *BoundedUniqueTimeoutStack) Contains(item any) bool {
	for k, _ := range buts.GetItemsMap() {
		if k == item {
			return true
		}
	}
	return false
}

// getOrder() returns the keys of the map in order of their creation time
// the desc parameter controls whether it is ordered ascending or descendingly
func (buts *BoundedUniqueTimeoutStack) getOrder() []any {
	result_set := make([]any, 0)
	for k, _ := range buts.items {
		result_set = append(result_set, k)
	}
	buts.sortslice(result_set)

	return result_set
}

func (buts *BoundedUniqueTimeoutStack) sortslice(slice []any) {
	sort.SliceStable(slice, func(i, j int) bool {
		return buts.items[slice[i]].After(buts.items[slice[j]])
	})
}
