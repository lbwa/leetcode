package shared

import (
	"reflect"
	"testing"
)

func Expect(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`got %d, want %d`, got, want)
	}
}
