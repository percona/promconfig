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

	"github.com/percona/promconfig"
)

func TestMask(t *testing.T) {
	maskedValue := promconfig.Secret("<secret>")
	t.Parallel()

	testCases := []struct {
		Name     string
		Config   *Config
		Expected *Config
	}{
		{
			Name: "global configuration variables should be masked",
			Config: &Config{
				Global: &GlobalConfig{
					SMTPFrom:         "hello@example.com",
					SMTPAuthUsername: "username",
					SMTPAuthPassword: "password",
				},
			},
			Expected: &Config{
				Global: &GlobalConfig{
					SMTPFrom:         "hello@example.com",
					SMTPAuthUsername: maskedValue,
					SMTPAuthPassword: maskedValue,
				},
				// Templates: []string{},
			},
		},
		{
			Name: "should work with struct",
			Config: &Config{
				Global: &GlobalConfig{
					HTTPConfig: promconfig.HTTPClientConfig{
						BearerToken: "Bearer xoxoxoxo",
					},
				},
			},
			Expected: &Config{
				Global: &GlobalConfig{
					HTTPConfig: promconfig.HTTPClientConfig{
						BearerToken: maskedValue,
					},
				},
				// Templates: []string{},
			},
		},
		{
			Name: "http configuration variables should be masked",
			Config: &Config{
				Global: &GlobalConfig{
					HTTPConfig: promconfig.HTTPClientConfig{
						BasicAuth: &promconfig.BasicAuth{
							Username:     "username",
							Password:     "password",
							PasswordFile: "/etc/passwd",
						},
						Authorization: &promconfig.Authorization{
							Type:            "bearer",
							Credentials:     "Something",
							CredentialsFile: "Meaningful",
						},
						OAuth2: &promconfig.OAuth2{
							ClientID:         "supersecret",
							ClientSecret:     "13r-0ihgf2r2n-i",
							ClientSecretFile: "/etc/passwd",
						},
						BearerToken: "Bearer xoxoxoxo",
					},
				},
			},
			Expected: &Config{
				Global: &GlobalConfig{
					HTTPConfig: promconfig.HTTPClientConfig{
						BasicAuth: &promconfig.BasicAuth{
							Username:     "username",
							Password:     maskedValue,
							PasswordFile: "/etc/passwd",
						},
						Authorization: &promconfig.Authorization{
							Type:            "bearer",
							Credentials:     maskedValue,
							CredentialsFile: "Meaningful",
						},
						OAuth2: &promconfig.OAuth2{
							ClientID:         "supersecret",
							ClientSecret:     maskedValue,
							ClientSecretFile: "/etc/passwd",
						},
						BearerToken: maskedValue,
					},
				},
				// Templates: []string{},
			},
		},
		{
			Name: "receiver configuration sensitive values should be masked",
			Config: &Config{
				Receivers: []*Receiver{
					&Receiver{
						EmailConfigs: []*EmailConfig{
							&EmailConfig{
								AuthUsername: "username",
								AuthPassword: "password",
							},
						},
					},
					&Receiver{
						PagerdutyConfigs: []*PagerdutyConfig{
							&PagerdutyConfig{
								HTTPConfig: promconfig.HTTPClientConfig{
									BasicAuth: &promconfig.BasicAuth{
										Username: "username",
										Password: "pa$$w0rd",
									},
								},
								ServiceKey: "asid-123ernkasd",
								RoutingKey: "HelloRouteK3y",
							},
						},
					},
					&Receiver{
						SlackConfigs: []*SlackConfig{
							&SlackConfig{
								HTTPConfig: promconfig.HTTPClientConfig{
									BasicAuth: &promconfig.BasicAuth{
										Username: "username",
										Password: "pa$$w0rd",
									},
								},
								Channel:  "general",
								Username: "<USOmeting>",
							},
						},
					},
					&Receiver{
						OpsGenieConfigs: []*OpsGenieConfig{
							&OpsGenieConfig{
								HTTPConfig: promconfig.HTTPClientConfig{
									BasicAuth: &promconfig.BasicAuth{
										Username: "username",
										Password: "pa$$w0rd",
									},
								},
								APIKey: "key",
								APIURL: "url",
							},
						},
					},
				},
			},
			Expected: &Config{
				Receivers: []*Receiver{
					&Receiver{
						EmailConfigs: []*EmailConfig{
							&EmailConfig{
								AuthUsername: maskedValue,
								AuthPassword: maskedValue,
							},
						},
					},
					&Receiver{
						PagerdutyConfigs: []*PagerdutyConfig{
							&PagerdutyConfig{
								HTTPConfig: promconfig.HTTPClientConfig{
									BasicAuth: &promconfig.BasicAuth{
										Username: "username",
										Password: maskedValue,
									},
								},
								ServiceKey: maskedValue,
								RoutingKey: maskedValue,
							},
						},
					},
					&Receiver{
						SlackConfigs: []*SlackConfig{
							&SlackConfig{
								HTTPConfig: promconfig.HTTPClientConfig{
									BasicAuth: &promconfig.BasicAuth{
										Username: "username",
										Password: maskedValue,
									},
								},
								Channel:  "general",
								Username: "<USOmeting>",
							},
						},
					},
					&Receiver{
						OpsGenieConfigs: []*OpsGenieConfig{
							&OpsGenieConfig{
								HTTPConfig: promconfig.HTTPClientConfig{
									BasicAuth: &promconfig.BasicAuth{
										Username: "username",
										Password: maskedValue,
									},
								},
								APIKey: maskedValue,
								APIURL: "url",
							},
						},
					},
				},
				// Templates: []string{},
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()
			testCase.Config.Mask()
			assert.Equal(t, testCase.Config, testCase.Expected)
		})
	}
}
