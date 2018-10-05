package atomic

import (
	"fmt"
	"sync"
	"testing"
)

const (
	concurrency = 4
	itemsAmount = 100
)

func BenchmarkSimpleStringMapGet(b *testing.B) {
	var wg sync.WaitGroup

	m := map[string]interface{}{}

	for i := 0; i < itemsAmount; i++ {
		m[fmt.Sprintf("%d", i)] = i
	}

	b.ResetTimer()
	for c := 0; c < concurrency; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < b.N; n++ {
				_ = m["0"]
			}
		}()
	}
	wg.Wait()
}

func BenchmarkAtomicStringMapGet(b *testing.B) {
	var wg sync.WaitGroup

	values := map[string]interface{}{}
	m := NewStringMap()

	for i := 0; i < itemsAmount; i++ {
		values[fmt.Sprintf("%d", i)] = i
	}
	m.SetMap(&values)

	b.ResetTimer()
	for c := 0; c < concurrency; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < b.N; n++ {
				_ = m.Get("0")
			}
		}()
	}
	wg.Wait()
}

func BenchmarkSimpleUint64MapGet(b *testing.B) {
	var wg sync.WaitGroup

	m := map[uint64]interface{}{}

	for i := uint64(0); i < itemsAmount; i++ {
		m[i] = i
	}

	b.ResetTimer()
	for c := 0; c < concurrency; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < b.N; n++ {
				_ = m[0]
			}
		}()
	}
	wg.Wait()
}

func BenchmarkAtomicUint64MapGet(b *testing.B) {
	var wg sync.WaitGroup

	values := map[uint64]interface{}{}
	m := NewUint64Map()

	for i := uint64(0); i < itemsAmount; i++ {
		values[i] = i
	}
	m.SetMap(&values)

	b.ResetTimer()
	for c := 0; c < concurrency; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < b.N; n++ {
				_ = m.Get(0)
			}
		}()
	}
	wg.Wait()
}

func BenchmarkSimpleUint64MapSetMap(b *testing.B) {
	var wg sync.WaitGroup

	b.ResetTimer()
	for c := 0; c < concurrency; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for n := 0; n < b.N; n++ {
				values := map[uint64]interface{}{}
				for i := uint64(0); i < itemsAmount; i++ {
					values[i] = i
				}
			}
		}()
	}
	wg.Wait()
}

func BenchmarkAtomicUint64MapSetMap(b *testing.B) {
	var wg sync.WaitGroup

	m := NewUint64Map()

	b.ResetTimer()
	for c := 0; c < concurrency; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for n := 0; n < b.N; n++ {
				values := map[uint64]interface{}{}
				for i := uint64(0); i < itemsAmount; i++ {
					values[i] = i
				}
				m.SetMap(&values)
			}
		}()
	}
	wg.Wait()
}
