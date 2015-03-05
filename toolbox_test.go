package silkroad

import "testing"

func TestStringInSlice(t *testing.T) {
	slice := []string{"a", "b", "c"}

	if got, want := StringInSlice(slice, "a"), true; got != want {
		t.Errorf("TestStringInSlice got %v, want %v", got, want)
	}

	if got, want := StringInSlice(slice, "d"), false; got != want {
		t.Errorf("TestStringInSlice got %v, want %v", got, want)
	}

}