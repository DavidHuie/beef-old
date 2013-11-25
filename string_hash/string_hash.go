package string_hash

import "hash/fnv"

const PRIME = 674506111

// Hash takes in a string to hash and an integer
// that corresponds to the hash function to apply.
// It then returns a uint64 hash of the string.
func Hash(str string, hash_number uint) uint64 {
	if hash_number == 0 {
		hash := fnv.New64()
		hash.Write([]byte(str))
		return hash.Sum64()
	} else if hash_number == 1 {
		hash := fnv.New64a()
		hash.Write([]byte(str))
		return hash.Sum64()
	} else {
		hash := Hash(str, 0)
		for i := uint(1); i < hash_number; i++ {
			partial_hash := Hash(str, i)
			hash *= PRIME * partial_hash
		}
		return hash
	}
}
