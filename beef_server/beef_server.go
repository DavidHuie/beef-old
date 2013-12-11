package main

import (
	"fmt"
	"github.com/DavidHuie/beef/bloom_filter"
)

const (
	// The size of an underlying bit array
	// is this number times the expected size.
	bit_array_factor = 12
	// The number of hash functions use in
	// new bloom filters.
	num_hashes = 3
)

type beef_server struct {
	bit_arrays map[string]*bloom_filter.BloomFilter
}

func New() *beef_server {
	bs := new(beef_server)
	bs.bit_arrays = make(map[string]*bloom_filter.BloomFilter)
	return bs
}

func (b *beef_server) CreateBF(name string, size uint64) {
	b.bit_arrays[name] = bloom_filter.New(size*bit_array_factor, num_hashes)
}

func (b *beef_server) DeleteBF(name string) {
	delete(b.bit_arrays, name)
}

func (b *beef_server) InsertBF(name string, value string) {
	b.bit_arrays[name].Insert(value)
}

func (b *beef_server) CheckBF(name string, value string) bool {
	return b.bit_arrays[name].Check(value)
}

func main() {
	server := New()
	server.CreateBF("hallo", 40000000000)
	server.CreateBF("vassap", 40000000000)
	fmt.Println(server.CheckBF("hallo", "a test"))
	server.InsertBF("hallo", "a test")
	fmt.Println(server.CheckBF("hallo", "a test"))
	fmt.Println(server)
}
