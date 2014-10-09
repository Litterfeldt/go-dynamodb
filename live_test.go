package dynamodb

import (
	"testing"
)

var test_key string = "test"

func test_table() *Table {
	return &Table{"test_table_123", "test_id", ""}
}

func init() {
	config := GetConfigFromEnv()
	ConfigureDbFromConfig(&config)
	test_table().Del(test_key)
}

func TestAddRow(t *testing.T) {
	attrs := make(map[string]AttrUpdate)
	attrs["test1"] = AttrUpdate{"test", ACTION_PUT}
	attrs["test2"] = AttrUpdate{"test", ACTION_PUT}
	resp, err := test_table().Add(attrs, test_key)

	if err != nil {
		t.Fatalf("An error occured when Uppdating test row\nError: %v", err)
	} else if resp.HashKey != "test" {
		t.Fatalf("Response Hashkey: %v did not match: %v", resp.HashKey, "test")
	} else if resp.Attributes["test1"] != "test" {
		t.Fatalf("Response attribute %v: %v did not match: %v", "test1", resp.HashKey, "test")
	} else if resp.Attributes["test2"] != "test" {
		t.Fatalf("Response attribute %v: %v did not match: %v", "test2", resp.HashKey, "test")
	}
}

func TestGetRow(t *testing.T) {
	resp, err := test_table().Get("test")

	if err != nil {
		t.Fatalf("An error occured when Getting test row\nError: %v", err)
	} else if resp.HashKey != "test" {
		t.Fatalf("Response Hashkey: %v did not match: %v", resp.HashKey, "test")
	} else if resp.Attributes["test1"] != "test" {
		t.Fatalf("Response attribute %v: %v did not match: %v", "test1", resp.HashKey, "test")
	} else if resp.Attributes["test2"] != "test" {
		t.Fatalf("Response attribute %v: %v did not match: %v", "test2", resp.HashKey, "test")
	}
}

func TestUpdateRow(t *testing.T) {
	attrs := make(map[string]AttrUpdate)
	attrs["test2"] = AttrUpdate{"test2", ACTION_PUT}
	resp, err := test_table().Add(attrs, test_key)

	if err != nil {
		t.Fatalf("An error occured when Uppdating test row\nError: %v", err)
	} else if resp.HashKey != "test" {
		t.Fatalf("Response Hashkey: %v did not match: %v", resp.HashKey, "test")
	} else if resp.Attributes["test1"] != "test" {
		t.Fatalf("Response attribute %v: %v did not match: %v", "test1", resp.HashKey, "test")
	} else if resp.Attributes["test2"] != "test2" {
		t.Fatalf("Response attribute %v: %v did not match: %v", "test2", resp.HashKey, "test2")
	}
}

func TestDeleteRow(t *testing.T) {
	attrs := make(map[string]AttrUpdate)
	attrs["test2"] = AttrUpdate{"test2", ACTION_PUT}
	err := test_table().Del(test_key)
	if err != nil {
		t.Fatalf("An error occured when deleting test row\nError: %v", err)
	}
	_, err = test_table().Get("test")

	if err == nil {
		t.Fatalf("Test row not deleted, still there")
	}
}
