package shared

import (
	"reflect"
	"testing"
)

// Expect is a simple testing matcher
func Expect(t *testing.T, got, want interface{}) {
	if !isEqual(got, want) {
		t.Errorf(`got %d, want %d`, got, want)
	}
}

func isEqual(a, b interface{}) bool {
	typeA := reflect.ValueOf(a)
	typeB := reflect.ValueOf(b)

	if typeA.Kind() != typeB.Kind() {
		panic("a, b should be same type")
	}

	switch typeA.Kind() {
	case reflect.Int:
		return a == b

	case reflect.String:
		return a == b

	default:
		return reflect.DeepEqual(a, b)
	}
}
