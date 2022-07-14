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

// HTTPClientConfig configures an HTTP client.
type HTTPClientConfig struct {
	// The HTTP basic authentication credentials for the targets.
	BasicAuth *BasicAuth `yaml:"basic_auth,omitempty"`
	// The HTTP authorization credentials for the targets.
	Authorization *Authorization `yaml:"authorization,omitempty"`
	// The OAuth2 client credentials used to fetch a token for the targets.
	OAuth2 *OAuth2 `yaml:"oauth2,omitempty"`
	// The bearer token for the targets.
	// Deprecated: use Authorization.Credentials instead.
	BearerToken string `yaml:"bearer_token,omitempty" secret:"true"`
	// The bearer token file for the targets.
	// Deprecated: use Authorization.Credentials instead.
	BearerTokenFile string `yaml:"bearer_token_file,omitempty"`
	// HTTP proxy server to use to connect to the targets.
	ProxyURL string `yaml:"proxy_url,omitempty"`
	// TLSConfig to use to connect to the targets.
	TLSConfig TLSConfig `yaml:"tls_config,omitempty"`
	// FollowRedirects specifies whether the client should follow HTTP 3xx redirects.
	// The omitempty flag is not set, because it would be hidden from the
	// marshalled configuration when set to false.
	FollowRedirects bool `yaml:"follow_redirects"`
}

// Authorization contains HTTP authorization credentials.
type Authorization struct {
	Type            string `yaml:"type,omitempty"`
	Credentials     string `yaml:"credentials,omitempty" secret:"true"`
	CredentialsFile string `yaml:"credentials_file,omitempty"`
}

// OAuth2 is the oauth2 client configuration.
type OAuth2 struct {
	ClientID         string            `yaml:"client_id"`
	ClientSecret     string            `yaml:"client_secret" secret:"true"`
	ClientSecretFile string            `yaml:"client_secret_file"`
	Scopes           []string          `yaml:"scopes,omitempty"`
	TokenURL         string            `yaml:"token_url"`
	EndpointParams   map[string]string `yaml:"endpoint_params,omitempty"`
}

// BasicAuth contains basic HTTP authentication credentials.
type BasicAuth struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password,omitempty" secret:"true"`
	PasswordFile string `yaml:"password_file,omitempty"`
}

// TLSConfig configures the options for TLS connections.
type TLSConfig struct {
	// The CA cert to use for the targets.
	CAFile string `yaml:"ca_file,omitempty"`
	// The client cert file for the targets.
	CertFile string `yaml:"cert_file,omitempty"`
	// The client key file for the targets.
	KeyFile string `yaml:"key_file,omitempty"`
	// Used to verify the hostname for the targets.
	ServerName string `yaml:"server_name,omitempty"`
	// Disable target certificate validation.
	InsecureSkipVerify bool `yaml:"insecure_skip_verify"`
}
