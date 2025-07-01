package runtime_test

import (
	"reflect"
	"slices"
	"testing"

	"github.com/tempo-lang/tempo/runtime"
)

func TestCopySliceSimple(t *testing.T) {
	array := []int{1, 2, 3, 4}

	copy := runtime.Copy(array)

	array[1] = 100

	if !slices.Equal(copy, []int{1, 2, 3, 4}) {
		t.Fatalf("wrong copy: %v", copy)
	}
}

func TestCopySliceNested(t *testing.T) {
	array := [][]int{{1, 2}, {3, 4}, {5, 6}}

	copy := runtime.Copy(array)

	array[1][1] = 100

	if !reflect.DeepEqual(copy, [][]int{{1, 2}, {3, 4}, {5, 6}}) {
		t.Fatalf("wrong copy: %v", copy)
	}
}

func TestCopyStructSlice(t *testing.T) {
	type StructSlice struct {
		Other string
		Inner []int
	}

	obj := []StructSlice{
		{Other: "Hello", Inner: []int{1, 2, 3}},
		{Other: "World", Inner: []int{4, 5, 6}},
	}

	copy := runtime.Copy(obj)

	obj[0].Inner[1] = 100
	obj[1].Inner[0] = 100
	obj[1].Other = "Changed"

	expected := []StructSlice{
		{Other: "Hello", Inner: []int{1, 2, 3}},
		{Other: "World", Inner: []int{4, 5, 6}},
	}

	if !reflect.DeepEqual(copy, expected) {
		t.Fatalf("wrong copy: %v", copy)
	}
}

func TestCopyAsync(t *testing.T) {

	obj := []int{1, 2, 3}
	asyncObj := runtime.FixedAsync(obj)

	asyncCopy := runtime.Copy(asyncObj)

	value1 := runtime.GetAsync(asyncObj)
	value2 := runtime.GetAsync(asyncCopy)

	value1[0] = 123
	value2[0] = 456

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(obj, expected) {
		t.Fatalf("wrong copy: %v", obj)
	}

	newCopy := runtime.GetAsync(asyncCopy)
	if !reflect.DeepEqual(newCopy, expected) {
		t.Fatalf("wrong copy: %v", newCopy)
	}
}
