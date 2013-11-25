package string_hash

import "testing"

var test_str string = "this is a test string"
var test_values = []struct {
	hash_number    uint
	expected_value uint64
}{
	{0, 196278022052473809},
	{1, 6353911612108993833},
	{2, 3586045451050446855},
	{3, 11260069749024759857},
	{9, 15496851938882022401},
}

func TestHashFunctionsAreUnique(t *testing.T) {
	for _, value := range test_values {
		hash := Hash(test_str, value.hash_number)
		if hash != value.expected_value {
			t.Errorf("Expected hash value of %v, got %v",
				value.expected_value, hash)
		}
	}
}
