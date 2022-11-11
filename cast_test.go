// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package cast

import (
	"encoding/json"
	"errors"
	"html/template"
	"reflect"
	"testing"
	"time"

	. "github.com/frankban/quicktest"
	"github.com/tidwall/gjson"
)

type testStep struct {
	input  any
	expect any
	iserr  bool
}

func createNumberTestSteps(zero, one, eight, eightnegative, eightpoint31, eightpoint31negative any) []testStep {
	var jeight, jminuseight, jfloateight json.Number
	_ = json.Unmarshal([]byte("8"), &jeight)
	_ = json.Unmarshal([]byte("-8"), &jminuseight)
	_ = json.Unmarshal([]byte("8.0"), &jfloateight)

	kind := reflect.TypeOf(zero).Kind()
	isUint := kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64

	// Some precision is lost when converting from float64 to float32.
	eightpoint31_32 := eightpoint31
	eightpoint31negative_32 := eightpoint31negative
	if kind == reflect.Float64 {
		eightpoint31_32 = float64(float32(eightpoint31.(float64)))
		eightpoint31negative_32 = float64(float32(eightpoint31negative.(float64)))
	}

	return []testStep{
		{int(8), eight, false},
		{int8(8), eight, false},
		{int16(8), eight, false},
		{int32(8), eight, false},
		{int64(8), eight, false},
		{time.Weekday(8), eight, false},
		{time.Month(8), eight, false},
		{uint(8), eight, false},
		{uint8(8), eight, false},
		{uint16(8), eight, false},
		{uint32(8), eight, false},
		{uint64(8), eight, false},
		{float32(8.31), eightpoint31_32, false},
		{float64(8.31), eightpoint31, false},
		{true, one, false},
		{false, zero, false},
		{"8", eight, false},
		{nil, zero, false},
		{int(-8), eightnegative, isUint},
		{int8(-8), eightnegative, isUint},
		{int16(-8), eightnegative, isUint},
		{int32(-8), eightnegative, isUint},
		{int64(-8), eightnegative, isUint},
		{float32(-8.31), eightpoint31negative_32, isUint},
		{float64(-8.31), eightpoint31negative, isUint},
		{"-8", eightnegative, isUint},
		{jeight, eight, false},
		{jminuseight, eightnegative, isUint},
		{jfloateight, eight, false},
		{"test", zero, true},
		{testing.T{}, zero, true},
	}
}

// Maybe Go 1.18 generics will make this less ugly?
func runNumberTest(c *C, tests []testStep, tove func(any) (any, error), tov func(any) any) {
	c.Helper()

	for i, test := range tests {
		errmsg := Commentf("i = %d", i)

		v, err := tove(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil, errmsg)
			continue
		}
		c.Assert(err, IsNil, errmsg)
		c.Assert(v, Equals, test.expect, errmsg)

		// Non-E test:
		v = tov(test.input)
		c.Assert(v, Equals, test.expect, errmsg)
	}
}

func TestToUintE(t *testing.T) {
	tests := createNumberTestSteps(uint(0), uint(1), uint(8), uint(0), uint(8), uint(8))

	runNumberTest(
		New(t),
		tests,
		func(v any) (any, error) { return ToUintE(v) },
		func(v any) any { return ToUint(v) },
	)
}

func TestToUint64E(t *testing.T) {
	tests := createNumberTestSteps(uint64(0), uint64(1), uint64(8), uint64(0), uint64(8), uint64(8))

	runNumberTest(
		New(t),
		tests,
		func(v any) (any, error) { return ToUint64E(v) },
		func(v any) any { return ToUint64(v) },
	)
}

func TestToUint32E(t *testing.T) {
	tests := createNumberTestSteps(uint32(0), uint32(1), uint32(8), uint32(0), uint32(8), uint32(8))

	runNumberTest(
		New(t),
		tests,
		func(v any) (any, error) { return ToUint32E(v) },
		func(v any) any { return ToUint32(v) },
	)
}

