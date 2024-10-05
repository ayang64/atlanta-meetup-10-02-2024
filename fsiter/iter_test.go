package fsiter

import (
	"io/fs"
	"os"
	"testing"
)

// NOTE: these benchmarks are pretty much useless.  they're just measuring
// filesystem performance and can be affected by caching etc.

func BenchmarkWalkFS(b *testing.B) {
	etc := os.DirFS("./testdata/testdir")
	items := 0
	for range b.N {
		fs.WalkDir(etc, ".", func(p string, d fs.DirEntry, e error) error {
			items++
			return nil
		})
	}
}

func BenchmarkAll(b *testing.B) {
	items := 0
	for range b.N {
		for range All(os.DirFS("./testdata/testdir"), ".") {
			items++
		}
	}
}

func TestIter(t *testing.T) {
	for name := range All(os.DirFS("./testdata/testdir"), ".") {
		t.Logf("%s", name)
	}
}
