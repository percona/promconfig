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

package promconfig

import "reflect"

func MaskSecret(c interface{}) {
	val := reflect.ValueOf(c)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		switch f.Kind() { //nolint:exhaustive
		case reflect.Ptr:
			if f.IsNil() {
				continue
			}
			MaskSecret(f.Interface())
		case reflect.Struct:
			MaskSecret(f.Addr().Interface())
		case reflect.Slice:
			for j := 0; j < f.Len(); j++ {
				MaskSecret(f.Index(j).Interface())
			}
		case reflect.String:
			// Mask only Secret datatypes
			if f.Type().Name() == "Secret" && f.CanSet() && f.String() != "" {
				f.SetString(secretToken)
			}
		}
	}
}
