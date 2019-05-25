package option

import "strings"
import "testing"

func TestNone(t *testing.T) {
	x := None{}
	if !x.IsNone() {
		t.Fail()
	}
	if x.IsSome() {
		t.Fail()
	}
	if x.Value() != "NONE" {
		t.Fail()
	}
}

func TestSome(t *testing.T) {
	x := Some{"foo"}
	if x.IsNone() {
		t.Fail()
	}
	if !x.IsSome() {
		t.Fail()
	}
	if x.Value() != "foo" {
		t.Fail()
	}
}

func TestOption(t *testing.T) {
	t.Run("map on None returns None", func(t *testing.T) {
		var x Option = None{}
		if !x.IsNone() {
			t.Fail()
		}
		y := x.Map(func(value string) string {
			return "mapped value"
		})
		if !y.IsNone() {
			t.Fail()
		}
	})
	t.Run("map on Some returns Some of mapped value", func(t *testing.T) {
		var x Option = Some{"foo"}
		y := x.
			Map(strings.ToUpper).
			Map(func(s string) string {
				return "Your string is: " + s
			})
		if y.Value() != "Your string is: FOO" {
			t.Fail()
		}
	})
}
