package model

import (
	"sync"
)

type SyncMap struct {
	sync.Map
}

func (e *SyncMap) ToMap() map[interface{}]interface{} {
	res := map[interface{}]interface{}{}
	e.ForEach(func(key, value interface{}) {
		res[key] = value
	})
	return res
}

func (e *SyncMap) ToStrMap() map[string]interface{} {
	res := map[string]interface{}{}
	e.ForEach(func(key, value interface{}) {
		res[key.(string)] = value
	})
	return res
}

func (e *SyncMap) Keys() []interface{} {
	var res []interface{}
	e.ForEach(func(key, _ interface{}) {
		res = append(res, key)
	})
	return res
}

func (e *SyncMap) Values() []interface{} {
	res := []interface{}{}
	e.ForEach(func(_, value interface{}) {
		res = append(res, value)
	})
	return res
}

func (e *SyncMap) ValuesMap() []map[string]interface{} {
	res := []map[string]interface{}{}
	e.ForEach(func(_, value interface{}) {
		res = append(res, value.(map[string]interface{}))
	})
	return res
}

func (e *SyncMap) ToArr(f func(key, value interface{}) interface{}) []interface{} {
	res := []interface{}{}
	e.Range(func(key, value interface{}) bool {
		res = append(res, f(key, value))
		return true
	})
	return res
}

func (e *SyncMap) ForEach(f func(key, value interface{})) {
	e.Range(func(key, value interface{}) bool {
		f(key, value)
		return true
	})
}

func (e *SyncMap) ForEachMap(f func(key string, value map[string]interface{})) {
	e.Range(func(key, value interface{}) bool {
		f(key.(string), value.(map[string]interface{}))
		return true
	})
}

func (e *SyncMap) ForEachArr(f func(key string, value []interface{})) {
	e.Range(func(key, value interface{}) bool {
		f(key.(string), value.([]interface{}))
		return true
	})
}

func (e *SyncMap) ForEachArrMap(f func(key string, value []map[string]interface{})) {
	e.Range(func(key, value interface{}) bool {
		f(key.(string), value.([]map[string]interface{}))
		return true
	})
}

func (e *SyncMap) StoreStr(key interface{}, value string) {
	e.Store(key, value)
}

func (e *SyncMap) LoadStr(key interface{}) (string, bool) {
	val, ok := e.Load(key)
	if ok {
		return val.(string), ok
	}
	return "", ok
}

func (e *SyncMap) GetStr(key interface{}) string {
	val, ok := e.Load(key)
	if ok {
		return val.(string)
	}
	return ""
}

func (e *SyncMap) StoreInt(key interface{}, value int) {
	e.Store(key, value)
}

func (e *SyncMap) LoadInt(key interface{}) (int, bool) {
	val, ok := e.Load(key)
	if ok {
		return val.(int), ok
	}
	return 0, ok
}

func (e *SyncMap) StoreInt64(key interface{}, value int64) {
	e.Store(key, value)
}

func (e *SyncMap) LoadInt64(key interface{}) (int64, bool) {
	val, ok := e.Load(key)
	if ok {
		return val.(int64), ok
	}
	return 0, ok
}

func (e *SyncMap) Size() int {
	count := 0
	e.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

func (e *SyncMap) LoadMap(key interface{}) (map[string]interface{}, bool) {
	val, ok := e.Load(key)
	if ok {
		return val.(map[string]interface{}), ok
	}
	return map[string]interface{}{}, ok
}