func TestToUint16E(t *testing.T) {
	tests := createNumberTestSteps(uint16(0), uint16(1), uint16(8), uint16(0), uint16(8), uint16(8))

	runNumberTest(
		New(t),
		tests,
		func(v any) (any, error) { return ToUint16E(v) },
		func(v any) any { return ToUint16(v) },
	)
}

func TestToUint8E(t *testing.T) {
	tests := createNumberTestSteps(uint8(0), uint8(1), uint8(8), uint8(0), uint8(8), uint8(8))

	runNumberTest(
		New(t),
		tests,
		func(v any) (any, error) { return ToUint8E(v) },
		func(v any) any { return ToUint8(v) },
	)
}
func TestToIntE(t *testing.T) {
	tests := createNumberTestSteps(int(0), int(1), int(8), int(-8), int(8), int(-8))

	runNumberTest(
		New(t),
		tests,
		func(v any) (any, error) { return ToIntE(v) },
		func(v any) any { return ToInt(v) },
	)
}

func TestToInt64E(t *testing.T) {
	tests := createNumberTestSteps(int64(0), int64(1), int64(8), int64(-8), int64(8), int64(-8))

	runNumberTest(
		New(t),
		tests,
		func(v any) (any, error) { return ToInt64E(v) },
		func(v any) any { return ToInt64(v) },
	)
}

func TestToInt32E(t *testing.T) {
	tests := createNumberTestSteps(int32(0), int32(1), int32(8), int32(-8), int32(8), int32(-8))

	runNumberTest(
		New(t),
		tests,
		func(v any) (any, error) { return ToInt32E(v) },
		func(v any) any { return ToInt32(v) },
	)
}

func TestToInt16E(t *testing.T) {
	tests := createNumberTestSteps(int16(0), int16(1), int16(8), int16(-8), int16(8), int16(-8))

	runNumberTest(
		New(t),
		tests,
		func(v any) (any, error) { return ToInt16E(v) },
		func(v any) any { return ToInt16(v) },
	)
}

func TestToInt8E(t *testing.T) {
	tests := createNumberTestSteps(int8(0), int8(1), int8(8), int8(-8), int8(8), int8(-8))

	runNumberTest(
		New(t),
		tests,
		func(v any) (any, error) { return ToInt8E(v) },
		func(v any) any { return ToInt8(v) },
	)
}

func TestToFloat64E(t *testing.T) {
	tests := createNumberTestSteps(float64(0), float64(1), float64(8), float64(-8), float64(8.31), float64(-8.31))

	runNumberTest(
		New(t),
		tests,
		func(v any) (any, error) { return ToFloat64E(v) },
		func(v any) any { return ToFloat64(v) },
	)
}

func TestToFloat32E(t *testing.T) {
	tests := createNumberTestSteps(float32(0), float32(1), float32(8), float32(-8), float32(8.31), float32(-8.31))

	runNumberTest(
		New(t),
		tests,
		func(v any) (any, error) { return ToFloat32E(v) },
		func(v any) any { return ToFloat32(v) },
	)
}

func TestToStringE(t *testing.T) {
	c := New(t)

	var jn json.Number
	_ = json.Unmarshal([]byte("8"), &jn)
	type Key struct {
		k string
	}
	key := &Key{"foo"}

	tests := []struct {
		input  any
		expect string
		iserr  bool
	}{
		{int(8), "8", false},
		{int8(8), "8", false},
		{int16(8), "8", false},
		{int32(8), "8", false},
		{int64(8), "8", false},
		{uint(8), "8", false},
		{uint8(8), "8", false},
		{uint16(8), "8", false},
		{uint32(8), "8", false},
		{uint64(8), "8", false},
		{float32(8.31), "8.31", false},
		{float64(8.31), "8.31", false},
		{jn, "8", false},
		{true, "true", false},
		{false, "false", false},
		{nil, "", false},
		{[]byte("one time"), "one time", false},
		{"one more time", "one more time", false},
		{template.HTML("one time"), "one time", false},
		{template.URL("http://somehost.foo"), "http://somehost.foo", false},
		{template.JS("(1+2)"), "(1+2)", false},
		{template.CSS("a"), "a", false},
		{template.HTMLAttr("a"), "a", false},
		// errors
		{testing.T{}, "", true},
		{key, "", true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToStringE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil, errmsg)
			continue
		}

		c.Assert(err, IsNil, errmsg)
		c.Assert(v, Equals, test.expect, errmsg)

		// Non-E test
		v = ToString(test.input)
		c.Assert(v, Equals, test.expect, errmsg)
	}
}

