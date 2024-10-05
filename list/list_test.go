package list

import (
	"math/rand"
	"testing"
)

func TestChanAll(t *testing.T) {
	var n *Node[int]
	for range 10 {
		lprepend(&n, rand.Intn(100))
	}

	for v := range ChanAll(n) {
		t.Logf("%v", v)
	}
}

func TestList(t *testing.T) {
	var n *Node[int]

	for range 10 {
		lprepend(&n, rand.Intn(100))
	}

	t.Logf("%s", n)
}

func TestZip(t *testing.T) {
	var n [3]*Node[int]

	for range 10 {
		lprepend(&n[0], rand.Intn(100))
		lprepend(&n[1], rand.Intn(100))
		lprepend(&n[2], rand.Intn(100))
	}

	for _, cur := range n {
		t.Logf("%s", cur)
	}

	for v := range Zip(n[:]...) {
		t.Logf("%v", v)
	}
}
