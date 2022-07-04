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
			MaskSensitiveData(f.Interface())
		case reflect.Slice:
			for j := 0; j < f.Len(); j++ {
				MaskSensitiveData(f.Index(j).Interface())
			}
		case reflect.String:
			processString(val.Type().Field(i).Tag.Get("masked"), f)
		}
	}
}

func processString(masked string, f reflect.Value) {
	if isTrue(masked) && f.CanSet() && f.String() != "" {
		f.SetString(maskedValue)
	}
}

func isTrue(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return b
}
