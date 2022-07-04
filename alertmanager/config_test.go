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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskSensitiveData(t *testing.T) {
	testCases := []struct {
		Config   *Config
		Expected *Config
	}{
		{
			Config: &Config{
				Global: &GlobalConfig{
					SMTPAuthUsername: "username",
					SMTPAuthPassword: "password",
				},
			},
			Expected: &Config{
				Global: &GlobalConfig{
					SMTPAuthUsername: maskedValue,
					SMTPAuthPassword: maskedValue,
				},
			},
		},
	}

	for _, testCase := range testCases {

		MaskSensitiveData(testCase.Config)
		assert.Equal(t, testCase.Config, testCase.Expected)
	}
}
