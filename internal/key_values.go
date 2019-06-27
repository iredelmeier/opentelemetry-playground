package internal

import "sync"

type Entry struct {
	Key   string
	Value string
}

type KeyValues struct {
	lock *sync.RWMutex
	kv   map[string]string
}

func (kv *KeyValues) Get(key string) (string, bool) {
	kv.lock.RLock()
	defer kv.lock.RUnlock()

	value, ok := kv.kv[key]

	return value, ok
}

func (kv *KeyValues) Set(key string, value string) {
	kv.lock.Lock()
	defer kv.lock.Unlock()

	kv.kv[key] = value
}

func (kv *KeyValues) Entries() []Entry {
	kv.lock.RLock()
	defer kv.lock.RUnlock()

	var entries []Entry

	for k, v := range kv.kv {
		entries = append(entries, Entry{
			Key:   k,
			Value: v,
		})
	}

	return entries
}
