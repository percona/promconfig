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
	"github.com/percona/promconfig"
)

// RuleGroups is a set of rule groups that are typically exposed in a file.
type RuleGroups struct {
	Groups []RuleGroup `yaml:"groups"`
}

// RuleGroup is a list of sequentially evaluated recording and alerting rules.
type RuleGroup struct {
	Name     string `yaml:"name"`
	Interval string `yaml:"interval,omitempty"`
	Rules    []Rule `yaml:"rules"`
}

// Rule describes an alerting or recording rule.
type Rule struct {
	Record      string              `yaml:"record,omitempty"`
	Alert       string              `yaml:"alert,omitempty"`
	Expr        string              `yaml:"expr"`
	For         promconfig.Duration `yaml:"for,omitempty"`
	Labels      map[string]string   `yaml:"labels,omitempty"`
	Annotations map[string]string   `yaml:"annotations,omitempty"`
}
