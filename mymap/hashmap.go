package mymap

import "hash/crc32"

// Unordered Map
// ㄴ Hash Map
// Ordered Map
// ㄴ Sorted Map

const arraySize = 1

type hashData[T any] struct {
	key   string
	value T
}

type HashMap[T any] struct {
	arr [arraySize][]hashData[T]
}

func hashfn(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (h *HashMap[T]) Add(key string, value T) {
	hash := hashfn(key)
	var hashData = hashData[T]{
		key:   key,
		value: value,
	}

	h.arr[hash%arraySize] = append(h.arr[hash%arraySize], hashData)
}

func (h *HashMap[T]) Get(key string) (T, bool) {
	hash := hashfn(key)

	for _, hd := range h.arr[hash%arraySize] {
		if hd.key == key {
			return hd.value, true
		}
	}

	var tempT T
	return tempT, false
}
