package bloomst

import (
	"github.com/willf/bloom"
)

type Bloomst struct {
	storage Storage
}

func New(fpath string) (*Bloomst, error) {
	s, e := NewBoltStorage(fpath)
	if e != nil {
		return nil, e
	}

	b := new(Bloomst)
	b.storage = s

	return b, nil
}

func (s *Bloomst) Test(key []byte, data []byte) (bool, error) {
	bf, e := s.newBf(key)
	if e != nil {
		return false, e
	}
	res := bf.Test(data)
	return res, nil
}

func (s *Bloomst) TestAndAdd(key []byte, data []byte) (bool, error) {
	bf, e := s.newBf(key)
	if e != nil {
		return false, e
	}
	res := bf.TestAndAdd(data)
	e = s.saveBf(key, bf)
	return res, e
}

func (s *Bloomst) saveBf(key []byte, bf *bloom.BloomFilter) error {
	bts, e := bf.GobEncode()
	if e != nil {
		return e
	}
	return s.storage.Set(key, bts)
}

func (s *Bloomst) newBf(key []byte) (*bloom.BloomFilter, error) {
	bts, e := s.storage.Get(key)
	if e != nil {
		return nil, e
	}
	bf := new(bloom.BloomFilter)
	if bts == nil {
		return bloom.New(20000, 5), nil
	}
	e = bf.GobDecode(bts)
	if e != nil {
		return nil, e
	}
	return bf, nil
}
