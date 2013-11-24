package bit_array

import "testing"

var test_new_bit_array_data = []struct {
	array             *BitArray
	expected_size     uint64
	expected_data_len uint64
}{
	{New(0), 0, 1},
	{New(1), 1, 1},
	{New(63), 63, 1},
	{New(64), 64, 2},
	{New(129), 129, 3},
}

func TestNewBitArray(t *testing.T) {
	for _, example := range test_new_bit_array_data {
		if example.array.size != example.expected_size {
			t.Errorf("Size of should be %v, got %v",
				example.expected_size,
				example.array.size)
		}
		if len(example.array.data) != int(example.expected_data_len) {
			t.Errorf("Expected length %v, got %v",
				example.expected_data_len,
				example.array.data)
		}
		for _, value := range example.array.data {
			if value != 0 {
				t.Errorf("Data should be zeroed out")
			}
		}

	}
}
func TestSetGet(t *testing.T) {
	ba := New(70)
	if ba.Get(70) != uint64(0) {
		t.Errorf("Entry at position %v should not be filled", 70)
	}
	ba.Set(70)
	if ba.Get(70) != uint64(1) {
		t.Errorf("Entry at position %v should be filled", 70)
	}
}
