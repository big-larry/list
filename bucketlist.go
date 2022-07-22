package list

import "errors"

type BucketList[T any] struct {
	buckets  *LinkedList[[]T]
	capacity uint32
}

func NewBucketList[T any](capacity uint32) *BucketList[T] {
	return &BucketList[T]{buckets: NewLinkedList[[]T](), capacity: capacity}

}

func (l *BucketList[T]) Add(item T) {
	last := l.buckets.GetLast()
	if last == nil || len(last.Value) == int(l.capacity) {
		last = l.buckets.Add(make([]T, 0, l.capacity))
	}
	last.Value = append(last.Value, item)
}

func (l *BucketList[T]) Get(index uint32) (T, error) {
	if l.buckets.last == nil {
		return getZero[T](), errors.New("List is empty")
	}
	bucket_index := index / l.capacity
	if bucket_index > l.buckets.GetLength() {
		return getZero[T](), errors.New("Out of range")
	}
	bucket := l.buckets.Get(bucket_index)
	if bucket == nil {
		return getZero[T](), errors.New("Bucket is empty")
	}
	bucket_item_index := index % l.capacity
	if bucket_item_index > uint32(len(bucket.Value)) {
		return getZero[T](), errors.New("Out of range")
	}
	return bucket.Value[bucket_item_index], nil
}

func getZero[T any]() T {
	var result T
	return result
}
