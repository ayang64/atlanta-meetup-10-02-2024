package list

import (
	"cmp"
	"fmt"
	"iter"
	"strings"
)

type Node[T cmp.Ordered] struct {
	value T
	next  *Node[T]
}

func (cur *Node[T]) String() string {
	var sb strings.Builder
	for v := range All(cur) {
		fmt.Fprintf(&sb, "%v->", v)
	}
	fmt.Fprintf(&sb, "nil")
	return sb.String()
}

func ChanAll[T cmp.Ordered](cur *Node[T]) chan T {
	rc := make(chan T)
	go func() {
		for ; cur != nil; cur = cur.next {
			rc <- cur.value
		}
		close(rc)
	}()
	return rc
}

func lprepend[T cmp.Ordered](cur **Node[T], v T) {
	*cur = &Node[T]{value: v, next: *cur}
}

func lappend[T cmp.Ordered](cur **Node[T], v T) {
	for ; cur != nil; cur = &(*cur).next {
	}
	*cur = &Node[T]{value: v}
}

func Zip[T cmp.Ordered](lists ...*Node[T]) iter.Seq[T] {
	var nexts []func() (T, bool)
	for _, list := range lists {
		n, _ := iter.Pull(All(list))
		nexts = append(nexts, n)
	}

	return func(yield func(v T) bool) {
		for {
			var success int
			for _, next := range nexts {
				v, ok := next()
				if !ok {
					continue
				}
				if !yield(v) {
					return
				}
				success++
			}
			if success == 0 {
				return
			}
		}
	}
}

func All[T cmp.Ordered](head *Node[T]) iter.Seq[T] {
	return func(yield func(v T) bool) {
		for cur := head; cur != nil; cur = cur.next {
			if !yield(cur.value) {
				return
			}
		}
	}
}
