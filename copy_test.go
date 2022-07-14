// promconfig
// Copyright 2020 Percona LLC
//
// Based on Joel Scoble (github.com/mohae).
// Copyright 2014-2016.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package promconfig

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

// just basic is this working stuff
func TestSimple(t *testing.T) {
	Strings := []string{"a", "b", "c"}
	cpyS := Copy(Strings).([]string)
	if (*reflect.SliceHeader)(unsafe.Pointer(&Strings)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpyS)).Data {
		t.Error("[]string: expected SliceHeader data pointers to point to different locations, they didn't")
	}
	for i, v := range Strings {
		if v != cpyS[i] {
			t.Errorf("[]string: got %v at index %d of the copy; want %v", cpyS[i], i, v)
		}
	}

	Bytes := []byte("hello")
	cpyBt := Copy(Bytes).([]byte)
	if (*reflect.SliceHeader)(unsafe.Pointer(&Strings)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpyBt)).Data {
		t.Error("[]byte: expected SliceHeader data pointers to point to different locations, they didn't")
	}
	if len(cpyBt) != len(Bytes) {
		t.Errorf("[]byte: len was %d; want %d", len(cpyBt), len(Bytes))
	}
	for i, v := range Bytes {
		if v != cpyBt[i] {
			t.Errorf("[]byte: got %v at index %d of the copy; want %v", cpyBt[i], i, v)
		}
	}

	Interfaces := []interface{}{"a", 42, true, 4.32}
	cpyIf := Copy(Interfaces).([]interface{})
	if (*reflect.SliceHeader)(unsafe.Pointer(&Strings)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpyIf)).Data {
		t.Error("[]interfaces: expected SliceHeader data pointers to point to different locations, they didn't")
		return
	}
	if len(cpyIf) != len(Interfaces) {
		t.Errorf("[]interface{}: len was %d; want %d", len(cpyIf), len(Interfaces))
		return
	}
	for i, v := range Interfaces {
		if v != cpyIf[i] {
			t.Errorf("[]interface{}: got %v at index %d of the copy; want %v", cpyIf[i], i, v)
		}
	}
}

type Basics struct {
	String      string
	Strings     []string
	StringArr   [4]string
	Bool        bool
	Bools       []bool
	Byte        byte
	Bytes       []byte
	Int         int
	Ints        []int
	Int8        int8
	Int8s       []int8
	Int16       int16
	Int16s      []int16
	Int32       int32
	Int32s      []int32
	Int64       int64
	Int64s      []int64
	Uint        uint
	Uints       []uint
	Uint8       uint8
	Uint8s      []uint8
	Uint16      uint16
	Uint16s     []uint16
	Uint32      uint32
	Uint32s     []uint32
	Uint64      uint64
	Uint64s     []uint64
	Float32     float32
	Float32s    []float32
	Float64     float64
	Float64s    []float64
	Complex64   complex64
	Complex64s  []complex64
	Complex128  complex128
	Complex128s []complex128
	Interface   interface{}
	Interfaces  []interface{}
}

// These tests test that all supported basic types are copied correctly.  This
// is done by copying a struct with fields of most of the basic types as []T.
func TestMostTypes(t *testing.T) {
	test := Basics{
		String:      "kimchi",
		Strings:     []string{"uni", "ika"},
		StringArr:   [4]string{"malort", "barenjager", "fernet", "salmiakki"},
		Bool:        true,
		Bools:       []bool{true, false, true},
		Byte:        'z',
		Bytes:       []byte("abc"),
		Int:         42,
		Ints:        []int{0, 1, 3, 4},
		Int8:        8,
		Int8s:       []int8{8, 9, 10},
		Int16:       16,
		Int16s:      []int16{16, 17, 18, 19},
		Int32:       32,
		Int32s:      []int32{32, 33},
		Int64:       64,
		Int64s:      []int64{64},
		Uint:        420,
		Uints:       []uint{11, 12, 13},
		Uint8:       81,
		Uint8s:      []uint8{81, 82},
		Uint16:      160,
		Uint16s:     []uint16{160, 161, 162, 163, 164},
		Uint32:      320,
		Uint32s:     []uint32{320, 321},
		Uint64:      640,
		Uint64s:     []uint64{6400, 6401, 6402, 6403},
		Float32:     32.32,
		Float32s:    []float32{32.32, 33},
		Float64:     64.1,
		Float64s:    []float64{64, 65, 66},
		Complex64:   complex64(-64 + 12i),
		Complex64s:  []complex64{complex64(-65 + 11i), complex64(66 + 10i)},
		Complex128:  complex128(-128 + 12i),
		Complex128s: []complex128{complex128(-128 + 11i), complex128(129 + 10i)},
		Interfaces:  []interface{}{42, true, "pan-galactic"},
	}

	cpy := Copy(test).(Basics)

	// see if they point to the same location
	if fmt.Sprintf("%p", &cpy) == fmt.Sprintf("%p", &test) {
		t.Error("address of copy was the same as original; they should be different")
		return
	}

	// Go through each field and check to see it got copied properly
	if cpy.String != test.String {
		t.Errorf("String: got %v; want %v", cpy.String, test.String)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Strings)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Strings)).Data {
		t.Error("Strings: address of copy was the same as original; they should be different")
		goto StringArr
	}

	if len(cpy.Strings) != len(test.Strings) {
		t.Errorf("Strings: len was %d; want %d", len(cpy.Strings), len(test.Strings))
		goto StringArr
	}
	for i, v := range test.Strings {
		if v != cpy.Strings[i] {
			t.Errorf("Strings: got %v at index %d of the copy; want %v", cpy.Strings[i], i, v)
		}
	}

