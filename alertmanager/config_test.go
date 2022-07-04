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
