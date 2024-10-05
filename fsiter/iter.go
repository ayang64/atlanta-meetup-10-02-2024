package fsiter

import (
	"fmt"
	"io/fs"
)

func All(fsys fs.FS, root string) func(func(string, fs.DirEntry) bool) {
	return func(yield func(string, fs.DirEntry) bool) {
		fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !yield(path, d) {
				return fmt.Errorf("iteration stopped")
			}
			return nil
		})
	}
}
