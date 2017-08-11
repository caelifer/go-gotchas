package main

import "testing"

type I interface {
	F()
}
type T struct{}

func (*T) F() {}

func TestCompareNils(t *testing.T) {
	var typ *T // default initialization to nil
	var itf I = typ

	t.Run("system", func(t *testing.T) {
		if typ != nil {
			t.Fatalf("Pointer value is always compatible to nil: [%v]", typ)
		}

		if itf == nil {
			t.Fatalf("Interface value cannot be compared to system nil: [%v]", itf)
		}
	})

	t.Run("empty-interface-casting", func(t *testing.T) {
		v1 := interface{}(itf)
		v2 := interface{}(typ)
		if v1 != v2 {
			t.Fatalf("Converted value and interface nils must be compatible: [%v] vs. [%v]",
				v1, v2)
		}
	})
}
