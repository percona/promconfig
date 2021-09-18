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

package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

var testData = `groups:
  - name: example
    rules:
      - record: job:http_inprogress_requests:sum
        expr: sum by (job) (http_inprogress_requests)
`

func TestGroup(t *testing.T) {
	groups := RuleGroups{
		Groups: []RuleGroup{
			{
				Name: "example",
				Rules: []Rule{
					{
						Record: "job:http_inprogress_requests:sum",
						Expr:   "sum by (job) (http_inprogress_requests)",
					},
				},
			},
		},
	}
	data, err := yaml.Marshal(groups)
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, testData, string(data))
}