type foo struct {
	val string
}

func (x foo) String() string {
	return x.val
}

func TestStringerToString(t *testing.T) {
	c := New(t)

	var x foo
	x.val = "bar"
	c.Assert(ToString(x), Equals, "bar")
}

type fu struct {
	val string
}

func (x fu) Error() string {
	return x.val
}

func TestErrorToString(t *testing.T) {
	c := New(t)

	var x fu
	x.val = "bar"
	c.Assert(ToString(x), Equals, "bar")
}

func TestStringMapStringSliceE(t *testing.T) {
	c := New(t)

	// ToStringMapString inputs/outputs
	var stringMapString = map[string]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var stringMapInterface = map[string]any{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapString = map[any]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapInterface = map[any]any{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}

	// ToStringMapStringSlice inputs/outputs
	var stringMapStringSlice = map[string][]string{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}
	var stringMapInterfaceSlice = map[string][]any{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}
	var stringMapInterfaceInterfaceSlice = map[string]any{"key 1": []any{"value 1", "value 2", "value 3"}, "key 2": []any{"value 1", "value 2", "value 3"}, "key 3": []any{"value 1", "value 2", "value 3"}}
	var stringMapStringSingleSliceFieldsResult = map[string][]string{"key 1": {"value", "1"}, "key 2": {"value", "2"}, "key 3": {"value", "3"}}
	var interfaceMapStringSlice = map[any][]string{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}
	var interfaceMapInterfaceSlice = map[any][]any{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}

	var stringMapStringSliceMultiple = map[string][]string{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}
	var stringMapStringSliceSingle = map[string][]string{"key 1": {"value 1"}, "key 2": {"value 2"}, "key 3": {"value 3"}}

	var stringMapInterface1 = map[string]any{"key 1": []string{"value 1"}, "key 2": []string{"value 2"}}
	var stringMapInterfaceResult1 = map[string][]string{"key 1": {"value 1"}, "key 2": {"value 2"}}

	var jsonStringMapString = `{"key 1": "value 1", "key 2": "value 2"}`
	var jsonStringMapStringArray = `{"key 1": ["value 1"], "key 2": ["value 2", "value 3"]}`
	var jsonStringMapStringArrayResult = map[string][]string{"key 1": {"value 1"}, "key 2": {"value 2", "value 3"}}

	type Key struct {
		k string
	}

	tests := []struct {
		input  any
		expect map[string][]string
		iserr  bool
	}{
		{stringMapStringSlice, stringMapStringSlice, false},
		{stringMapInterfaceSlice, stringMapStringSlice, false},
		{stringMapInterfaceInterfaceSlice, stringMapStringSlice, false},
		{stringMapStringSliceMultiple, stringMapStringSlice, false},
		{stringMapStringSliceMultiple, stringMapStringSlice, false},
		{stringMapString, stringMapStringSliceSingle, false},
		{stringMapInterface, stringMapStringSliceSingle, false},
		{stringMapInterface1, stringMapInterfaceResult1, false},
		{interfaceMapStringSlice, stringMapStringSlice, false},
		{interfaceMapInterfaceSlice, stringMapStringSlice, false},
		{interfaceMapString, stringMapStringSingleSliceFieldsResult, false},
		{interfaceMapInterface, stringMapStringSingleSliceFieldsResult, false},
		{jsonStringMapStringArray, jsonStringMapStringArrayResult, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{map[any]any{"foo": testing.T{}}, nil, true},
		{map[any]any{Key{"foo"}: "bar"}, nil, true}, // ToStringE(Key{"foo"}) should fail
		{jsonStringMapString, nil, true},
		{"", nil, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToStringMapStringSliceE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil, errmsg)
			continue
		}

		c.Assert(err, IsNil, errmsg)
		c.Assert(v, DeepEquals, test.expect, errmsg)

		// Non-E test
		v = ToStringMapStringSlice(test.input)
		c.Assert(v, DeepEquals, test.expect, errmsg)
	}
}

func TestToStringMapE(t *testing.T) {
	c := New(t)

	tests := []struct {
		input  any
		expect map[string]any
		iserr  bool
	}{
		{map[any]any{"tag": "tags", "group": "groups"}, map[string]any{"tag": "tags", "group": "groups"}, false},
		{map[string]any{"tag": "tags", "group": "groups"}, map[string]any{"tag": "tags", "group": "groups"}, false},
		{`{"tag": "tags", "group": "groups"}`, map[string]any{"tag": "tags", "group": "groups"}, false},
		{`{"tag": "tags", "group": true}`, map[string]any{"tag": "tags", "group": true}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{"", nil, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToStringMapE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil, errmsg)
			continue
		}

		c.Assert(err, IsNil, errmsg)
		c.Assert(v, DeepEquals, test.expect, errmsg)

		// Non-E test
		v = ToStringMap(test.input)
		c.Assert(v, DeepEquals, test.expect, errmsg)
	}
}

func TestToStringMapBoolE(t *testing.T) {
	c := New(t)

	tests := []struct {
		input  any
		expect map[string]bool
		iserr  bool
	}{
		{map[any]any{"v1": true, "v2": false}, map[string]bool{"v1": true, "v2": false}, false},
		{map[string]any{"v1": true, "v2": false}, map[string]bool{"v1": true, "v2": false}, false},
		{map[string]bool{"v1": true, "v2": false}, map[string]bool{"v1": true, "v2": false}, false},
		{`{"v1": true, "v2": false}`, map[string]bool{"v1": true, "v2": false}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{"", nil, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToStringMapBoolE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil, errmsg)
			continue
		}

		c.Assert(err, IsNil, errmsg)
		c.Assert(v, DeepEquals, test.expect, errmsg)

		// Non-E test
		v = ToStringMapBool(test.input)
		c.Assert(v, DeepEquals, test.expect, errmsg)
	}
}

func TestToStringMapIntE(t *testing.T) {
	c := New(t)

	tests := []struct {
		input  any
		expect map[string]int
		iserr  bool
	}{
		{map[any]any{"v1": 1, "v2": 222}, map[string]int{"v1": 1, "v2": 222}, false},
		{map[string]any{"v1": 342, "v2": 5141}, map[string]int{"v1": 342, "v2": 5141}, false},
		{map[string]int{"v1": 33, "v2": 88}, map[string]int{"v1": 33, "v2": 88}, false},
		{map[string]int32{"v1": int32(33), "v2": int32(88)}, map[string]int{"v1": 33, "v2": 88}, false},
		{map[string]uint16{"v1": uint16(33), "v2": uint16(88)}, map[string]int{"v1": 33, "v2": 88}, false},
		{map[string]float64{"v1": float64(8.22), "v2": float64(43.32)}, map[string]int{"v1": 8, "v2": 43}, false},
		{`{"v1": 67, "v2": 56}`, map[string]int{"v1": 67, "v2": 56}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{"", nil, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToStringMapIntE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil, errmsg)
			continue
		}

		c.Assert(err, IsNil, errmsg)
		c.Assert(v, DeepEquals, test.expect, errmsg)

		// Non-E test
		v = ToStringMapInt(test.input)
		c.Assert(v, DeepEquals, test.expect, errmsg)
	}
}

func TestToStringMapInt64E(t *testing.T) {
	c := New(t)

	tests := []struct {
		input  any
		expect map[string]int64
		iserr  bool
	}{
		{map[any]any{"v1": int32(8), "v2": int32(888)}, map[string]int64{"v1": int64(8), "v2": int64(888)}, false},
		{map[string]any{"v1": int64(45), "v2": int64(67)}, map[string]int64{"v1": 45, "v2": 67}, false},
		{map[string]int64{"v1": 33, "v2": 88}, map[string]int64{"v1": 33, "v2": 88}, false},
		{map[string]int{"v1": 33, "v2": 88}, map[string]int64{"v1": 33, "v2": 88}, false},
		{map[string]int32{"v1": int32(33), "v2": int32(88)}, map[string]int64{"v1": 33, "v2": 88}, false},
		{map[string]uint16{"v1": uint16(33), "v2": uint16(88)}, map[string]int64{"v1": 33, "v2": 88}, false},
		{map[string]float64{"v1": float64(8.22), "v2": float64(43.32)}, map[string]int64{"v1": 8, "v2": 43}, false},
		{`{"v1": 67, "v2": 56}`, map[string]int64{"v1": 67, "v2": 56}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{"", nil, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToStringMapInt64E(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil)
			continue
		}

		c.Assert(err, IsNil)
		c.Assert(v, DeepEquals, test.expect, errmsg)

		// Non-E test
		v = ToStringMapInt64(test.input)
		c.Assert(v, DeepEquals, test.expect, errmsg)
	}
}

func TestToStringMapStringE(t *testing.T) {
	c := New(t)

	var stringMapString = map[string]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var stringMapInterface = map[string]any{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapString = map[any]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapInterface = map[any]any{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var jsonString = `{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}`
	var invalidJsonString = `{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"`
	var emptyString = ""

	tests := []struct {
		input  any
		expect map[string]string
		iserr  bool
	}{
		{stringMapString, stringMapString, false},
		{stringMapInterface, stringMapString, false},
		{interfaceMapString, stringMapString, false},
		{interfaceMapInterface, stringMapString, false},
		{jsonString, stringMapString, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{invalidJsonString, nil, true},
		{emptyString, nil, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToStringMapStringE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil)
			continue
		}

		c.Assert(err, IsNil)
		c.Assert(v, DeepEquals, test.expect, errmsg)

		// Non-E test
		v = ToStringMapString(test.input)
		c.Assert(v, DeepEquals, test.expect, errmsg)
	}
}

func TestToBoolSliceE(t *testing.T) {
	c := New(t)

	tests := []struct {
		input  any
		expect []bool
		iserr  bool
	}{
		{[]bool{true, false, true}, []bool{true, false, true}, false},
		{[]any{true, false, true}, []bool{true, false, true}, false},
		{[]int{1, 0, 1}, []bool{true, false, true}, false},
		{[]string{"true", "false", "true"}, []bool{true, false, true}, false},
		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{[]string{"foo", "bar"}, nil, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToBoolSliceE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil)
			continue
		}

		c.Assert(err, IsNil)
		c.Assert(v, DeepEquals, test.expect, errmsg)

		// Non-E test
		v = ToBoolSlice(test.input)
		c.Assert(v, DeepEquals, test.expect, errmsg)
	}
}

func TestToIntSliceE(t *testing.T) {
	c := New(t)

	tests := []struct {
		input  any
		expect []int
		iserr  bool
	}{
		{[]int{1, 3}, []int{1, 3}, false},
		{[]any{1.2, 3.2}, []int{1, 3}, false},
		{[]string{"2", "3"}, []int{2, 3}, false},
		{[2]string{"2", "3"}, []int{2, 3}, false},
		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{[]string{"foo", "bar"}, nil, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToIntSliceE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil)
			continue
		}

		c.Assert(err, IsNil)
		c.Assert(v, DeepEquals, test.expect, errmsg)

		// Non-E test
		v = ToIntSlice(test.input)
		c.Assert(v, DeepEquals, test.expect, errmsg)
	}
}

func TestToSliceE(t *testing.T) {
	c := New(t)

	tests := []struct {
		input  any
		expect []any
		iserr  bool
	}{
		{[]any{1, 3}, []any{1, 3}, false},
		{[]map[string]any{{"k1": 1}, {"k2": 2}}, []any{map[string]any{"k1": 1}, map[string]any{"k2": 2}}, false},
		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToSliceE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil)
			continue
		}

		c.Assert(err, IsNil)
		c.Assert(v, DeepEquals, test.expect, errmsg)

		// Non-E test
		v = ToSlice(test.input)
		c.Assert(v, DeepEquals, test.expect, errmsg)
	}
}

