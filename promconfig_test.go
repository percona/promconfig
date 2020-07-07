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

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

var goldenF = flag.Bool("golden", false, "update both golden .json files files")

func TestGoldenData(t *testing.T) {
	matches, err := filepath.Glob("testdata/*.yml")
	require.NoError(t, err)
	require.NotEmpty(t, matches)

	for _, yf := range matches {
		b, err := ioutil.ReadFile(yf)
		require.NoError(t, err)

		var cfg Config
		err = yaml.Unmarshal(b, &cfg)
		require.NoError(t, err)
		actualB, err := json.MarshalIndent(cfg, "", "  ")
		require.NoError(t, err)
		actualB = append(actualB, '\n')

		jf := strings.TrimSuffix(yf, filepath.Ext(yf)) + ".json"

		if *goldenF {
			err = ioutil.WriteFile(jf, actualB, 0644)
			require.NoError(t, err)
		}

		expectedB, err := ioutil.ReadFile(jf)
		require.NoError(t, err)

		expectedS := string(expectedB)
		actualS := string(actualB)
		assert.Equal(t, expectedS, actualS)
	}
}
