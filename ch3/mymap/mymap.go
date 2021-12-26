package mymap

import "sync"

// 封装一个数据结构 MyMap，实现并发安全的 Load，Store，Delete，LoadAndDelete，LoadOrStore
// 几个 API(禁止使用 sync.Map)，不用考虑性能

type MyMap struct {
	sync.Mutex

	m map[interface{}]interface{}
}

func NewMyMap() *MyMap {
	return &MyMap{
		m: make(map[interface{}]interface{}),
	}
}

func (my *MyMap) Load(key interface{}) (value interface{}, ok bool) {
	my.Lock()
	defer my.Unlock()
	value, ok = my.m[key]
	return
}

func (my *MyMap) Delete(key interface{}) {
	my.Lock()
	defer my.Unlock()

	delete(my.m, key)
}

func (my *MyMap) LoadAndDelete(key interface{}) (value interface{}, loaded bool) {
	my.Lock()
	defer my.Unlock()

	if value, loaded = my.m[key]; loaded {
		delete(my.m, key)
	}
	return
}

func (my *MyMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	my.Lock()
	defer my.Unlock()

	if actual, loaded = my.m[key]; loaded {
		my.m[key] = value
	}

	return
}

func (my *MyMap) Store(key, value interface{}) {
	my.Lock()
	defer my.Unlock()

	my.m[key] = value
}
