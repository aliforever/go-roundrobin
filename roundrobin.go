package roundrobin

import (
	"errors"
	"sync"
)

type RoundRobin struct {
	m       sync.Mutex
	counter int
	items   []string
}

func New() *RoundRobin {
	return &RoundRobin{}
}

func (r *RoundRobin) Add(items ...string) {
	r.m.Lock()
	defer r.m.Unlock()

	for _, item := range items {
		for _, storedItem := range r.items {
			if item == storedItem {
				return
			}
		}
	}

	r.items = append(r.items, items...)
}

func (r *RoundRobin) RemoveItem(item string) {
	r.m.Lock()
	defer r.m.Unlock()

	newItems := []string{}

	for _, i := range r.items {
		if i != item {
			newItems = append(newItems, i)
		}
	}

	r.items = newItems
}

func (r *RoundRobin) Next() (item string, err error) {
	r.m.Lock()
	defer r.m.Unlock()

	if len(r.items) == 0 {
		err = errors.New("no_items_available")
		return
	}

	if r.counter >= len(r.items) {
		r.counter = 0
	}

	item = r.items[r.counter]

	r.counter++

	return
}

func (r *RoundRobin) Items() []string {
	r.m.Lock()
	defer r.m.Unlock()

	return r.items
}
