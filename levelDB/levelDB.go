package levelDB

import (
	"encoding/json"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"path/filepath"
)

var server Server

const ConstLeveldbFileName = "db"

type Server interface {
	Open() error
	Close() error
	Instance() *leveldb.DB
	IsExistFromLevelDB(key string) (bool, error)
	FindByPrefix(pre string) (files map[string]string, err error)
	FindAll() (map[string]string, error)
	FindLimit(start, limit string) (data map[string]string, err error)
	BatchInsert(attr map[string]string) error
	RemoveKeyFromLevelDB(key string) error
	Insert(key string, val interface{}) error
	FindByKey(key string) (data []byte, err error)
}

func NewServer(dir string) Server {
	s := new(ServerImpl)
	s.dirPath = dir
	return s
}

func InitServer(dir string) error {
	server = NewServer(dir)
	err := server.Open()
	if err != nil {
		return err
	}
	return nil
}

func GetServer() Server {
	return server
}

type ServerImpl struct {
	ldb     *leveldb.DB
	dirPath string
}

func (s *ServerImpl) Open() error {
	if s.ldb != nil {
		return nil
	}
	_ = &opt.Options{
		CompactionTableSize: 1024 * 1024 * 2,
		WriteBuffer:         1024 * 1024 * 2,
		Filter:              filter.NewBloomFilter(10),
	}
	var err error
	//s.ldb, err = leveldb.OpenFile(filepath.Join(s.dirPath, ConstLeveldbFileName), opts)
	s.ldb, err = leveldb.OpenFile(filepath.Join(s.dirPath, ConstLeveldbFileName), nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerImpl) Close() error {
	err := s.ldb.Close()
	return err
}

func (s *ServerImpl) Instance() *leveldb.DB {
	if s.ldb == nil {
		return nil
	}
	return s.ldb
}

func (s *ServerImpl) RemoveKeyFromLevelDB(key string) error {
	err := s.ldb.Delete([]byte(key), nil)
	return err
}

func (s *ServerImpl) IsExistFromLevelDB(key string) (bool, error) {
	return s.ldb.Has([]byte(key), nil)
}

// attr[key]val
func (s *ServerImpl) BatchInsert(attr map[string]string) error {
	batch := new(leveldb.Batch)
	for k, v := range attr {
		batch.Put([]byte(k), []byte(v))
	}
	err := s.ldb.Write(batch, nil)
	return err
}

func (s *ServerImpl) FindAll() (map[string]string, error) {
	iter := s.ldb.NewIterator(nil, nil)
	data := make(map[string]string, 0)
	for iter.Next() {
		key := iter.Key()
		val := iter.Value()
		data[string(key)] = string(val)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, err
	}
	return data, nil
}

//	比如key是key1-key1000
//	start = key1 limit = key2 就是从key1-key199
//	start = key100 limit = key105 就是从key100-key104
func (s *ServerImpl) FindLimit(start, limit string) (map[string]string, error) {
	data := make(map[string]string, 0)
	iter := s.ldb.NewIterator(&util.Range{Start: []byte(start), Limit: []byte(limit)}, nil)
	for iter.Next() {
		key := iter.Key()
		val := iter.Value()
		data[string(key)] = string(val)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *ServerImpl) FindByPrefix(pre string) (map[string]string, error) { // 模糊查询
	files := make(map[string]string,0)
	iter := s.ldb.NewIterator(util.BytesPrefix([]byte(pre)), nil)
	for iter.Next() {
		key := iter.Key()
		val := iter.Value()
		files[string(key)] = string(val)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return files, err
	}
	return files,nil
}

func (s *ServerImpl) Insert(key string, val interface{}) error {
	var (
		err  error
		data []byte
	)
	if s.ldb == nil {
		return errors.New("copyfile is null or db is null")
	}
	if data, err = json.Marshal(val); err != nil {
		return err
	}
	if err = s.ldb.Put([]byte(key), data, nil); err != nil {
		return err
	}
	return nil
}

func (s *ServerImpl) FindByKey(key string) (data []byte, err error) {
	if data, err = s.ldb.Get([]byte(key), nil); err != nil {
		return data, err
	}
	return data, nil
}
