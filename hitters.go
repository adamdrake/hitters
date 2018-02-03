package hitters

import (
	"errors"
	"sync"
)

/*

This is an implementation of _Frequent_ from Misra and Gries (1982).
It will take a list (or stream) consisting of n items, and find all items
which appears with frequency > n/k where k is the length of the data stream,
using only k - 1 counters.  

Manku and Motwani's LossyCounting (2002) and Metwally's SpaceSaving (2005)
algorithms are similar, situationally more accurate, but not always needed.

*/

type Hitters struct {
	items                    map[string]int
	capacity, processedCount int
	sync.RWMutex
}

// New takes a number k representing how many of the top items to count
func New(k int) (*Hitters, error) {
	if k < 1 {
		return &Hitters{}, errors.New("Constructor requires a positive argument")
	}
	return &Hitters{capacity: k, items: make(map[string]int)}, nil
}

func (t *Hitters) addOne(item string) {
	t.Lock()
	defer t.Unlock()
	t.processedCount++

	for k := range t.items {
		if k == item {
			t.items[k]++
			return
		}
	}

	// Item is not present in the current map.  First, check if current size is less than max and if so, add item.
	if len(t.items) < t.capacity-1 {
		t.items[item] = 1
		return
	}

	// Items list is full and current item is not in list, so decrement all and remove any less than 1
	for k := range t.items {
		t.items[k]--
		if t.items[k] < 1 {
			delete(t.items, k)
		}
	}
}

// Add will put an item into the top k list
func (t *Hitters) Add(items ...string) {
	for _, item := range items {
		t.addOne(item)
	}
}

// Get will return the count for the provided key if it exists, and 0 otherwise
func (t *Hitters) Get(k string) int {
	t.RLock()
	defer t.RUnlock()
	return t.items[k]
}

// Items returns all items in the Hitters
func (t *Hitters) Items() map[string]int {
	its := make(map[string]int)
	t.RLock()
	defer t.RUnlock()
	for k, v := range t.items {
		its[k] = v
	}
	return its
}
