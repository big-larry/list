package list

import "errors"

type BucketListCursor[T any] struct {
	current_bucket       *LinkedListNode[[]T]
	current_bucket_index uint32
}

func NewBucketListCursor[T any](list *BucketList[T]) (*BucketListCursor[T], error) {
	if list == nil || list.buckets.root == nil || len(list.buckets.root.Value) == 0 {
		return nil, errors.New("List is empty")
	}
	return &BucketListCursor[T]{current_bucket: list.buckets.root}, nil
}

func (cursor *BucketListCursor[T]) GetNext() (T, bool) {
	if cursor.current_bucket_index == uint32(len(cursor.current_bucket.Value)) {
		if cursor.current_bucket.next == nil {
			return getZero[T](), false
		}
		cursor.current_bucket = cursor.current_bucket.next
		cursor.current_bucket_index = 0
	}
	cursor.current_bucket_index++
	return cursor.current_bucket.Value[cursor.current_bucket_index-1], true
}
