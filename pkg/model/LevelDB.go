package model

import (
	"encoding/json"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/xiaojun207/go-base-utils/math"
	"log"
)

type LevelDB struct {
	*leveldb.DB
	Path string
}

func NewLevelDB(path string) LevelDB {
	db := LevelDB{
		Path: path,
	}
	db.Init()
	return db
}

func (e *LevelDB) Init() {
	var err error
	e.DB, err = leveldb.OpenFile(e.Path, nil)
	if err != nil {
		log.Println(err)
	}
}

func (e *LevelDB) Close() {
	e.Close()
}

func (e *LevelDB) put(key string, value []byte) error {
	if value == nil {
		return errors.New("LevelDB put value is nil")
	}
	return e.Put([]byte(key), value, nil)
}

func (e LevelDB) get(key string) ([]byte, error) {
	return e.Get([]byte(key), nil)
}

func (e *LevelDB) Remove(key string) error {
	return e.Delete([]byte(key), nil)
}

func (e LevelDB) Has(key string) bool {
	has, err := e.DB.Has([]byte(key), nil)
	if err != nil {
		log.Println("Has", err)
		log.Panic("Has", err)
	}
	return has
}

func (e *LevelDB) Store(key string, value interface{}) error {
	byteValue, err := json.Marshal(value)
	if err != nil {
		log.Println("Store.err:", err)
		return err
	}
	return e.put(key, byteValue)
}

func (e *LevelDB) Load(key string) ([]byte, error) {
	return e.get(key)
}

func (e *LevelDB) ToMap() map[interface{}]interface{} {
	res := map[interface{}]interface{}{}
	e.ForEach(func(key string, value []byte) {
		var tmp interface{}
		json.Unmarshal(value, &tmp)
		res[key] = value
	})
	return res
}

func (e *LevelDB) ToStrMap() map[string]interface{} {
	res := map[string]interface{}{}
	e.ForEach(func(key string, value []byte) {
		var tmp interface{}
		json.Unmarshal(value, &tmp)
		res[key] = tmp
	})
	return res
}

func (e *LevelDB) Keys() []string {
	var res []string
	e.ForEach(func(key string, _ []byte) {
		res = append(res, key)
	})
	return res
}

func (e *LevelDB) Values() []interface{} {
	res := []interface{}{}
	e.ForEach(func(_ string, value []byte) {
		res = append(res, value)
	})
	return res
}

func (e *LevelDB) ValuesMap() []map[string]interface{} {
	res := []map[string]interface{}{}
	e.ForEach(func(_ string, value []byte) {
		var tmp map[string]interface{}
		json.Unmarshal(value, &tmp)
		res = append(res, tmp)
	})
	return res
}

func (e *LevelDB) Range(f func(key string, value []byte) bool) {
	iter := e.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		if !f(string(key), value) {
			return
		}
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Println(err)
	}
}

func (e *LevelDB) ToArr(f func(key, value interface{}) interface{}) []interface{} {
	res := []interface{}{}
	e.Range(func(key string, value []byte) bool {
		res = append(res, f(key, value))
		return true
	})
	return res
}

func (e *LevelDB) ForEach(f func(key string, value []byte)) {
	e.Range(func(key string, value []byte) bool {
		f(key, value)
		return true
	})
}

func (e *LevelDB) ForEachMap(f func(key string, value map[string]interface{})) {
	e.Range(func(key string, value []byte) bool {
		var res map[string]interface{}
		json.Unmarshal(value, &res)
		f(key, res)
		return true
	})
}

func (e *LevelDB) ForEachArr(f func(key string, value []interface{})) {
	e.Range(func(key string, value []byte) bool {
		var res []interface{}
		json.Unmarshal(value, &res)
		f(key, res)
		return true
	})
}

func (e *LevelDB) ForEachArrMap(f func(key string, value []map[string]interface{})) {
	e.Range(func(key string, value []byte) bool {
		var res []map[string]interface{}
		json.Unmarshal(value, &res)
		f(key, res)
		return true
	})
}

func (e *LevelDB) StoreStr(key string, value string) {
	e.put(key, []byte(value))
}

func (e *LevelDB) LoadStr(key string) (string, bool) {
	val, err := e.Load(key)
	if err != nil {
		return "", false
	}
	return string(val), true
}

func (e *LevelDB) GetStr(key string) string {
	val, _ := e.LoadStr(key)
	return val
}

func (e *LevelDB) StoreInt(key string, value int) {
	e.Store(key, value)
}

func (e *LevelDB) LoadInt(key string) int {
	res, err := e.Load(key)
	if err != nil {
		log.Println("LoadInt.err:", err)
		return 0
	}
	return math.Byte2Int(res)
}

func (e *LevelDB) StoreInt64(key string, value int64) {
	e.Store(key, value)
}

func (e *LevelDB) LoadInt64(key string) int64 {
	res, err := e.Load(key)
	if err != nil {
		log.Println("LoadInt.err:", err)
		return 0
	}
	return int64(math.ByteToUint64(res))
}

func (e *LevelDB) Size() int {
	count := 0
	e.Range(func(key string, value []byte) bool {
		count++
		return true
	})
	return count
}

func (e *LevelDB) LoadMap(key string) map[string]interface{} {
	val, err := e.Load(key)
	if err != nil {
		return map[string]interface{}{}
	}
	var res map[string]interface{}
	json.Unmarshal(val, &res)
	return res
}
