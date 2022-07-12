package promconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestMarshalYAML(t *testing.T) {
	type secret struct {
		S Secret
	}
	tmp := secret{S: "Something"}
	data, err := yaml.Marshal(tmp)
	assert.NoError(t, err)

	tmp2 := &secret{}
	err = yaml.Unmarshal(data, &tmp2)
	assert.NoError(t, err)
	assert.Equal(t, tmp2.S.String(), secretToken)
}
