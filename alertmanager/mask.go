// promconfig
// Copyright 2020 Percona LLC
//
// Based on Prometheus systems and service monitoring server.
// Copyright 2015 The Prometheus Authors.
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

package alertmanager

import (
	"reflect"
	"strconv"
)

const maskedValue = "xxxxxxxx"

// MaskSensitiveData loops over Config struct and masks sensitive data.
func MaskSensitiveData(c interface{}) {
	val := reflect.ValueOf(c)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		switch f.Kind() {
		case reflect.Ptr:
			if f.IsNil() {
				continue
			}
			MaskSensitiveData(f.Interface())
		case reflect.Struct:
			MaskSensitiveData(f.Addr().Interface())
		case reflect.Slice:
			for j := 0; j < f.Len(); j++ {
				MaskSensitiveData(f.Index(j).Interface())
			}
		case reflect.String:
			// masked struct tag must be equal to true to mask values
			processString(val.Type().Field(i).Tag.Get("masked"), f)
			//missing cases in switch of type reflect.Kind: Array, Bool, Chan, Complex128, Complex64, Float32, Float64, Func, Int, Int16, Int32, Int64, Int8, Interface, Inval
			// id, Map, Uint, Uint16, Uint32, Uint64, Uint8, Uintptr, UnsafePointer
		case reflect.Bool, reflect.Chan:
		case reflect.Array:
		case reflect.Complex128, reflect.Complex64:
		case reflect.Float32, reflect.Float64:
		case reflect.Func:
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		case reflect.Uintptr, reflect.UnsafePointer:
		}
	}
}

// processString checks masked tag and masks sensitive values.
func processString(masked string, f reflect.Value) {
	if isTrue(masked) && f.CanSet() && f.String() != "" {
		f.SetString(maskedValue)
	}
}

// isTrue is a helper function to convert string to boolean.
func isTrue(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return b
}
