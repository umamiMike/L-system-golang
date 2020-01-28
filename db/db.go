package db

import (
	"log"
	"time"

	"github.com/dgraph-io/badger"
)

type Badger struct {
	db *badger.DB
}

type Operation struct {
	Key   string
	Value []byte
	Op    byte
}

func NewBadger(storageDir string) *Badger {
	storage := &Badger{}
	opts := badger.DefaultOptions(storageDir)
	opts.SyncWrites = true
	opts.Dir = storageDir
	opts.ValueDir = storageDir
	var err error
	storage.db, err = badger.Open(opts)
	if err != nil {
		panic(err)
	}

	go storage.runStorageGC()

	return storage
}

func (db *Badger) Write(key string, data string) {
	err := db.Set(key, []byte(data))
	if err != nil {
		log.Fatal(err)
	}

}
func get(key string) []byte {
	db := NewBadger("/tmp/badger")
	val, err := db.Get(key)
	if err != nil {
		log.Fatal(err)
	}
	db.db.Close()
	return val
}

// Del deletes a key
func (storage *Badger) Del(key string) (err error) {
	return storage.db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})
}

// Get returns value by key
func (storage *Badger) Get(key string) (value []byte, err error) {
	err = storage.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(value)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// Set adds a key-value pair to the database
func (storage *Badger) Set(key string, value []byte) (err error) {
	return storage.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), value)
		return err
	})
}

// DbStorage represent base db storage interface
type DbStorage interface {
	Set(key string, value []byte) (err error)
	Del(key string) (err error)
	Get(key string) (value []byte, err error)
	Iterate(fn func(key []byte, value []byte))
	IterateByPrefix(prefix []byte, limit uint64, fn func(key []byte, value []byte)) uint64
	IterateByPrefixFrom(prefix []byte, from []byte, limit uint64, fn func(key []byte, value []byte)) uint64
	DeleteByPrefix(prefix []byte)
	KeysByPrefixCount(prefix []byte) uint64
	ProcessBatch(batch []*Operation) (err error)
	Close() error
}

func (storage *Badger) runStorageGC() {
	timer := time.NewTicker(10 * time.Minute)
	for {
		select {
		case <-timer.C:
			storage.storageGC()
		}
	}
}
func (storage *Badger) storageGC() {

again:
	err := storage.db.RunValueLogGC(0.5)
	if err == nil {
		goto again
	}
}