func TestToStringSliceE(t *testing.T) {
	c := New(t)

	tests := []struct {
		input  any
		expect []string
		iserr  bool
	}{
		{[]int{1, 2}, []string{"1", "2"}, false},
		{[]int8{int8(1), int8(2)}, []string{"1", "2"}, false},
		{[]int32{int32(1), int32(2)}, []string{"1", "2"}, false},
		{[]int64{int64(1), int64(2)}, []string{"1", "2"}, false},
		{[]float32{float32(1.01), float32(2.01)}, []string{"1.01", "2.01"}, false},
		{[]float64{float64(1.01), float64(2.01)}, []string{"1.01", "2.01"}, false},
		{[]string{"a", "b"}, []string{"a", "b"}, false},
		{[]any{1, 3}, []string{"1", "3"}, false},
		{any(1), []string{"1"}, false},
		{[]error{errors.New("a"), errors.New("b")}, []string{"a", "b"}, false},
		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToStringSliceE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil)
			continue
		}

		c.Assert(err, IsNil)
		c.Assert(v, DeepEquals, test.expect, errmsg)

		// Non-E test
		v = ToStringSlice(test.input)
		c.Assert(v, DeepEquals, test.expect, errmsg)
	}
}

func TestToDurationSliceE(t *testing.T) {
	c := New(t)

	tests := []struct {
		input  any
		expect []time.Duration
		iserr  bool
	}{
		{[]string{"1s", "1m"}, []time.Duration{time.Second, time.Minute}, false},
		{[]int{1, 2}, []time.Duration{1, 2}, false},
		{[]any{1, 3}, []time.Duration{1, 3}, false},
		{[]time.Duration{1, 3}, []time.Duration{1, 3}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{[]string{"invalid"}, nil, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToDurationSliceE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil)
			continue
		}

		c.Assert(err, IsNil)
		c.Assert(v, DeepEquals, test.expect, errmsg)

		// Non-E test
		v = ToDurationSlice(test.input)
		c.Assert(v, DeepEquals, test.expect, errmsg)
	}
}

