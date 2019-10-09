package singleflight

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	var g Group
	v, err := g.Do("key", func() (interface{}, error) {
		return "date", nil
	})

	if got, want := fmt.Sprintf("%v (%T)", v, v), "date (string)"; got != want {
		t.Errorf("Do = %v; want %v", got, want)
	}

	if err != nil {
		t.Errorf("Do error = %v", err)
	}
}
