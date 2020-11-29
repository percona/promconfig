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

// ServiceDiscoveryConfig configures lists of different service discovery mechanisms.
type ServiceDiscoveryConfig struct {
	// List of labeled target groups for this job.
	StaticConfigs []*Group `yaml:"static_configs,omitempty"`
	// List of file service discovery configurations.
	FileSDConfigs []*FilesSDConfig `yaml:"file_sd_configs,omitempty"`
	// List of Kubernetes service discovery configurations.
	KubernetesSDConfigs []*KubernetesSDConfig `yaml:"kubernetes_sd_configs,omitempty"`
	// List of AWS EC2 service discovery configurations.
	EC2SDConfigs []*EC2SDConfig `yaml:"ec2_sd_configs,omitempty"`
	// List of Google cloud GCE service discovery configurations.
	GceSDConfigs []*GceSDConfig `yaml:"gce_sd_configs,omitempty"`
	// List of azure cloud service discovery configurations.
	AzureSDConfigs []*AzureSDConfig `yaml:"azure_sd_configs,omitempty"`
	// List of digitalocean droplet service discovery configurations.
	DigitaloceanSDConfigs []*DigitaloceanSDConfig `yaml:"digitalocean_sd_configs,omitempty"`
	// List of consul cataloge service discovery configurations.
	ConsulSDConfigs []*ConsulSDConfig `yaml:"consul_sd_configs,omitempty"`
	// List of docker swarm service discovery configurations.
	DockerswarmSDConfigs []*DockerswarmSDConfig `yaml:"dockerswarm_sd_configs,omitempty"`
	// List of dns-based service discovery configurations.
	DNSSDConfigs []*DNSSDConfig `yaml:"dns_sd_configs,omitempty"`
}

// Group is a set of targets with a common label set(production , test, staging etc.).
type Group struct {
	// Targets is a list of targets identified by a label set. Each target is
	// uniquely identifiable in the group by its address label.
	Targets []string `yaml:"targets,omitempty"`
	// Labels is a set of labels that is common across all targets in the group.
	Labels map[string]string `yaml:"labels,omitempty"`
}

// FilesSDConfig is the configuration for file based discovery.
type FilesSDConfig struct {
	Files           []string `yaml:"files"`
	RefreshInterval Duration `yaml:"refresh_interval,omitempty"`
}

// KubernetesSDConfig is the configuration for Kubernetes service discovery.
type KubernetesSDConfig struct {
	APIServer          string           `yaml:"api_server,omitempty"`
	Role               string           `yaml:"role"`
	HTTPClientConfig   HTTPClientConfig `yaml:",inline"`
	NamespaceDiscovery []string         `yaml:"namespaces,omitempty"`
}

// EC2SDConfig is the configuration for AWS EC2 instance service discovery.
type EC2SDConfig struct {
	Region          string    `yaml:"region,omitempty"`
	Endpoint        string    `yaml:"endpoint,omitempty"`
	AccessKey       string    `yaml:"access_key,omitempty"`
	SecretKey       string    `yaml:"secret_key,omitempty"`
	Profile         string    `yaml:"profile,omitempty"`
	RoleArn         string    `yaml:"role_arn,omitempty"`
	RefreshInterval Duration  `yaml:"refresh_interval,omitempty"`
	Port            int       `yaml:"port,omitempty"`
	Filters         []*Filter `yaml:"filters,omitempty"`
}

// GceSDConfig is the configuration for Google cloud GCE instance service discovery
type GceSDConfig struct {
	Project         string   `yaml:"project"`
	Zone            string   `yaml:"zone"`
	Filter          string   `yaml:"filter,omitempty"`
	RefreshInterval Duration `yaml:"refresh_interval,omitempty"`
	Port            int      `yaml:"port,omitempty"`
	TagSeprator     string   `yaml:"tag_separator,omitempty"`
}

// AzureSDConfig is the configuration for Azure cloud service discovery
type AzureSDConfig struct {
	Environment     string   `yaml:"environment,omitempty"`
	SubscriptionID  string   `yaml:"subscription_id"`
	TenantID        string   `yaml:"tenant_id,omitempty"`
	ClientID        string   `yaml:"client_id,omitempty"`
	ClientSecret    string   `yaml:"client_secret,omitempty"`
	RefreshInterval Duration `yaml:"refresh_interval,omitempty"`
	Port            int      `yaml:"port,omitempty"`
}

// DigitaloceanSDConfig is the configuration for digitalocean droplet service discovery
type DigitaloceanSDConfig struct {
	HTTPClientConfig HTTPClientConfig `yaml:",inline"`
	RefreshInterval  Duration         `yaml:"refresh_interval,omitempty"`
	Port             int              `yaml:"port,omitempty"`
}

// ConsulSDConfig is the configuration for the consul cataloge service discovery
type ConsulSDConfig struct {
	Server          string            `yaml:"server,omitempty"`
	Token           string            `yaml:"token"`
	Datacenter      string            `yaml:"datacenter"`
	Scheme          string            `yaml:"scheme,omitempty"`
	Username        string            `yaml:"username"`
	Password        string            `yaml:"password"`
	TLSConfig       TLSConfig         `yaml:"tls_config,omitempty"`
	Services        []string          `yaml:"services,omitempty"`
	Tags            []string          `yaml:"tags,omitempty"`
	NodeMeta        map[string]string `yaml:"node_meta,omitempty"`
	TagSeprator     string            `yaml:"tag_seprator,omitempty"`
	AllowStale      bool              `yaml:"allow_stale"`
	RefreshInterval Duration          `yaml:"refresh_interval,omitempty"`
}

// DockerswarmSDConfig is the configuration for service discovery of docker services, tasks or nodes.
type DockerswarmSDConfig struct {
	HTTPClientConfig HTTPClientConfig `yaml:",inline"`
	Host             string           `yaml:"host"`
	Role             string           `yaml:"role"`
	RefreshInterval  Duration         `yaml:"refresh_interval,omitempty"`
	Port             int              `yaml:"port,omitempty"`
	Filters          []*Filter        `yaml:"filters,omitempty"`
}

// DNSSDConfig is configuration for dns based service discovery.
type DNSSDConfig struct {
	Names           []string `yaml:"names"`
	Type            string   `yaml:"type,omitempty"`
	RefreshInterval Duration `yaml:"refresh_interval,omitempty"`
	Port            int      `yaml:"port,omitempty"`
}

// Filter to limit service discovery
type Filter struct {
	Name   string   `yaml:"name"`
	Values []string `yaml:"values"`
}