StringArr:
	if unsafe.Pointer(&test.StringArr) == unsafe.Pointer(&cpy.StringArr) {
		t.Error("StringArr: address of copy was the same as original; they should be different")
		goto Bools
	}
	for i, v := range test.StringArr {
		if v != cpy.StringArr[i] {
			t.Errorf("StringArr: got %v at index %d of the copy; want %v", cpy.StringArr[i], i, v)
		}
	}

Bools:
	if cpy.Bool != test.Bool {
		t.Errorf("Bool: got %v; want %v", cpy.Bool, test.Bool)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Bools)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Bools)).Data {
		t.Error("Bools: address of copy was the same as original; they should be different")
		goto Bytes
	}
	if len(cpy.Bools) != len(test.Bools) {
		t.Errorf("Bools: len was %d; want %d", len(cpy.Bools), len(test.Bools))
		goto Bytes
	}
	for i, v := range test.Bools {
		if v != cpy.Bools[i] {
			t.Errorf("Bools: got %v at index %d of the copy; want %v", cpy.Bools[i], i, v)
		}
	}

Bytes:
	if cpy.Byte != test.Byte {
		t.Errorf("Byte: got %v; want %v", cpy.Byte, test.Byte)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Bytes)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Bytes)).Data {
		t.Error("Bytes: address of copy was the same as original; they should be different")
		goto Ints
	}
	if len(cpy.Bytes) != len(test.Bytes) {
		t.Errorf("Bytes: len was %d; want %d", len(cpy.Bytes), len(test.Bytes))
		goto Ints
	}
	for i, v := range test.Bytes {
		if v != cpy.Bytes[i] {
			t.Errorf("Bytes: got %v at index %d of the copy; want %v", cpy.Bytes[i], i, v)
		}
	}

Ints:
	if cpy.Int != test.Int {
		t.Errorf("Int: got %v; want %v", cpy.Int, test.Int)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Ints)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Ints)).Data {
		t.Error("Ints: address of copy was the same as original; they should be different")
		goto Int8s
	}
	if len(cpy.Ints) != len(test.Ints) {
		t.Errorf("Ints: len was %d; want %d", len(cpy.Ints), len(test.Ints))
		goto Int8s
	}
	for i, v := range test.Ints {
		if v != cpy.Ints[i] {
			t.Errorf("Ints: got %v at index %d of the copy; want %v", cpy.Ints[i], i, v)
		}
	}

Int8s:
	if cpy.Int8 != test.Int8 {
		t.Errorf("Int8: got %v; want %v", cpy.Int8, test.Int8)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Int8s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Int8s)).Data {
		t.Error("Int8s: address of copy was the same as original; they should be different")
		goto Int16s
	}
	if len(cpy.Int8s) != len(test.Int8s) {
		t.Errorf("Int8s: len was %d; want %d", len(cpy.Int8s), len(test.Int8s))
		goto Int16s
	}
	for i, v := range test.Int8s {
		if v != cpy.Int8s[i] {
			t.Errorf("Int8s: got %v at index %d of the copy; want %v", cpy.Int8s[i], i, v)
		}
	}