func TestToBoolE(t *testing.T) {
	c := New(t)

	var jf, jt, je json.Number
	_ = json.Unmarshal([]byte("0"), &jf)
	_ = json.Unmarshal([]byte("1"), &jt)
	_ = json.Unmarshal([]byte("1.0"), &je)
	tests := []struct {
		input  any
		expect bool
		iserr  bool
	}{
		{0, false, false},
		{jf, false, false},
		{nil, false, false},
		{"false", false, false},
		{"FALSE", false, false},
		{"False", false, false},
		{"f", false, false},
		{"F", false, false},
		{false, false, false},

		{"true", true, false},
		{"TRUE", true, false},
		{"True", true, false},
		{"t", true, false},
		{"T", true, false},
		{1, true, false},
		{jt, true, false},
		{je, true, false},
		{true, true, false},
		{-1, true, false},

		// errors
		{"test", false, true},
		{testing.T{}, false, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToBoolE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil)
			continue
		}

		c.Assert(err, IsNil)
		c.Assert(v, Equals, test.expect, errmsg)

		// Non-E test
		v = ToBool(test.input)
		c.Assert(v, Equals, test.expect, errmsg)
	}
}

func BenchmarkTooBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !ToBool(true) {
			b.Fatal("ToBool returned false")
		}
	}
}

