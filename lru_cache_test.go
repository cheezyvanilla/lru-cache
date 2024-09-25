package main

import "testing"

// create test for this flow -> put(1,1) then put(2,2) then get(1) then put(3,3) then get(2)
func TestLruCache(t *testing.T) {
	obj := NewLruCache(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	if obj.Get(1) != 1 {
		t.Errorf("Expected 1 but got %d", obj.Get(1))
	}

	obj.Put(3, 3)
	if obj.Get(2) != -1 {
		t.Errorf("Expected -1 but got %d", obj.Get(2))
	}

	obj.Put(4, 4)

	if obj.Get(1) != -1 {
		t.Errorf("Expected -1 but got %d", obj.Get(1))
	}
	if obj.Get(3) != 3 {
		t.Errorf("Expected 3 but got %d", obj.Get(3))
	}
	if obj.Get(4) != 4 {
		t.Errorf("Expected 4 but got %d", obj.Get(4))
	}

	// test 2
	obj = NewLruCache(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)
	if obj.Get(1) != -1 {
		t.Errorf("Expected -1 but got %d", obj.Get(1))
	}
	if obj.Get(2) != 3 {
		t.Errorf("Expected 3 but got %d", obj.Get(2))
	}
}