Int16s:
	if cpy.Int16 != test.Int16 {
		t.Errorf("Int16: got %v; want %v", cpy.Int16, test.Int16)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Int16s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Int16s)).Data {
		t.Error("Int16s: address of copy was the same as original; they should be different")
		goto Int32s
	}
	if len(cpy.Int16s) != len(test.Int16s) {
		t.Errorf("Int16s: len was %d; want %d", len(cpy.Int16s), len(test.Int16s))
		goto Int32s
	}
	for i, v := range test.Int16s {
		if v != cpy.Int16s[i] {
			t.Errorf("Int16s: got %v at index %d of the copy; want %v", cpy.Int16s[i], i, v)
		}
	}

Int32s:
	if cpy.Int32 != test.Int32 {
		t.Errorf("Int32: got %v; want %v", cpy.Int32, test.Int32)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Int32s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Int32s)).Data {
		t.Error("Int32s: address of copy was the same as original; they should be different")
		goto Int64s
	}
	if len(cpy.Int32s) != len(test.Int32s) {
		t.Errorf("Int32s: len was %d; want %d", len(cpy.Int32s), len(test.Int32s))
		goto Int64s
	}
	for i, v := range test.Int32s {
		if v != cpy.Int32s[i] {
			t.Errorf("Int32s: got %v at index %d of the copy; want %v", cpy.Int32s[i], i, v)
		}
	}

Int64s:
	if cpy.Int64 != test.Int64 {
		t.Errorf("Int64: got %v; want %v", cpy.Int64, test.Int64)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Int64s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Int64s)).Data {
		t.Error("Int64s: address of copy was the same as original; they should be different")
		goto Uints
	}
	if len(cpy.Int64s) != len(test.Int64s) {
		t.Errorf("Int64s: len was %d; want %d", len(cpy.Int64s), len(test.Int64s))
		goto Uints
	}
	for i, v := range test.Int64s {
		if v != cpy.Int64s[i] {
			t.Errorf("Int64s: got %v at index %d of the copy; want %v", cpy.Int64s[i], i, v)
		}
	}

Uints:
	if cpy.Uint != test.Uint {
		t.Errorf("Uint: got %v; want %v", cpy.Uint, test.Uint)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Uints)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Uints)).Data {
		t.Error("Uints: address of copy was the same as original; they should be different")
		goto Uint8s
	}
	if len(cpy.Uints) != len(test.Uints) {
		t.Errorf("Uints: len was %d; want %d", len(cpy.Uints), len(test.Uints))
		goto Uint8s
	}
	for i, v := range test.Uints {
		if v != cpy.Uints[i] {
			t.Errorf("Uints: got %v at index %d of the copy; want %v", cpy.Uints[i], i, v)
		}
	}

Uint8s:
	if cpy.Uint8 != test.Uint8 {
		t.Errorf("Uint8: got %v; want %v", cpy.Uint8, test.Uint8)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Uint8s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Uint8s)).Data {
		t.Error("Uint8s: address of copy was the same as original; they should be different")
		goto Uint16s
	}
	if len(cpy.Uint8s) != len(test.Uint8s) {
		t.Errorf("Uint8s: len was %d; want %d", len(cpy.Uint8s), len(test.Uint8s))
		goto Uint16s
	}
	for i, v := range test.Uint8s {
		if v != cpy.Uint8s[i] {
			t.Errorf("Uint8s: got %v at index %d of the copy; want %v", cpy.Uint8s[i], i, v)
		}
	}

Uint16s:
	if cpy.Uint16 != test.Uint16 {
		t.Errorf("Uint16: got %v; want %v", cpy.Uint16, test.Uint16)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Uint16s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Uint16s)).Data {
		t.Error("Uint16s: address of copy was the same as original; they should be different")
		goto Uint32s
	}
	if len(cpy.Uint16s) != len(test.Uint16s) {
		t.Errorf("Uint16s: len was %d; want %d", len(cpy.Uint16s), len(test.Uint16s))
		goto Uint32s
	}
	for i, v := range test.Uint16s {
		if v != cpy.Uint16s[i] {
			t.Errorf("Uint16s: got %v at index %d of the copy; want %v", cpy.Uint16s[i], i, v)
		}
	}

