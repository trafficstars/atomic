package atomic

import (
	"sync/atomic"
	"unsafe"
)

type Uint64Map interface {
	SetMap(input *map[uint64]interface{})
	Get(key uint64) interface{}
}

type uint64Map struct {
	data *map[uint64]interface{}
}

func NewUint64Map() Uint64Map {
	return &uint64Map{
		data: &map[uint64]interface{}{},
	}
}

func (m *uint64Map) load() *map[uint64]interface{} {
	return (*map[uint64]interface{})(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&m.data))))
}

func (m *uint64Map) Get(key uint64) interface{} {
	loadedMap := m.load()
	return (*loadedMap)[key]
}

func (m *uint64Map) SetMap(input *map[uint64]interface{}) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&m.data)), unsafe.Pointer(input))
}
