package internal

import "sync"

type Attribute struct {
	Key   string
	Value string
}

type Attributes struct {
	lock       *sync.RWMutex
	attributes map[string]string
}

func NewAttributes() Attributes {
	return Attributes{
		lock:       &sync.RWMutex{},
		attributes: make(map[string]string),
	}
}

func (a Attributes) Get(key string) (string, bool) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	value, ok := a.attributes[key]

	return value, ok
}

func (a Attributes) Set(key string, value string) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.attributes[key] = value
}

func (a Attributes) Entries() []Attribute {
	a.lock.RLock()
	defer a.lock.RUnlock()

	var entries []Attribute

	for k, v := range a.attributes {
		entries = append(entries, Attribute{
			Key:   k,
			Value: v,
		})
	}

	return entries
}