Uint32s:
	if cpy.Uint32 != test.Uint32 {
		t.Errorf("Uint32: got %v; want %v", cpy.Uint32, test.Uint32)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Uint32s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Uint32s)).Data {
		t.Error("Uint32s: address of copy was the same as original; they should be different")
		goto Uint64s
	}
	if len(cpy.Uint32s) != len(test.Uint32s) {
		t.Errorf("Uint32s: len was %d; want %d", len(cpy.Uint32s), len(test.Uint32s))
		goto Uint64s
	}
	for i, v := range test.Uint32s {
		if v != cpy.Uint32s[i] {
			t.Errorf("Uint32s: got %v at index %d of the copy; want %v", cpy.Uint32s[i], i, v)
		}
	}

Uint64s:
	if cpy.Uint64 != test.Uint64 {
		t.Errorf("Uint64: got %v; want %v", cpy.Uint64, test.Uint64)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Uint64s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Uint64s)).Data {
		t.Error("Uint64s: address of copy was the same as original; they should be different")
		goto Float32s
	}
	if len(cpy.Uint64s) != len(test.Uint64s) {
		t.Errorf("Uint64s: len was %d; want %d", len(cpy.Uint64s), len(test.Uint64s))
		goto Float32s
	}
	for i, v := range test.Uint64s {
		if v != cpy.Uint64s[i] {
			t.Errorf("Uint64s: got %v at index %d of the copy; want %v", cpy.Uint64s[i], i, v)
		}
	}

Float32s:
	if cpy.Float32 != test.Float32 {
		t.Errorf("Float32: got %v; want %v", cpy.Float32, test.Float32)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Float32s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Float32s)).Data {
		t.Error("Float32s: address of copy was the same as original; they should be different")
		goto Float64s
	}
	if len(cpy.Float32s) != len(test.Float32s) {
		t.Errorf("Float32s: len was %d; want %d", len(cpy.Float32s), len(test.Float32s))
		goto Float64s
	}
	for i, v := range test.Float32s {
		if v != cpy.Float32s[i] {
			t.Errorf("Float32s: got %v at index %d of the copy; want %v", cpy.Float32s[i], i, v)
		}
	}

Float64s:
	if cpy.Float64 != test.Float64 {
		t.Errorf("Float64: got %v; want %v", cpy.Float64, test.Float64)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Float64s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Float64s)).Data {
		t.Error("Float64s: address of copy was the same as original; they should be different")
		goto Complex64s
	}
	if len(cpy.Float64s) != len(test.Float64s) {
		t.Errorf("Float64s: len was %d; want %d", len(cpy.Float64s), len(test.Float64s))
		goto Complex64s
	}
	for i, v := range test.Float64s {
		if v != cpy.Float64s[i] {
			t.Errorf("Float64s: got %v at index %d of the copy; want %v", cpy.Float64s[i], i, v)
		}
	}

Complex64s:
	if cpy.Complex64 != test.Complex64 {
		t.Errorf("Complex64: got %v; want %v", cpy.Complex64, test.Complex64)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Complex64s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Complex64s)).Data {
		t.Error("Complex64s: address of copy was the same as original; they should be different")
		goto Complex128s
	}
	if len(cpy.Complex64s) != len(test.Complex64s) {
		t.Errorf("Complex64s: len was %d; want %d", len(cpy.Complex64s), len(test.Complex64s))
		goto Complex128s
	}
	for i, v := range test.Complex64s {
		if v != cpy.Complex64s[i] {
			t.Errorf("Complex64s: got %v at index %d of the copy; want %v", cpy.Complex64s[i], i, v)
		}
	}

Complex128s:
	if cpy.Complex128 != test.Complex128 {
		t.Errorf("Complex128s: got %v; want %v", cpy.Complex128s, test.Complex128s)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Complex128s)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Complex128s)).Data {
		t.Error("Complex128s: address of copy was the same as original; they should be different")
		goto Interfaces
	}
	if len(cpy.Complex128s) != len(test.Complex128s) {
		t.Errorf("Complex128s: len was %d; want %d", len(cpy.Complex128s), len(test.Complex128s))
		goto Interfaces
	}
	for i, v := range test.Complex128s {
		if v != cpy.Complex128s[i] {
			t.Errorf("Complex128s: got %v at index %d of the copy; want %v", cpy.Complex128s[i], i, v)
		}
	}

Interfaces:
	if cpy.Interface != test.Interface {
		t.Errorf("Interface: got %v; want %v", cpy.Interface, test.Interface)
	}

	if (*reflect.SliceHeader)(unsafe.Pointer(&test.Interfaces)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpy.Interfaces)).Data {
		t.Error("Interfaces: address of copy was the same as original; they should be different")
		return
	}
	if len(cpy.Interfaces) != len(test.Interfaces) {
		t.Errorf("Interfaces: len was %d; want %d", len(cpy.Interfaces), len(test.Interfaces))
		return
	}
	for i, v := range test.Interfaces {
		if v != cpy.Interfaces[i] {
			t.Errorf("Interfaces: got %v at index %d of the copy; want %v", cpy.Interfaces[i], i, v)
		}
	}
}

