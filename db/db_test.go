package db

import (
	"testing"
)

func TestNewBadger(t *testing.T) {
	storageDir := "/tmp/badger"
	badger := NewBadger(storageDir)
	//assert.Equal(t, storageDir, badger.opts.Dir, "new badger should return Badger with values, opts containaing storageDir")

}
