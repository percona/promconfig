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
	common "github.com/percona/promconfig/common"
)

// A Route is a node that contains definitions of how to handle alerts.
type Route struct {
	Receiver string `yaml:"receiver,omitempty" json:"receiver,omitempty"`

	GroupByStr []string `yaml:"group_by,omitempty" json:"group_by,omitempty"`
	GroupBy    []string `yaml:"-" json:"-"`
	GroupByAll bool     `yaml:"-" json:"-"`

	Match    map[string]string `yaml:"match,omitempty" json:"match,omitempty"`
	MatchRE  MatchRegexps      `yaml:"match_re,omitempty" json:"match_re,omitempty"`
	Continue bool              `yaml:"continue" json:"continue,omitempty"`
	Routes   []*Route          `yaml:"routes,omitempty" json:"routes,omitempty"`

	GroupWait      common.Duration `yaml:"group_wait,omitempty" json:"group_wait,omitempty"`
	GroupInterval  common.Duration `yaml:"group_interval,omitempty" json:"group_interval,omitempty"`
	RepeatInterval common.Duration `yaml:"repeat_interval,omitempty" json:"repeat_interval,omitempty"`
}
