package rules_test

import (
	"testing"

	"github.com/percona/promconfig/rules"
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
	groups := rules.RuleGroups{
		Groups: []rules.RuleGroup{
			{
				Name: "example",
				Rules: []rules.RuleNode{
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
