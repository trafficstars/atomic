package atomic

import (
	"sync/atomic"
	"unsafe"
)

type StringMap interface {
	SetMap(input *map[string]interface{})
	Get(key string) interface{}
}

type stringMap struct {
	data *map[string]interface{}
}

func NewStringMap() StringMap {
	return &stringMap{
		data: &map[string]interface{}{},
	}
}

func (m *stringMap) load() *map[string]interface{} {
	return (*map[string]interface{})(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&m.data))))
}

func (m *stringMap) Get(key string) interface{} {
	loadedMap := m.load()
	return (*loadedMap)[key]
}

func (m *stringMap) SetMap(input *map[string]interface{}) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&m.data)), unsafe.Pointer(input))
}