// not meant to be exhaustive
func TestComplexSlices(t *testing.T) {
	orig3Int := [][][]int{[][]int{[]int{1, 2, 3}, []int{11, 22, 33}}, [][]int{[]int{7, 8, 9}, []int{66, 77, 88, 99}}}
	cpyI := Copy(orig3Int).([][][]int)
	if (*reflect.SliceHeader)(unsafe.Pointer(&orig3Int)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&cpyI)).Data {
		t.Error("[][][]int: address of copy was the same as original; they should be different")
		return
	}
	if len(orig3Int) != len(cpyI) {
		t.Errorf("[][][]int: len of copy was %d; want %d", len(cpyI), len(orig3Int))
	}
	for i, v := range orig3Int {
		if len(v) != len(cpyI[i]) {
			t.Errorf("[][][]int: len of element %d was %d; want %d", i, len(cpyI[i]), len(v))
			continue
		}
		for j, vv := range v {
			if len(vv) != len(cpyI[i][j]) {
				t.Errorf("[][][]int: len of element %d:%d was %d, want %d", i, j, len(cpyI[i][j]), len(vv))
				continue
			}
			for k, vvv := range vv {
				if vvv != cpyI[i][j][k] {
					t.Errorf("[][][]int: element %d:%d:%d was %d, want %d", i, j, k, cpyI[i][j][k], vvv)
				}
			}
		}

	}
}

type A struct {
	Int    int
	String string
	UintSl []uint
	NilSl  []string
	Map    map[string]int
	MapB   map[string]*B
	SliceB []B
	B
	T time.Time
}

type B struct {
	Vals []string
}

var AStruct = A{
	Int:    42,
	String: "Konichiwa",
	UintSl: []uint{0, 1, 2, 3},
	Map:    map[string]int{"a": 1, "b": 2},
	MapB: map[string]*B{
		"hi":  &B{Vals: []string{"hello", "bonjour"}},
		"bye": &B{Vals: []string{"good-bye", "au revoir"}},
	},
	SliceB: []B{
		B{Vals: []string{"Ciao", "Aloha"}},
	},
	B: B{Vals: []string{"42"}},
	T: time.Now(),
}

func TestStructA(t *testing.T) {
	cpy := Copy(AStruct).(A)
	if &cpy == &AStruct {
		t.Error("expected copy to have a different address than the original; it was the same")
		return
	}
	if cpy.Int != AStruct.Int {
		t.Errorf("A.Int: got %v, want %v", cpy.Int, AStruct.Int)
	}
	if cpy.String != AStruct.String {
		t.Errorf("A.String: got %v; want %v", cpy.String, AStruct.String)
	}
	if (*reflect.SliceHeader)(unsafe.Pointer(&cpy.UintSl)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&AStruct.UintSl)).Data {
		t.Error("A.Uintsl: expected the copies address to be different; it wasn't")
		goto NilSl
	}
	if len(cpy.UintSl) != len(AStruct.UintSl) {
		t.Errorf("A.UintSl: got len of %d, want %d", len(cpy.UintSl), len(AStruct.UintSl))
		goto NilSl
	}
	for i, v := range AStruct.UintSl {
		if cpy.UintSl[i] != v {
			t.Errorf("A.UintSl %d: got %d, want %d", i, cpy.UintSl[i], v)
		}
	}

NilSl:
	if cpy.NilSl != nil {
		t.Error("A.NilSl: expected slice to be nil, it wasn't")
	}

}

func TestPointerToStruct(t *testing.T) {
	type Foo struct {
		Bar int
	}

	f := &Foo{Bar: 42}
	cpy := Copy(f)
	if f == cpy {
		t.Errorf("expected copy to point to a different location: orig: %p; copy: %p", f, cpy)
	}
	if !reflect.DeepEqual(f, cpy) {
		t.Errorf("expected the copy to be equal to the original (except for memory location); it wasn't: got %#v; want %#v", f, cpy)
	}
}

type I struct {
	A string
}

func (i *I) DeepCopy() interface{} {
	return &I{A: "custom copy"}
}

type NestI struct {
	I *I
}
