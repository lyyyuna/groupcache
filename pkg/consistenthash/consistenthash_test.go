package consistenhash

import (
	"testing"
)

func TestConsistency(t *testing.T) {
	hash1 := New(1, nil)
	hash2 := New(1, nil)
	hash1.Add("Bill", "bob", "Bonny")
	hash2.Add("bob", "Bonny", "Bill")

	if hash1.Get("Ben") != hash2.Get("Ben") {
		t.Error("Fetching 'ben' from both hashes should be same")
	}

	hash2.Add("Becky", "Ben", "Bobby")

	if hash2.Get("Ben") != hash1.Get("Ben") ||
		hash1.Get("bob0") != hash2.Get("bob0") ||
		hash1.Get("Bonny") != hash2.Get("Bonny") {
		t.Error("Direct match should always return the same entry")
	}
}
