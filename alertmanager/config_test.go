package alertmanager

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMaskSensitiveData(t *testing.T) {
	c := &Config{
		Global: &GlobalConfig{
			SMTPAuthUsername: "username",
			SMTPAuthPassword: "password",
		},
	}
	MaskSensitiveData(c)
	spew.Dump(c)
}
