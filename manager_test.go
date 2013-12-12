package beef

import "testing"

func TestAllFunctions(t *testing.T) {
	server := NewManager()

	bf_name := "test"

	if server.ExistsBF(bf_name) {
		t.Errorf("Test bloom filter should not exist")
	}

	server.CreateBF(bf_name, 1000)

	if !server.ExistsBF(bf_name) {
		t.Errorf("Test bloom filter should exist")
	}

	if server.CheckBF(bf_name, "test value") {
		t.Errorf("Value should not be set")
	}

	server.InsertBF(bf_name, "test value")

	if !server.CheckBF(bf_name, "test value") {
		t.Errorf("Value should now be set")
	}
}
