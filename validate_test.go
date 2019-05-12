package validate

import (
	"testing"
	"time"
)

func TestType(t *testing.T) {
	if nil == Validate(1) {
		t.Errorf("validate should validate int type")
	}

	if nil == Validate(map[string]string{
		"a": "a",
	}) {
		t.Errorf("validate should validate map type")
	}

	if nil == Validate([]string{
		"a", "b",
	}) {
		t.Errorf("validate should validate slice type")
	}

	v := 1
	if nil == Validate(&v) {
		t.Errorf("validate should validate &int type")
	}

	s := struct {
		field int
	}{
		field: 1,
	}
	if nil != Validate(s) {
		t.Errorf("validate does no validate struct type")
	}
	if nil != Validate(s) {
		t.Errorf("validate does no validate struct pointer type")
	}
}

func TestMinTagForDuration(t *testing.T) {
	if nil == Validate(struct {
		field time.Duration `min:"0s"`
	}{
		field: -time.Second,
	}) {
		t.Errorf("min tag does not validate for time.Duratuon")
	}

	if nil == Validate(struct {
		field time.Duration `min:"-1s"`
	}{
		field: -time.Minute,
	}) {
		t.Errorf("min tag does not validate for time.Duratuon")
	}

	if nil != Validate(struct {
		field time.Duration `min:"0s"`
	}{
		field: 0,
	}) {
		t.Errorf("min tag does not validate for time.Duratuon")
	}

	if nil != Validate(struct {
		field time.Duration `min:"-1s"`
	}{
		field: -time.Millisecond,
	}) {
		t.Errorf("min tag does not validate for time.Duratuon")
	}
}

func TestMaxTagForDuration(t *testing.T) {
	if nil == Validate(struct {
		field time.Duration `max:"0s"`
	}{
		field: time.Second,
	}) {
		t.Errorf("max tag does not validate for time.Duratuon")
	}

	if nil == Validate(struct {
		field time.Duration `max:"1s"`
	}{
		field: time.Minute,
	}) {
		t.Errorf("max tag does not validate for time.Duratuon")
	}

	if nil != Validate(struct {
		field time.Duration `max:"0s"`
	}{
		field: 0,
	}) {
		t.Errorf("max tag does not validate for time.Duratuon")
	}

	if nil != Validate(struct {
		field time.Duration `max:"1s"`
	}{
		field: time.Millisecond,
	}) {
		t.Errorf("max tag does not validate for time.Duratuon")
	}
}

