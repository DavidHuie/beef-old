package beef

import "testing"

func TestAllFunctions(t *testing.T) {
	server := NewManager()

	bf_name := "test"

	if server.ExistsBF(bf_name) {
		t.Errorf("Test bloom filter should not exist")
	}

	error := server.CreateBF(bf_name, 1000)

	if error != nil {
		t.Errorf("Creating a bloom filter should not throw an error")
	}

	if !server.ExistsBF(bf_name) {
		t.Errorf("Test bloom filter should exist")
	}

	if value, _ := server.CheckBF(bf_name, "test value"); value {
		t.Errorf("Value should not be set")
	}

	if _, error := server.CheckBF(bf_name, "test value"); error != nil {
		t.Errorf("There should not be errors because the bloom filter has been created")
	}

	if _, error := server.CheckBF("xxx", "test value"); error == nil {
		t.Errorf("Checking for non-existing bloom filter should throw an error")
	}

	error = server.InsertBF(bf_name, "test value")

	if error != nil {
		t.Errorf("There should not be an error inserting into existing bloom filter")
	}

	if _, error := server.CheckBF(bf_name, "test value"); error != nil {
		t.Errorf("Checking for a set value should not throw an error")
	}

	if value, _ := server.CheckBF(bf_name, "test value"); !value {
		t.Errorf("Checking for a set value should not throw an error")
	}
}
