package bloomst

import (
	"github.com/boltdb/bolt"
)

type Storage interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, val []byte) error
}

var (
	DefaultBoltBucketName = []byte{0x13}
)

type BoltStorage struct {
	bdb        *bolt.DB
	BucketName []byte
}

func NewBoltStorage(path string) (*BoltStorage, error) {
	bdb, e := bolt.Open(path, 0777, nil)
	if e != nil {
		return nil, e
	}
	e = bdb.Update(func(tx *bolt.Tx) error {
		_, e := tx.CreateBucketIfNotExists(DefaultBoltBucketName)
		return e
	})
	if e != nil {
		return nil, e
	}

	s := new(BoltStorage)
	s.BucketName = DefaultBoltBucketName
	s.bdb = bdb
	return s, nil
}

func (b *BoltStorage) Get(key []byte) ([]byte, error) {
	var (
		res []byte
	)
	e := b.bdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(b.BucketName)
		res = b.Get(key)
		return nil
	})
	return res, e
}
func (b *BoltStorage) Set(key []byte, val []byte) error {
	e := b.bdb.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(b.BucketName)
		e := b.Put(key, val)
		return e
	})
	return e
}
