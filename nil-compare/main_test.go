package main

import (
	"reflect"
	"testing"
)

// Safely comparing values
func safeCompare(i1, i2 interface{}) bool {
	v1 := reflect.ValueOf(i1)
	v2 := reflect.ValueOf(i1)
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

	t.Run("system-compare", func(t *testing.T) {
		if typ != nil {
			t.Fatalf("Pointer value is always compatible to nil: [%v]", typ)
		}

		// Gotcha!
		if itf == nil {
			t.Fatalf("Interface value cannot be safely compared to \"naked\" nil: [%v]", itf)
		}
	})

	t.Run("safe-compare", func(t *testing.T) {
		// Converting to interface{} allows for comparison
		if !safeCompare(itf, nil) {
			t.Fatalf("Converted interface value and nil must be compatible: [%[1]v, %[1]T] vs. [%[2]v, %[2]T]",
				itf, nil)
		}
	})
}
