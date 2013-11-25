package bloom_filter

import "testing"

var test_hash_values = []uint64{1348, 9754, 3048, 4528}
var test_string = "test string"

func TestHashValues(t *testing.T) {
	bf := New(10000, 4)
	result := bf.hash_values(test_string)
	for i, value := range result {
		if test_hash_values[i] != value {
			t.Errorf("Expected %v, got %v",
				test_hash_values[i], value)
		}
	}
}

func TestInsertAndCheck(t *testing.T) {
	bf := New(1000, 3)
	if bf.Check(test_string) {
		t.Errorf("%v should not be initially set", test_string)
	}
	bf.Insert(test_string)
	if !bf.Check(test_string) {
		t.Errorf("%v should be set", test_string)
	}
}