func BenchmarkTooInt(b *testing.B) {
	convert := func(num52 any) {
		if v := ToInt(num52); v != 52 {
			b.Fatalf("ToInt returned wrong value, got %d, want %d", v, 32)
		}
	}
	for i := 0; i < b.N; i++ {
		convert("52")
		convert(52.0)
		convert(uint64(52))
	}
}

func BenchmarkTrimZeroDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strToInt("")
		strToInt("123")
		strToInt("120")
		strToInt("120.00")
	}
}

func TestIndirectPointers(t *testing.T) {
	c := New(t)

	x := 13
	y := &x
	z := &y

	c.Assert(ToInt(y), Equals, 13)
	c.Assert(ToInt(z), Equals, 13)

}

func TestToTime(t *testing.T) {
	c := New(t)

	var jntime, jnetime json.Number
	_ = json.Unmarshal([]byte("1234567890"), &jntime)
	_ = json.Unmarshal([]byte("123.4567890"), &jnetime)
	tests := []struct {
		input  any
		expect time.Time
		iserr  bool
	}{
		{"2009-11-10 23:00:00 +0000 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},   // Time.String()
		{"Tue Nov 10 23:00:00 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},        // ANSIC
		{"Tue Nov 10 23:00:00 UTC 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},    // UnixDate
		{"Tue Nov 10 23:00:00 +0000 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},  // RubyDate
		{"10 Nov 09 23:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},             // RFC822
		{"10 Nov 09 23:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},           // RFC822Z
		{"Tuesday, 10-Nov-09 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false}, // RFC850
		{"Tue, 10 Nov 2009 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},   // RFC1123
		{"Tue, 10 Nov 2009 23:00:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false}, // RFC1123Z
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},            // RFC3339
		{"2018-10-21T23:21:29+0200", time.Date(2018, 10, 21, 21, 21, 29, 0, time.UTC), false},      // RFC3339 without timezone hh:mm colon
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},            // RFC3339Nano
		{"11:00PM", time.Date(0, 1, 1, 23, 0, 0, 0, time.UTC), false},                              // Kitchen
		{"Nov 10 23:00:00", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},                    // Stamp
		{"Nov 10 23:00:00.000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},                // StampMilli
		{"Nov 10 23:00:00.000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},             // StampMicro
		{"Nov 10 23:00:00.000000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},          // StampNano
		{"2016-03-06 15:28:01-00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},        // RFC3339 without T
		{"2016-03-06 15:28:01-0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},         // RFC3339 without T or timezone hh:mm colon
		{"2016-03-06 15:28:01", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 -0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 -00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 +0900", time.Date(2016, 3, 6, 6, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 +09:00", time.Date(2016, 3, 6, 6, 28, 1, 0, time.UTC), false},
		{"2006-01-02", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{"02 Jan 2006", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{1472574600, time.Date(2016, 8, 30, 16, 30, 0, 0, time.UTC), false},
		{int(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC), false},
		{int64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{int32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{uint(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC), false},
		{uint64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{uint32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{jntime, time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		// errors
		{"2006", time.Time{}, true},
		{jnetime, time.Time{}, true},
		{testing.T{}, time.Time{}, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToTimeE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil)
			continue
		}

		c.Assert(err, IsNil)
		c.Assert(v.UTC(), Equals, test.expect, errmsg)

		// Non-E test
		v = ToTime(test.input)
		c.Assert(v.UTC(), Equals, test.expect, errmsg)
	}
}

func TestToDurationE(t *testing.T) {
	c := New(t)

	var td time.Duration = 5
	var jn json.Number
	_ = json.Unmarshal([]byte("5"), &jn)

	tests := []struct {
		input  any
		expect time.Duration
		iserr  bool
	}{
		{time.Duration(5), td, false},
		{int(5), td, false},
		{int64(5), td, false},
		{int32(5), td, false},
		{int16(5), td, false},
		{int8(5), td, false},
		{uint(5), td, false},
		{uint64(5), td, false},
		{uint32(5), td, false},
		{uint16(5), td, false},
		{uint8(5), td, false},
		{float64(5), td, false},
		{float32(5), td, false},
		{jn, td, false},
		{string("5"), td, false},
		{string("5ns"), td, false},
		{string("5us"), time.Microsecond * td, false},
		{string("5µs"), time.Microsecond * td, false},
		{string("5ms"), time.Millisecond * td, false},
		{string("5s"), time.Second * td, false},
		{string("5m"), time.Minute * td, false},
		{string("5h"), time.Hour * td, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for i, test := range tests {
		errmsg := Commentf("i = %d", i) // assert helper message

		v, err := ToDurationE(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil)
			continue
		}

		c.Assert(err, IsNil)
		c.Assert(v, Equals, test.expect, errmsg)

		// Non-E test
		v = ToDuration(test.input)
		c.Assert(v, Equals, test.expect, errmsg)
	}
}

func TestTrimZero(t *testing.T) {
	c := New(t)

	c.Assert(ToInt("10.0"), Equals, 10)
	c.Assert(ToInt("10.00"), Equals, 10)
	c.Assert(ToFloat64("10.010"), Equals, 10.01)
	c.Assert(ToInt("0.0000000000"), Equals, 0)
	c.Assert(ToFloat64("0.00000000001"), Equals, 0.00000000001)
}

func assertTimeEqual(t *testing.T, expected, actual time.Time) {
	t.Helper()
	// Compare the dates using a numeric zone as there are cases where
	// time.Parse will assign a dummy location.
	Assert(t, actual.Format(time.RFC1123Z), Equals, expected.Format(time.RFC1123Z))
}

func assertLocationEqual(t *testing.T, expected, actual *time.Location) {
	t.Helper()
	Assert(t, locationEqual(expected, actual), IsTrue)
}

func locationEqual(a, b *time.Location) bool {
	// A note about comparring time.Locations:
	//   - can't only compare pointers
	//   - can't compare loc.String() because locations with the same
	//     name can have different offsets
	//   - can't use reflect.DeepEqual because time.Location has internal
	//     caches

	if a == b {
		return true
	} else if a == nil || b == nil {
		return false
	}

	// Check if they're equal by parsing times with a format that doesn't
	// include a timezone, which will interpret it as being a local time in
	// the given zone, and comparing the resulting local times.
	tA, err := time.ParseInLocation("2006-01-02", "2016-01-01", a)
	if err != nil {
		return false
	}

	tB, err := time.ParseInLocation("2006-01-02", "2016-01-01", b)
	if err != nil {
		return false
	}

	return tA.Equal(tB)
}

func TestError(t *testing.T) {
	c := New(t)

	c.Assert(ToError(errors.New("foo")).Error(), Equals, "foo")
	c.Assert(ToError("bar").Error(), Equals, "bar")
}

func TestFlatten(t *testing.T) {
	c := New(t)

	c.Assert(flatten(map[string]any{
		"foo": "bar",
		"baz": map[string]any{
			"qux": []any{
				"quux",
				"quuz",
			},
		},
	}), DeepEquals, map[string]any{
		"foo":       "bar",
		"baz.qux.0": "quux",
		"baz.qux.1": "quuz",
	})

	data, err := json.Marshal(map[string]any{
		"foo": "bar",
		"baz": map[string]any{
			"qux": []any{
				"quux",
				"quuz",
			},
		},
	})
	c.Assert(err, IsNil)
	c.Assert(gjson.GetBytes(data, "foo").String(), Equals, "bar")
	c.Assert(gjson.GetBytes(data, "baz.qux.0").String(), Equals, "quux")
	c.Assert(gjson.GetBytes(data, "baz.qux.1").String(), Equals, "quuz")

	c.Assert(flatten([]any{
		"foo",
		"bar",
		map[string]any{
			"baz": "qux",
		},
	}), DeepEquals, map[string]any{
		"0":     "foo",
		"1":     "bar",
		"2.baz": "qux",
	})

	data, err = json.Marshal([]any{
		"foo",
		"bar",
		map[string]any{
			"baz": "qux",
		},
	})
	c.Assert(err, IsNil)
	c.Assert(gjson.GetBytes(data, "0").String(), Equals, "foo")
	c.Assert(gjson.GetBytes(data, "1").String(), Equals, "bar")
	c.Assert(gjson.GetBytes(data, "2.baz").String(), Equals, "qux")
}
