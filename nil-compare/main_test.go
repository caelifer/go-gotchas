package main

import "testing"

// Safely comparing values
func safeCompare(v1, v2 interface{}) bool {
	return v1 == v2
}

type I interface {
	F()
}
type T struct{}

func (*T) F() {}

func TestCompareNils(t *testing.T) {
	var typ *T = nil // Explicit initialization for better readability
	var itf I = typ  // Gotcha! itf == *tuple{nil, *T}, not nil

	t.Run("system", func(t *testing.T) {
		if typ != nil {
			t.Fatalf("Pointer value is always compatible to nil: [%v]", typ)
		}

		// Gotcha!
		if itf == nil {
			t.Fatalf("Interface value cannot be safely compared to \"naked\" nil: [%v]", itf)
		}
	})

	t.Run("empty-interface-casting", func(t *testing.T) {
		// Converting to interface{} allows for comparison
		if !safeCompare(itf, typ) {
			t.Fatalf("Converted value and interface nils must be compatible: [%v] vs. [%v]",
				itf, typ)
		}
	})
}
