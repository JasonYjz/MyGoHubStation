package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"log"
	"time"
)

type Person struct {
	Name string
	Age int
	Addr string
}

type MyCache struct {
	cache *bigcache.BigCache
}

func main() {
	mc := MyCache{cache: getCache()}

	p := Person{
		Name: "Jason",
		Age:  36,
		Addr: "Chengdu",
	}

	mc.Set("tst", p)

	v, err := mc.Get("tst")
	if err != nil {

	}

	person := v.(Person)
	fmt.Printf("Name:%s, Age:%d, Addr:%s\n", person.Name, person.Age, person.Addr)

	mc.cache.Delete("tst")
}

func getCache() *bigcache.BigCache {
	config := bigcache.Config {
		// number of shards (must be a power of 2)
		Shards: 1024,

		// time after which entry can be evicted
		LifeWindow: 10 * time.Minute,

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		CleanWindow: 5 * time.Minute,

		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,

		// prints information about additional memory allocation
		Verbose: true,

		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 8192,

		// callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,

		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		// Ignored if OnRemove is specified.
		OnRemoveWithReason: nil,
	}

	cache, initErr := bigcache.NewBigCache(config)
	if initErr != nil {
		log.Fatal(initErr)
	}

	return cache
}

func serializeGOB(value interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	gob.Register(value)

	err := enc.Encode(&value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func deserializeGOB(valueBytes []byte) (interface{}, error) {
	var value interface{}
	buf := bytes.NewBuffer(valueBytes)
	dec := gob.NewDecoder(buf)

	err := dec.Decode(&value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (c *MyCache) Set(key, value interface{}) error {
	keyString, ok := key.(string)
	if !ok {
		return errors.New("a cache key must be a string")
	}

	valueBytes, err := serializeGOB(value)
	if err != nil {
		return err
	}

	return c.cache.Set(keyString, valueBytes)
}

func (c *MyCache) Get(key interface{}) (interface{}, error) {
	// Assert the key is of string type
	keyString, ok := key.(string)
	if !ok {
		return nil, errors.New("a cache key must be a string")
	}

	// Get the value in the byte format it is stored in
	valueBytes, err := c.cache.Get(keyString)
	if err != nil {
		return nil, err
	}

	// Deserialize the bytes of the value
	value, err := deserializeGOB(valueBytes)
	if err != nil {
		return nil, err
	}

	return value, nil
}