func TestMinTagForInt(t *testing.T) {
	if nil == Validate(struct {
		field int `min:"0"`
	}{
		field: -1,
	}) {
		t.Errorf("min tag does not validate for int")
	}

	if nil == Validate(struct {
		field int8 `min:"0"`
	}{
		field: -1,
	}) {
		t.Errorf("min tag does not validate for int8")
	}

	if nil == Validate(struct {
		field int16 `min:"0"`
	}{
		field: -1,
	}) {
		t.Errorf("min tag does not validate for int16")
	}

	if nil == Validate(struct {
		field int32 `min:"0"`
	}{
		field: -1,
	}) {
		t.Errorf("min tag does not validate for int32")
	}

	if nil == Validate(struct {
		field int64 `min:"0"`
	}{
		field: -1,
	}) {
		t.Errorf("min tag does not validate for int64")
	}

	if nil != Validate(struct {
		field int `min:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("min tag does not validate for int")
	}

	if nil != Validate(struct {
		field int8 `min:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("min tag does not validate for int8")
	}

	if nil != Validate(struct {
		field int16 `min:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("min tag does not validate for int16")
	}

	if nil != Validate(struct {
		field int32 `min:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("min tag does not validate for int32")
	}

	if nil != Validate(struct {
		field int64 `min:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("min tag does not validate for int64")
	}
}

func TestMaxTagForInt(t *testing.T) {
	if nil == Validate(struct {
		field int `max:"0"`
	}{
		field: 1,
	}) {
		t.Errorf("max tag does not validate for int")
	}

	if nil == Validate(struct {
		field int8 `max:"0"`
	}{
		field: 1,
	}) {
		t.Errorf("max tag does not validate for int8")
	}

	if nil == Validate(struct {
		field int16 `max:"0"`
	}{
		field: 1,
	}) {
		t.Errorf("max tag does not validate for int16")
	}

	if nil == Validate(struct {
		field int32 `max:"0"`
	}{
		field: 1,
	}) {
		t.Errorf("max tag does not validate for int32")
	}

	if nil == Validate(struct {
		field int64 `max:"0"`
	}{
		field: 1,
	}) {
		t.Errorf("max tag does not validate for int64")
	}

	if nil != Validate(struct {
		field int `max:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("max tag does not validate for int")
	}

	if nil != Validate(struct {
		field int8 `max:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("max tag does not validate for int8")
	}

	if nil != Validate(struct {
		field int16 `max:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("max tag does not validate for int16")
	}

	if nil != Validate(struct {
		field int32 `max:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("max tag does not validate for int32")
	}

	if nil != Validate(struct {
		field int64 `max:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("max tag does not validate for int64")
	}
}

func TestMinTagForUint(t *testing.T) {
	if nil == Validate(struct {
		field uint `min:"10"`
	}{
		field: 9,
	}) {
		t.Errorf("min tag does not validate for uint")
	}

	if nil == Validate(struct {
		field uint8 `min:"10"`
	}{
		field: 9,
	}) {
		t.Errorf("min tag does not validate for uint8")
	}

	if nil == Validate(struct {
		field uint16 `min:"10"`
	}{
		field: 9,
	}) {
		t.Errorf("min tag does not validate for uint16")
	}

	if nil == Validate(struct {
		field uint32 `min:"10"`
	}{
		field: 9,
	}) {
		t.Errorf("min tag does not validate for uint32")
	}

	if nil == Validate(struct {
		field uint64 `min:"10"`
	}{
		field: 9,
	}) {
		t.Errorf("min tag does not validate for uint64")
	}

	if nil != Validate(struct {
		field uint `min:"10"`
	}{
		field: 10,
	}) {
		t.Errorf("min tag does not validate for uint")
	}

	if nil != Validate(struct {
		field uint8 `min:"10"`
	}{
		field: 10,
	}) {
		t.Errorf("min tag does not validate for uint8")
	}

	if nil != Validate(struct {
		field uint16 `min:"10"`
	}{
		field: 10,
	}) {
		t.Errorf("min tag does not validate for uint16")
	}

	if nil != Validate(struct {
		field uint32 `min:"10"`
	}{
		field: 10,
	}) {
		t.Errorf("min tag does not validate for uint32")
	}

	if nil != Validate(struct {
		field uint64 `min:"10"`
	}{
		field: 10,
	}) {
		t.Errorf("min tag does not validate for uint64")
	}
}

func TestMaxTagForUint(t *testing.T) {
	if nil == Validate(struct {
		field uint `max:"0"`
	}{
		field: 1,
	}) {
		t.Errorf("max tag does not validate for uint")
	}

	if nil == Validate(struct {
		field uint8 `max:"0"`
	}{
		field: 1,
	}) {
		t.Errorf("max tag does not validate for uint8")
	}

	if nil == Validate(struct {
		field uint16 `max:"0"`
	}{
		field: 1,
	}) {
		t.Errorf("max tag does not validate for uint16")
	}

	if nil == Validate(struct {
		field uint32 `max:"0"`
	}{
		field: 1,
	}) {
		t.Errorf("max tag does not validate for uint32")
	}

	if nil == Validate(struct {
		field uint64 `max:"0"`
	}{
		field: 1,
	}) {
		t.Errorf("max tag does not validate for uint64")
	}

	if nil != Validate(struct {
		field uint `max:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("max tag does not validate for uint")
	}

	if nil != Validate(struct {
		field uint8 `max:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("max tag does not validate for uint8")
	}

	if nil != Validate(struct {
		field uint16 `max:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("max tag does not validate for uint16")
	}

	if nil != Validate(struct {
		field uint32 `max:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("max tag does not validate for uint32")
	}

	if nil != Validate(struct {
		field uint64 `max:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("max tag does not validate for uint64")
	}
}

func TestMinTagForFloat(t *testing.T) {
	if nil == Validate(struct {
		field float32 `min:"0"`
	}{
		field: -0.1,
	}) {
		t.Errorf("min tag does not validate for flaot32")
	}

	if nil == Validate(struct {
		field float64 `min:"0"`
	}{
		field: -0.1,
	}) {
		t.Errorf("min tag does not validate for flaot64")
	}

	if nil != Validate(struct {
		field float32 `min:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("min tag does not validate for flaot32")
	}

	if nil != Validate(struct {
		field float64 `min:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("min tag does not validate for flaot64")
	}
}

func TestMaxTagForFloat(t *testing.T) {
	if nil == Validate(struct {
		field float32 `max:"0"`
	}{
		field: 0.1,
	}) {
		t.Errorf("max tag does not validate for flaot32")
	}

	if nil == Validate(struct {
		field float64 `max:"0"`
	}{
		field: 0.1,
	}) {
		t.Errorf("max tag does not validate for flaot64")
	}

	if nil != Validate(struct {
		field float32 `max:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("max tag does not validate for flaot32")
	}

	if nil != Validate(struct {
		field float64 `max:"0"`
	}{
		field: 0,
	}) {
		t.Errorf("max tag does not validate for flaot64")
	}
}

func TestMinTagForString(t *testing.T) {
	if nil == Validate(struct {
		field string `min:"2"`
	}{
		field: "a",
	}) {
		t.Errorf("min tag does not validate for string")
	}

	if nil != Validate(struct {
		field string `min:"2"`
	}{
		field: "ab",
	}) {
		t.Errorf("min tag does not validate for string")
	}
}

func TestMaxTagForString(t *testing.T) {
	if nil == Validate(struct {
		field string `max:"2"`
	}{
		field: "abc",
	}) {
		t.Errorf("max tag does not validate for string")
	}

	if nil != Validate(struct {
		field string `max:"2"`
	}{
		field: "ab",
	}) {
		t.Errorf("max tag does not validate for string")
	}
}

func TestMinTagForMap(t *testing.T) {
	if nil == Validate(struct {
		field map[string]string `min:"2"`
	}{
		field: map[string]string{
			"a": "a",
		},
	}) {
		t.Errorf("min tag does not validate for map")
	}

	if nil != Validate(struct {
		field map[string]string `min:"2"`
	}{
		field: map[string]string{
			"a": "a",
			"b": "b",
		},
	}) {
		t.Errorf("min tag does not validate for map")
	}
}

func TestMaxForMap(t *testing.T) {
	if nil == Validate(struct {
		field map[string]string `max:"2"`
	}{
		field: map[string]string{
			"a": "a",
			"b": "b",
			"c": "c",
		},
	}) {
		t.Errorf("max tag does not validate for map")
	}

	if nil != Validate(struct {
		field map[string]string `max:"2"`
	}{
		field: map[string]string{
			"a": "a",
			"b": "b",
		},
	}) {
		t.Errorf("max tag does not validate for map")
	}
}

func TestMinTagForSlice(t *testing.T) {
	if nil == Validate(struct {
		field []string `min:"2"`
	}{
		field: []string{"a"},
	}) {
		t.Errorf("min tag does not validate for string")
	}

	if nil != Validate(struct {
		field []string `min:"2"`
	}{
		field: []string{"a", "b"},
	}) {
		t.Errorf("min tag does not validate for string")
	}
}

func TestMaxTagForSlice(t *testing.T) {
	if nil == Validate(struct {
		field []string `max:"2"`
	}{
		field: []string{"a", "b", "c"},
	}) {
		t.Errorf("min tag does not validate for string")
	}

	if nil != Validate(struct {
		field []string `max:"2"`
	}{
		field: []string{"a", "b"},
	}) {
		t.Errorf("min tag does not validate for string")
	}
}

func TestIsEmptyTagForString(t *testing.T) {
	if nil == Validate(struct {
		field string `is_empty:"true"`
	}{
		field: "a",
	}) {
		t.Errorf("is_empty tag does not validate for string")
	}

	if nil != Validate(struct {
		field string `is_empty:"true"`
	}{
		field: "",
	}) {
		t.Errorf("is_empty tag does not validate for string")
	}

	if nil == Validate(struct {
		field string `is_empty:"false"`
	}{
		field: "",
	}) {
		t.Errorf("is_empty tag does not validate for string")
	}

	if nil != Validate(struct {
		field string `is_empty:"false"`
	}{
		field: "a",
	}) {
		t.Errorf("is_empty tag does not validate for string")
	}
}

func TestIsEmptyTagForMap(t *testing.T) {
	if nil == Validate(struct {
		field map[string]string `is_empty:"true"`
	}{
		field: map[string]string{"a": "a"},
	}) {
		t.Errorf("is_empty tag does not validate for map")
	}

	if nil != Validate(struct {
		field map[string]string `is_empty:"true"`
	}{
		field: map[string]string{},
	}) {
		t.Errorf("is_empty tag does not validate for map")
	}

	if nil == Validate(struct {
		field map[string]string `is_empty:"false"`
	}{
		field: map[string]string{},
	}) {
		t.Errorf("is_empty tag does not validate for map")
	}

	if nil != Validate(struct {
		field map[string]string `is_empty:"false"`
	}{
		field: map[string]string{"a": "a"},
	}) {
		t.Errorf("is_empty tag does not validate for map")
	}
}

func TestIsEmptyTagForSlice(t *testing.T) {
	if nil == Validate(struct {
		field []string `is_empty:"true"`
	}{
		field: []string{
			"a",
		},
	}) {
		t.Errorf("is_empty tag does not validate for slice")
	}

	if nil != Validate(struct {
		field []string `is_empty:"true"`
	}{
		field: []string{},
	}) {
		t.Errorf("is_empty tag does not validate for sclie")
	}

	if nil == Validate(struct {
		field []string `is_empty:"false"`
	}{
		field: []string{},
	}) {
		t.Errorf("is_empty tag does not validate for slice")
	}

	if nil != Validate(struct {
		field []string `is_empty:"false"`
	}{
		field: []string{"a"},
	}) {
		t.Errorf("is_empty tag does not validate for slice")
	}
}

func TestIsNilTagForPtr(t *testing.T) {
	if nil == Validate(struct {
		field *int `is_nil:"true"`
	}{
		field: new(int),
	}) {
		t.Errorf("is_nill tag does not validate for pointer")
	}

	if nil != Validate(struct {
		field *int `is_nil:"true"`
	}{
		field: nil,
	}) {
		t.Errorf("is_nill tag does not validate for pointer")
	}

	if nil == Validate(struct {
		field *int `is_nil:"false"`
	}{
		field: nil,
	}) {
		t.Errorf("is_nill tag does not validate for pointer")
	}

	if nil != Validate(struct {
		field *int `is_nil:"false"`
	}{
		field: new(int),
	}) {
		t.Errorf("is_nill tag does not validate for pointer")
	}
}

func TestChildTagsForSlice(t *testing.T) {
	if nil == Validate(struct {
		field []int `child_min:"0"`
	}{
		field: []int{0, -1},
	}) {
		t.Errorf("child_min tag does not validate for slice")
	}

	if nil != Validate(struct {
		field []int `child_min:"0"`
	}{
		field: []int{0, 0},
	}) {
		t.Errorf("child_min tag does not validate for slice")
	}

	if nil == Validate(struct {
		field []int `child_max:"0"`
	}{
		field: []int{0, 1},
	}) {
		t.Errorf("child_max tag does not validate for slice")
	}

	if nil != Validate(struct {
		field []int `child_max:"0"`
	}{
		field: []int{0, 0},
	}) {
		t.Errorf("child_max tag does not validate for slice")
	}

	if nil == Validate(struct {
		field [][]int `child_is_empty:"true"`
	}{
		field: [][]int{
			[]int{0},
		},
	}) {
		t.Errorf("child_is_empty tag does not validate for slice")
	}

	if nil != Validate(struct {
		field [][]int `child_is_empty:"true"`
	}{
		field: [][]int{
			[]int{},
		},
	}) {
		t.Errorf("child_is_empty tag does not validate for slice")
	}

	if nil == Validate(struct {
		field [][]int `child_is_empty:"false"`
	}{
		field: [][]int{
			[]int{},
		},
	}) {
		t.Errorf("child_is_empty tag does not validate for slice")
	}

	if nil != Validate(struct {
		field [][]int `child_is_empty:"false"`
	}{
		field: [][]int{
			[]int{0},
		},
	}) {
		t.Errorf("child_is_empty tag does not validate for slice")
	}

	if nil == Validate(struct {
		field []*int `child_is_nil:"true"`
	}{
		field: []*int{
			new(int),
		},
	}) {
		t.Errorf("child_is_nil tag does not validate for slice")
	}

	if nil != Validate(struct {
		field []*int `child_is_nil:"true"`
	}{
		field: []*int{nil},
	}) {
		t.Errorf("child_is_nil tag does not validate for slice")
	}

	if nil == Validate(struct {
		field []*int `child_is_nil:"false"`
	}{
		field: []*int{nil},
	}) {
		t.Errorf("child_is_nil tag does not validate for slice")
	}

	if nil != Validate(struct {
		field []*int `child_is_nil:"false"`
	}{
		field: []*int{new(int)},
	}) {
		t.Errorf("child_is_nil tag does not validate for slice")
	}
}

func TestChildTagsForPtr(t *testing.T) {
	minusOne := -1
	zero := 0
	one := 1
	empty := ""
	notEmpty := "a"
	onePtr := &one
	var nilPtr *int

	if nil == Validate(struct {
		field *int `child_min:"0"`
	}{
		field: &minusOne,
	}) {
		t.Errorf("child_min tag does not validate for pointer")
	}

	if nil != Validate(struct {
		field *int `child_min:"0"`
	}{
		field: &zero,
	}) {
		t.Errorf("child_min tag does not validate for pointer")
	}

	if nil == Validate(struct {
		field *int `child_max:"0"`
	}{
		field: &one,
	}) {
		t.Errorf("child_max tag does not validate for pointer")
	}

	if nil != Validate(struct {
		field *int `child_max:"0"`
	}{
		field: &zero,
	}) {
		t.Errorf("child_max tag does not validate for pointer")
	}

	if nil == Validate(struct {
		field *string `child_is_empty:"true"`
	}{
		field: &notEmpty,
	}) {
		t.Errorf("child_is_empty tag does not validate for pointer")
	}

	if nil != Validate(struct {
		field *string `child_is_empty:"true"`
	}{
		field: &empty,
	}) {
		t.Errorf("child_is_empty tag does not validate for pointer")
	}

	if nil == Validate(struct {
		field *string `child_is_empty:"false"`
	}{
		field: &empty,
	}) {
		t.Errorf("child_is_empty tag does not validate for pointer")
	}

	if nil != Validate(struct {
		field *string `child_is_empty:"false"`
	}{
		field: &notEmpty,
	}) {
		t.Errorf("child_is_empty tag does not validate for pointer")
	}

	if nil == Validate(struct {
		field **int `child_is_nil:"true"`
	}{
		field: &onePtr,
	}) {
		t.Errorf("child_is_nil tag does not validate for pointer")
	}

	if nil != Validate(struct {
		field **int `child_is_nil:"true"`
	}{
		field: &nilPtr,
	}) {
		t.Errorf("child_is_nil tag does not validate for pointer")
	}

	if nil == Validate(struct {
		field **int `child_is_nil:"false"`
	}{
		field: &nilPtr,
	}) {
		t.Errorf("child_is_nil tag does not validate for pointer")
	}

	if nil != Validate(struct {
		field **int `child_is_nil:"false"`
	}{
		field: &onePtr,
	}) {
		t.Errorf("child_is_nil tag does not validate for pointer")
	}
}