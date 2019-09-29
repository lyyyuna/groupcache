package consistenhash

import (
	"testing"
)

func TestConsistency(t *testing.T) {
	hash1 := New(1)
	hash2 := New(1)
	hash1.Add("Bill", "Bob", "Bonny")
	hash2.Add("Bob", "Bonny", "Bill")

	if hash1.Get("Ben") != hash2.Get("Ben") {
		t.Error("Fetching 'ben' from both hashes should be same")
	}

	hash2.Add("Becky", "Ben", "Bobby")

	if hash2.Get("Ben") != hash1.Get("Ben") ||
		hash1.Get("Bob") != hash2.Get("Bob") ||
		hash1.Get("Bonny") != hash2.Get("Bonny") {
		t.Error("Direct match should always return the same entry")
	}
}
